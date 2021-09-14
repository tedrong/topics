import React from "react";
import { useDispatch } from "react-redux";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";

const useStyles = makeStyles({
  root: {
    textAlign: "center",
    marginTop: "20%",
  },
  mainTitle: {
    color: "#A3B3C2",
    fontSize: "4vw",
  },
  subTitle: {
    color: "#1D4157",
    fontSize: "3vw",
    fontWeight: "bold",
  },
});

export default function Welcome() {
  const dispatch = useDispatch();
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <div className={classes.mainTitle}>{"Welcome"}</div>
      <Typography className={classes.subTitle}>{"Topics"}</Typography>
    </div>
  );
}
