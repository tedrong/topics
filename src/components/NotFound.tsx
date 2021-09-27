import { Alert, AlertTitle } from "@mui/material";

export default function NotFound() {
  return (
    <Alert severity="error">
      <AlertTitle>Error</AlertTitle>
      <strong>Page Not Found!</strong>
    </Alert>
  );
}
