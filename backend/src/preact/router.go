package preact

import (
	"encoding/json"
	"github.com/azer/url-router"
	"io/ioutil"
)

func NewRouter() (*Router, error) {
	r := &Router{}

	if err := r.Reload(); err != nil {
		return nil, err
	}

	return r, nil
}

type RoutingOptions struct {
	Component string `json:"component"`
	Title     string `json:"title"`
	Template  string `json:"template"`
}

type Router struct {
	State    int
	Document map[string]*RoutingOptions
	Matcher  *urlrouter.Router
}

func (router *Router) Reload() error {
	if err := router.LoadDocument(); err != nil {
		return err
	}

	router.CreateMatcher()
	return nil
}

func (router *Router) LoadDocument() error {
	doc, err := ioutil.ReadFile("./routes.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(doc, &router.Document)
}

func (router *Router) Match(path string) *RoutingOptions {
	match := router.Matcher.Match(path)
	if match == nil {
		return nil
	}

	return router.Document[match.Pattern]
}

func (router *Router) CreateMatcher() {
	router.Matcher = urlrouter.New()
	for k := range router.Document {
		router.Matcher.Add(k)
	}
}
