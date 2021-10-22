import React, { Suspense, useState } from "react";
import {
  Link,
  matchPath,
  Redirect,
  Router,
  Route,
  Switch,
  useHistory,
  useLocation,
} from "react-router-dom";
import styled from "styled-components";
import { useSelector } from "react-redux";
import {
  AppBar,
  Avatar,
  Badge,
  Box,
  Button,
  Divider,
  Drawer,
  Hidden,
  IconButton,
  List,
  ListItem,
  Toolbar,
  Typography,
} from "@mui/material";

import { IconType, IconContext } from "react-icons";
import { FaBell, FaSignOutAlt, FaBars, FaUser } from "react-icons/fa";
import { BsPieChartFill, BsXSquareFill } from "react-icons/bs";
import logo from "../material/leave.png";
import theme from "../theme";

import { getLoginSelector } from "../store/user/selectors";

import Pending from "../components/Pending";
import NotFound from "../components/NotFound";
const Welcome = React.lazy(() => import("./Welcome"));
const Dashboard = React.lazy(() => import("./Dashboard"));
const Account = React.lazy(() => import("./Account"));

const LayoutRoot = styled.div`
  display: flex;
  overflow: hidden;
  height: 100vh;
  background-color: #fafafa;
`;

const LayoutWrapper = styled("div")(() => ({
  display: "flex",
  flex: "1 1 auto",
  overflow: "hidden",
  paddingTop: 64,
  [theme.breakpoints.up("lg")]: {
    paddingLeft: 256,
  },
}));

const LayoutContainer = styled.div`
  display: flex;
  flex: 1 1 auto;
  overflow: hidden;
`;

const LayoutContent = styled.div`
  flex: 1 1 auto;
  height: 100%;
  overflow: auto;
`;

const StyledImg = styled.img`
  height: 3vh;
`;

interface menuStatus {
  mobile: boolean;
  setMobile: React.Dispatch<React.SetStateAction<boolean>>;
}
interface menuItem {
  href: string;
  icon: IconType;
  title: string;
}

export default function Frame() {
  const login = useSelector(getLoginSelector);
  const [MobileMode, setMobileMode] = useState(false);

  if (login.token.access_token !== "") {
    return (
      <LayoutRoot>
        <IconContext.Provider value={{ style: { marginRight: "10px" } }}>
          <NavBar mobile={MobileMode} setMobile={setMobileMode} />
          <SideMenu mobile={MobileMode} setMobile={setMobileMode} />
        </IconContext.Provider>
        <LayoutWrapper>
          <LayoutContainer>
            <LayoutContent>
              <PrivateRouting />
            </LayoutContent>
          </LayoutContainer>
        </LayoutWrapper>
      </LayoutRoot>
    );
  } else {
    return <Redirect to="/login" />;
  }
}

function NavBar(status: menuStatus) {
  const handleLogout = () => {
    console.log("logout");
  };
  return (
    <AppBar position="fixed">
      <Toolbar>
        <Link to="/home/welcome">
          <StyledImg src={logo} />
        </Link>
        <Box sx={{ flexGrow: 1 }} />
        <Hidden xlDown>
          <IconButton color="inherit">
            <Badge badgeContent={10} color="info" variant="dot"></Badge>
            <FaBell />
          </IconButton>
          <IconButton color="inherit" onClick={handleLogout}>
            <FaSignOutAlt />
          </IconButton>
        </Hidden>
        <Hidden xlUp>
          <IconButton
            color="inherit"
            onClick={() => {
              status.setMobile(true);
            }}
          >
            <FaBars />
          </IconButton>
        </Hidden>
      </Toolbar>
    </AppBar>
  );
}

function SideMenu(status: menuStatus) {
  const user = {
    avatar: "",
    jobTitle: "R&D",
    name: "Ted Rong",
  };
  const items = [
    {
      href: "/home/dashboard",
      icon: BsPieChartFill,
      title: "Dashboard",
    },
    {
      href: "/home/account",
      icon: FaUser,
      title: "Account",
    },
    {
      href: "/home/notfound",
      icon: BsXSquareFill,
      title: "NotFound",
    },
  ];
  const content = (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
      }}
    >
      <Box
        sx={{
          alignItems: "center",
          display: "flex",
          flexDirection: "column",
          p: 2,
        }}
      >
        <Avatar
          component={Link}
          src={user.avatar}
          sx={{
            cursor: "pointer",
            width: 64,
            height: 64,
          }}
          to="/app/account"
        />
        <Typography color="textPrimary" variant="h5">
          {user.name}
        </Typography>
        <Typography color="textSecondary" variant="body2">
          {user.jobTitle}
        </Typography>
      </Box>
      <Divider />
      <Box sx={{ p: 2 }}>
        <List>
          {items.map((item) => (
            <SideMenuItem
              href={item.href}
              key={item.title}
              title={item.title}
              icon={item.icon}
            />
          ))}
        </List>
      </Box>
    </Box>
  );
  return (
    <>
      <Hidden lgUp>
        <Drawer
          anchor="left"
          onClose={() => {
            status.setMobile(false);
          }}
          open={status.mobile}
          variant="temporary"
          PaperProps={{
            sx: {
              width: 256,
            },
          }}
        >
          {content}
        </Drawer>
      </Hidden>
      <Hidden lgDown>
        <Drawer
          anchor="left"
          open
          variant="persistent"
          PaperProps={{
            sx: {
              width: 256,
              top: 64,
              height: "calc(100% - 64px)",
            },
          }}
        >
          {content}
        </Drawer>
      </Hidden>
    </>
  );
}

function SideMenuItem(item: menuItem) {
  const location = useLocation();
  const active = item.href ? !!matchPath(item.href, location.pathname) : false;
  let Icon = item.icon;
  return (
    <ListItem
      disableGutters
      sx={{
        display: "flex",
        py: 0,
      }}
      // {...rest}
    >
      <Button
        component={Link}
        sx={{
          color: "text.secondary",
          fontWeight: "medium",
          justifyContent: "flex-start",
          letterSpacing: 0,
          py: 1.25,
          textTransform: "none",
          width: "100%",
          ...(active && {
            color: "primary.main",
          }),
          "& svg": {
            mr: 1,
          },
        }}
        to={item.href}
      >
        <Icon size="20" />
        <Typography>{item.title}</Typography>
      </Button>
    </ListItem>
  );
}

function PrivateRouting() {
  return (
    <Router history={useHistory()}>
      <Suspense fallback={<Pending />}>
        <Switch>
          <Route exact path={"/home/welcome"} component={Welcome} />
          <Route exact path={"/home/dashboard"} component={Dashboard} />
          <Route exact path={"/home/account"} component={Account} />
          <Route component={NotFound} />
        </Switch>
      </Suspense>
    </Router>
  );
}
