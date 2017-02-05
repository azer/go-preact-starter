include .env

PROJECTNAME=$(shell basename "$(PWD)")
GOBASE=$(shell pwd)/backend
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard backend/*.go)
PID=/tmp/$(PROJECTNAME)-backend.pid
FRONTENDBASE=$(shell pwd)/frontend
FRONTENDBIN=$(FRONTENDBASE)/node_modules/.bin

all: help

install: go-get frontend-install
publish: compile-backend compile-frontend start-backend
stop: stop-watching stop-remote-control stop-backend

develop:
	bash -c "trap 'make stop' EXIT; $(MAKE) internal-develop"

internal-develop: watch
	DEVELOP=1 $(MAKE) compile-frontend
	$(MAKE) start-remote-control
	$(MAKE) recompile-backend
	while true;	do sleep 3; done

watch: stop-watching
	@echo "  >  Watching for changes..."
	@$(FRONTENDBIN)/chokidar --silent 'frontend' 'backend' 'routes.json' \
		-i 'backend/pkg' -i 'backend/bin' -i 'frontend/node_modules' \
		-c 'if [[ {path} == *.go ]]; then make recompile-backend; elif [[ {path} == *.css ]]; then make recompile-css; elif [[ {path} == *.js ]]; then make recompile-js; elif [[ "{path}" == *"routes.json"* ]]; then make recompile-backend && make recompile-js; fi' & echo $$! > /tmp/$(PROJECTNAME)-watch.pid

stop-watching:
	@touch /tmp/$(PROJECTNAME)-watch.pid
	@-kill `cat /tmp/$(PROJECTNAME)-watch.pid` 2> /dev/null || true

# Backend commands
go-build:
	@echo "  >  Building backend..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-get:
	@echo "  >  Downloading backend dependencies..."
	@cd backend && GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get .

go-install:
	@cd backend && GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

go-clean:
	@echo "  >  Cleaning Go build cache"
	@cd backend && GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

start-backend: stop-backend
	@echo "  >  Starting $(PROJECTNAME) backend at $(ADDR)"
	@-$(GOBIN)/$(PROJECTNAME) & echo $$! > $(PID)
	@echo "  >  Backend Process ID: "$(shell cat $(PID))

stop-backend:
	@echo "  >  Stopping $(PROJECTNAME) backend"
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true

compile-backend: go-clean go-get go-build
recompile-backend:
	$(MAKE) compile-backend
	LOG=* DEVELOP=1 $(MAKE) start-backend
	@-curl http://localhost:9066/commands/reload 2> /dev/null > /dev/null

# Front-end commands
compile-frontend: compile-js compile-css

compile-js:
	@echo "  >  Compiling JavaScript..."
	@cd frontend && $(FRONTENDBIN)/browserify client-side.js --debug -o ../public/dist.js

compile-css:
	@echo "  >  Compiling CSS..."
	@cd frontend && cat *.css components/**/*.css > ../public/dist.css

recompile-js: compile-js
	curl http://localhost:9966/commands/reload 2> /dev/null > /dev/null

recompile-css: compile-css
	curl http://localhost:9966/commands/reload-css 2> /dev/null > /dev/null

browserify-for-backend:
	@cd frontend && $(FRONTENDBIN)/browserify server-side.js

clean-frontend:
	@echo "  >  Cleaning front-end builds"
	@-rm public/{dist.js,dist.css} 2> /dev/null || true

start-remote-control: stop-remote-control
	@$(FRONTENDBIN)/remote-control -p 9966 -h localhost --quiet & echo $$! > /tmp/$(PROJECTNAME)-remote-control.pid

stop-remote-control:
	@touch /tmp/$(PROJECTNAME)-remote-control.pid
	@-kill `cat /tmp/$(PROJECTNAME)-remote-control.pid` 2> /dev/null || true

frontend-install:
	@echo "  >  Downloading dependencies from NPM."
	@cd frontend && npm install

help:
	@cat docs/man | less

.PHONY: default develop start-remote-control stop-remote-control

ifndef VERBOSE
.SILENT:
endif
