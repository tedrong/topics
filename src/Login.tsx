import React from "react";
import { Helmet } from "react-helmet";
import { Redirect, Link as RouterLink } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import {
  getTokenSelector,
  getLoginErrorSelector,
} from "./store/auth/selectors";
import { Token } from "./store/auth/types";
import { fetchLoginRequest } from "./store/auth/actions";

import * as Yup from "yup";
import { Formik } from "formik";
import {
  Avatar,
  Box,
  Button,
  Container,
  CssBaseline,
  Link,
  TextField,
  Typography,
} from "@mui/material";

import styled from "styled-components";
import logo from "./material/landscape.png";

const StyledBox = styled(Box)`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  min-width: 300px;
  background-color: #fafafa;
`;

const StyledAvatar = styled(Avatar)`
  margin: auto;
  border-radius: 0%;
  background-color: white;
`;

const StyledForm = styled.form`
  padding: 10%;
  align-items: center;
  text-align: center;
  border: 1px lightgray solid;
  background-color: white;
`;

export default function LoginBox() {
  const dispatch = useDispatch();
  const token: Token = useSelector(getTokenSelector);
  const error: string | null = useSelector(getLoginErrorSelector);

  if (error === null && token.access_token !== "") {
    return <Redirect to="/home/welcome" />;
  }

  return (
    <React.Fragment>
      <Helmet>
        <title>Login | Topics</title>
      </Helmet>
      <StyledBox>
        <Container maxWidth="sm">
          <Formik
            initialValues={{
              email: "",
              password: "",
            }}
            validationSchema={Yup.object().shape({
              email: Yup.string()
                .email("Must be a valid email")
                .max(255)
                .required("Email is required"),
              password: Yup.string().max(255).required("Password is required"),
            })}
            onSubmit={(values, { setSubmitting }) => {
              dispatch(fetchLoginRequest(values));
              setSubmitting(false);
            }}
          >
            {({
              errors,
              handleBlur,
              handleChange,
              handleSubmit,
              isSubmitting,
              touched,
              values,
            }) => (
              <StyledForm onSubmit={handleSubmit}>
                <Box sx={{ mb: 3 }}>
                  <StyledAvatar
                    src={logo}
                    alt="logo"
                    variant={"square"}
                    sx={{ width: "8vw", height: "auto" }}
                  />
                  <CssBaseline />
                  <Typography color="textPrimary" variant="h2">
                    Topics
                  </Typography>
                  <Typography
                    color="textSecondary"
                    gutterBottom
                    variant="body2"
                  >
                    Login in on the internal platform
                  </Typography>
                </Box>
                <TextField
                  error={Boolean(touched.email && errors.email)}
                  fullWidth
                  helperText={touched.email && errors.email}
                  label="Email Address"
                  margin="normal"
                  name="email"
                  onBlur={handleBlur}
                  onChange={handleChange}
                  type="email"
                  value={values.email}
                  variant="outlined"
                />
                <TextField
                  error={Boolean(touched.password && errors.password)}
                  fullWidth
                  helperText={touched.password && errors.password}
                  label="Password"
                  margin="normal"
                  name="password"
                  onBlur={handleBlur}
                  onChange={handleChange}
                  type="password"
                  value={values.password}
                  variant="outlined"
                />
                <Box sx={{ py: 2 }}>
                  <Button
                    color="success"
                    disabled={isSubmitting}
                    fullWidth
                    size="large"
                    type="submit"
                    variant="contained"
                  >
                    Login
                  </Button>
                </Box>
                <Typography color="textSecondary" variant="body1">
                  Don&apos;t have an account?{" "}
                  <Link component={RouterLink} to="/register" underline="hover">
                    Sign up
                  </Link>
                </Typography>
              </StyledForm>
            )}
          </Formik>
        </Container>
      </StyledBox>
    </React.Fragment>
  );
}