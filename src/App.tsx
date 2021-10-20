import { Route, Switch, Redirect } from "react-router-dom";
import { ThemeProvider } from "@mui/material/styles";

import theme from "./theme";
import mockup from "./store/mockData";
import init from "./store/api";

import Login from "./Login";
import Register from "./route/Register";
import Frame from "./route/Frame";

export default function App() {
  if ("demo" === process.env.REACT_APP_STAGE) {
    mockup();
  } else {
    init();
  }
  return (
    <ThemeProvider theme={theme}>
      <Switch>
        <Route exact path="/">
          <Redirect to="/login" />
        </Route>
        <Route path="/login" component={Login} />
        <Route path="/register" component={Register} />S
        <Route path="/home" component={Frame} />
      </Switch>
    </ThemeProvider>
  );
}
