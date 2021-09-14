import React, { useState, useEffect } from "react";
import { makeStyles } from "@material-ui/core/styles";
import CircularProgress from "@material-ui/core/CircularProgress";

const useStyles = makeStyles((theme) => ({
  root: {
    display: "table",
    margin: "27% auto 0 auto",
  },
  wrap: {
    width: "10vw",
    textAlign: "center",
  },
  text: {
    paddingTop: "10px",
  },
}));

export default function Pending() {
  const [timer, setTimer] = useState("");
  const classes = useStyles();

  useEffect(() => {
    let interval: null | ReturnType<typeof setTimeout> = null;
    interval = setInterval(() => {
      setTimer((timer) => timer + ".");
    }, 500);
    return () => clearInterval();
  }, [timer]);
  if (timer.length > 3) {
    setTimer("");
  }

  return (
    <div className={classes.root}>
      <div className={classes.wrap}>
        <CircularProgress />
        <div className={classes.text}>Loading{timer}</div>
      </div>
    </div>
  );
}
