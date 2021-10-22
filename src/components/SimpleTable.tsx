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
  TableSortLabel,
  Tooltip,
} from "@mui/material";

import { Log } from "../store/dashboard/types";

const orders = [
  {
    id: 1,
    ref: "CDD1049",
    amount: 30.5,
    customer: {
      name: "Ekaterina Tankova",
    },
    createdAt: 1555016400000,
    status: "pending",
  },
  {
    id: 2,
    ref: "CDD1048",
    amount: 25.1,
    customer: {
      name: "Cao Yu",
    },
    createdAt: 1555016400000,
    status: "delivered",
  },
  {
    id: 3,
    ref: "CDD1047",
    amount: 10.99,
    customer: {
      name: "Alexa Richardson",
    },
    createdAt: 1554930000000,
    status: "refunded",
  },
  {
    id: 4,
    ref: "CDD1046",
    amount: 96.43,
    customer: {
      name: "Anje Keizer",
    },
    createdAt: 1554757200000,
    status: "pending",
  },
  {
    id: 5,
    ref: "CDD1045",
    amount: 32.54,
    customer: {
      name: "Clarke Gillebert",
    },
    createdAt: 1554670800000,
    status: "delivered",
  },
  {
    id: 6,
    ref: "CDD1044",
    amount: 16.76,
    customer: {
      name: "Adam Denisov",
    },
    createdAt: 1554670800000,
    status: "delivered",
  },
];
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
