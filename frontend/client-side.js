import { h, render } from "preact"
import Router, { Route } from "preact-router"
import { routesArray } from "./server-side"
import components from "./components"

render((
    <Router>
    {
      routesArray().map(r => <Route {...r} component={components[r.component]} />)
    }
  </Router>), document.body)
