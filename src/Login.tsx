import React, { useState, useEffect } from "react";
import { Redirect } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import {
  getPendingSelector,
  getLoginSelector,
  getErrorSelector,
} from "./store/login/selectors";
import { fetchLoginRequest } from "./store/login/actions";

import { Formik, Form, Field } from "formik";
import {
  TextField,
  fieldToTextField,
  TextFieldProps,
} from "formik-material-ui";
import {
  Avatar,
  Button,
  Container,
  CssBaseline,
  InputAdornment,
  IconButton,
  LinearProgress,
  makeStyles,
  Paper,
  TextField as MuiTextField,
  Typography,
} from "@material-ui/core";
import { BsFillEyeFill, BsFillEyeSlashFill } from "react-icons/bs";
import logo from "./material/landscape.png";

const useStyles = makeStyles(() => ({
  container: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    height: "100vh",
    minWidth: "300px",
  },
  paper: {
    width: "32vw",
    padding: "5%",
    textAlign: "center",
  },
  avatar: {
    margin: "auto",
    width: "60%",
    height: "auto",
    borderRadius: "0%",
    backgroundColor: "white",
  },
  submitBtn: {
    height: "5vh",
    marginTop: "1vh",
  },
}));

interface Values {
  email: string;
  password: string;
}

export default function LoginBox() {
  const classes = useStyles();
  const dispatch = useDispatch();
  const pending = useSelector(getPendingSelector);
  const login = useSelector(getLoginSelector);
  const error = useSelector(getErrorSelector);

  if (login.access_token != "") {
    return <Redirect to="/home/welcome" />;
  }

  return (
    <React.Fragment>
      <Container className={classes.container} maxWidth="md">
        <Paper className={classes.paper} variant="outlined">
          <CssBaseline />
          <Avatar className={classes.avatar} src={logo} alt="logo" />
          <Typography component="h6" variant="h5">
            Frontend
          </Typography>
          <Formik
            initialValues={{
              email: "",
              password: "",
            }}
            validate={(values) => {
              const errors: Partial<Values> = {};
              if (!values.email) {
                errors.email = "Required";
              } else if (
                !/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)
              ) {
                errors.email = "Invalid email address";
              }
              return errors;
            }}
            onSubmit={(values, { setSubmitting }) => {
              dispatch(fetchLoginRequest());
              setSubmitting(false);
            }}
          >
            {({ submitForm, isSubmitting }) => (
              <Form>
                <Field
                  fullWidth
                  autoFocus
                  component={TextField}
                  variant="outlined"
                  margin="normal"
                  id="email"
                  name="email"
                  type="email"
                  label="Email"
                />
                <br />
                <Field
                  fullWidth
                  autoFocus
                  component={PasswordTextField}
                  variant="outlined"
                  margin="normal"
                  id="password"
                  name="password"
                  type="password"
                  label="Password"
                  autoComplete="current-password"
                />
                {pending && <LinearProgress />}
                <br />
                <Button
                  className={classes.submitBtn}
                  fullWidth
                  variant="contained"
                  color="primary"
                  disabled={pending}
                  onClick={submitForm}
                >
                  Submit
                </Button>
              </Form>
            )}
          </Formik>
        </Paper>
      </Container>
    </React.Fragment>
  );
}

function PasswordTextField(props: TextFieldProps) {
  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => {
    setShowPassword(!showPassword);
  };
  const handleMouseDownPassword = (event: React.SyntheticEvent) => {
    event.preventDefault();
  };
  return (
    <MuiTextField
      {...fieldToTextField(props)}
      type={showPassword ? "text" : "password"}
      InputProps={{
        endAdornment: (
          <InputAdornment position="end">
            <IconButton
              aria-label="toggle password visibility"
              onClick={handleClickShowPassword}
              onMouseDown={handleMouseDownPassword}
            >
              {showPassword ? <BsFillEyeFill /> : <BsFillEyeSlashFill />}
            </IconButton>
          </InputAdornment>
        ),
      }}
    />
  );
}
