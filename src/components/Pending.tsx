import React, { useState, useEffect } from "react";
import { makeStyles } from "@mui/styles";
import { CircularProgress } from "@mui/material";

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
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <div className={classes.wrap}>
        <CircularProgress />
        <div className={classes.text}>Loading</div>
      </div>
    </div>
  );
}
