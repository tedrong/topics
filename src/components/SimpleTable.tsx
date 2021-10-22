import moment from "moment";
import {
  Box,
  Button,
  Card,
  CardHeader,
  Chip,
  Divider,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from "@mui/material";

import { Log } from "../store/dashboard/types";

interface Prop {
  list: Log[];
}

export default function LatestOrders(prop: Prop) {
  return (
    <Card>
      <CardHeader title="Latest Logs" />
      <Divider />
      <Box sx={{ minWidth: 800 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Date</TableCell>
              <TableCell>Message</TableCell>
              {/* <TableCell sortDirection="desc">
                <Tooltip enterDelay={300} title="Sort">
                  <TableSortLabel active direction="desc">
                    Date
                  </TableSortLabel>
                </Tooltip>
              </TableCell> */}
              <TableCell>Level</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {prop.list.map((log, index) => (
              <TableRow hover key={index}>
                <TableCell>
                  {moment(log.time * 1000).format("YYYY/MM/DD")}
                </TableCell>
                <TableCell>{log.message}</TableCell>
                <TableCell>{levelTag(log.level)}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Box>
      <Box
        sx={{
          display: "flex",
          justifyContent: "flex-end",
          p: 2,
        }}
      >
        <Button
          color="primary"
          //   endIcon={<ArrowRightIcon />}
          size="small"
          variant="text"
        >
          View all
        </Button>
      </Box>
    </Card>
  );
}

function levelTag(level: string) {
  switch (level) {
    case "info":
      return <Chip color="info" label={"Info"} size="small" />;
    case "warning":
      return <Chip color="warning" label={"Warn"} size="small" />;
    case "error":
      return <Chip color="error" label={"Error"} size="small" />;
    default:
      return <Chip color="primary" label={"Default"} size="small" />;
  }
}
