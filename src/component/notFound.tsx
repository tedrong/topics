import React from "react";
import { Alert, AlertTitle } from "@material-ui/lab";

export default function NotFound() {
  return (
    <Alert severity="error">
      <AlertTitle>Error</AlertTitle>
      <strong>Page Not Found!</strong>
    </Alert>
  );
}
