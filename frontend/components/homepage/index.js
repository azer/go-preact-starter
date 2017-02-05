import { h, Component } from 'preact'

export default class Homepage extends Component {
  render () {
    return (
      <div class="homepage">
        <h1>Hi {this.props.matches.name || "there"}</h1>
        Wow, It worked!
      </div>
    )
  }
}
