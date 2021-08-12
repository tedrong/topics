import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";
import Login from "./Login";

function App() {
  return (
    <BrowserRouter>
      <Switch>
        <Route exact path="/">
          <Redirect to="/login" />
        </Route>
        <Route path="/login" component={Login} />
        <Route path="/home" />
      </Switch>
    </BrowserRouter>
  );
}

export default App;
