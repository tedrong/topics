import React, { Suspense, useEffect } from "react";
import {
  Router,
  Route,
  Switch,
  Redirect,
  useHistory,
  useLocation,
  Link,
} from "react-router-dom";
import styled from "styled-components";
import { useDispatch, useSelector } from "react-redux";
import {
  TextField,
  fieldToTextField,
  TextFieldProps,
} from "formik-material-ui";
import {
  Badge,
  Button,
  Container,
  makeStyles,
  TextField as MuiTextField,
} from "@material-ui/core";

import {
  FaBell,
  FaSignOutAlt,
  // FaUsers,
  // FaDollyFlatbed,
  // FaShippingFast,
  // FaMicrochip,
  // FaMobileAlt,
  // FaDollarSign,
  // FaChartBar,
  // FaClipboardList,
  // FaUser,
} from "react-icons/fa";

import Pending from "../component/pending";
import NotFound from "../component/notFound";
const Welcome = React.lazy(() => import("./welcome"));

const Layout = styled.div`
  background-color: #f8f9fa;
  display: flex;
  position: relative;
  flex-grow: 1;
  flex-direction: column;
  overflow-y: auto;
  height: 100vh;
`;
const Content = styled.div`
  padding: 0;
  // margin: 1vw;
  height: 100vh;
  overflow: auto;
  border: 1px solid lightgray;
  background-color: white;
`;

const StyledNavBar = styled.div`
  background-color: rgba(74, 74, 74);
  display: table;
  width: auto;
  height: 6%;
  // height: 1vh;
  // margin: 1vh;
`;
const NavBarItem = styled.div`
  display: table-cell;
  vertical-align: middle;
  text-align: right;
  padding-right: 1em;
`;
const CollapseIcon = styled.div`
  display: table-cell;
  vertical-align: middle;
  cursor: pointer;
  width: 2%;
  font-size: 18px;
  &:hover {
    color: DeepSkyBlue;
  }
`;

export default function Frame() {
  return (
    <Layout>
      <NavBar />
      <Content>
        <PrivateRouting />
      </Content>
    </Layout>
  );
}

function NavBar() {
  const handleLogout = () => {};
  return (
    <StyledNavBar>
      <CollapseIcon></CollapseIcon>
      <NavBarItem>
        <Button style={{ color: "white" }} onClick={handleLogout}>
          <FaSignOutAlt />
        </Button>
        <Badge badgeContent={4} color="primary">
          <FaBell />
        </Badge>
      </NavBarItem>
    </StyledNavBar>
  );
}

function PrivateRouting() {
  return (
    <Router history={useHistory()}>
      <Suspense fallback={<Pending />}>
        <Switch>
          <Route exact path={"/home/welcome"} component={Welcome} />
          <Route component={NotFound} />
        </Switch>
      </Suspense>
    </Router>
  );
}
