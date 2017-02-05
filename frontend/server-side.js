import Router, { Route } from "preact-router"
import render from "preact-render-to-string"
import components from "./components"
import routes from "../routes.json"

import { h } from "preact"

// "send" is a global variable set by server-side VM context
if (typeof send !== 'undefined') {
  send(renderURL)
}

function renderURL (url, params) {
  return render(routerEl(url))
}

export function routerEl (url) {
  return (
    <Router url={url}>
      {
        routesArray().map(r => <Route {...r} component={components[r.component]} />)
      }
    </Router>
  )
}

export function routesArray () {
  const list = []

  var key;
  for (key in routes) {
    list.push({ path: key, component: routes[key].component })
  }

  return list
}
