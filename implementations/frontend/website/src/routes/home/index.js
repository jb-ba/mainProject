import { h, Component } from "preact";
import Card from "preact-material-components/Card";
import "preact-material-components/Card/style.css";
import "preact-material-components/Button/style.css";
import style from "./style";

export default class Home extends Component {
  state = {
    btnAction: "Turn On",
    device: {
      building: 0,
      room: 0,
      label: "Front",
      ledOn: false
    }
  };

  log = () => {
    console.log("log test");
  };

  render() {
    return (
      <div class={`${style.home} page`}>
        <h1>Home route</h1>
        <Card>
          <div class={style.cardHeader}>
            <h2 class=" mdc-typography--title">Light</h2>
            <div class={[" mdc-typography--caption", style.bold].join(" ")}>
              Building {this.state.device.building}, Room{" "}
              {this.state.device.room}, Label {this.state.device.label}
            </div>
          </div>
          <div class={style.cardBody}>
            The device only offers the ability to be turned on and off at the
            moment.
          </div>
          <Card.Actions>
            <Card.ActionButton onClick={this.log}>
              {this.state.btnAction}
            </Card.ActionButton>
          </Card.Actions>
        </Card>
      </div>
    );
  }
}
