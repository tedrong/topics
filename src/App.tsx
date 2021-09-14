import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";
import Login from "./Login";
import Frame from "./route/frame";
import mockup from "./store/mockData";

export default function App() {
  if ("demo" === process.env.REACT_APP_STAGE) {
    mockup();
  }
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/">
          <Redirect to="/login" />
        </Route>
        <Route path="/login" component={Login} />
        <Route path="/home" component={Frame} />
      </Switch>
    </BrowserRouter>
  );
}
