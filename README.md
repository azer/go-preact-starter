## go-preact-starter

Starter for combining Go and [Preact](https://github.com/developit/preact) in any web project. 

### How It Works?

* Renders Preact components on serverside using [go-duktape](https://github.com/olebedev/go-duktape)
* Unifies client-side and server-side routing on one JSON file.
* Watches Go, JS, CSS files, compiles them automatically and refreshes your browser.
* Uses Browserify, no configuration needed.
* Supports server-side templating for HTML documents that wraps the application. So you can have multiple pages.
* Locates all Go dependencies inside the project for security by giving you a convenient GOPATH setup.
* Provides caching when development mode is disabled.

### Install

Clone the repo and install the dependencies:

```bash
git clone git@github.com:azer/go-preact-starter.git hello-world
cd hello-world
make install # install dependencies needed
make develop # start developing! visit localhost:9000 to see your website!
```

### Coding Notes

* Create UI components under frontend/components
* Define your routes on `routes.json`, point them to a valid component.
* Run `make go-get` and `make frontend-install` to install new dependencies.
* Edit .env file to choose different host/port for serving.
