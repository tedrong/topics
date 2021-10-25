import moment from "moment";
import { Bar } from "react-chartjs-2";
import { ChartOptions } from "chart.js";
import {
  Box,
  Button,
  Card,
  CardContent,
  CardHeader,
  Divider,
  useTheme,
  colors,
} from "@mui/material";
import { BsFillCaretDownFill, BsFillCaretRightFill } from "react-icons/bs";

interface Prop {
  label: number[];
  cpu: number[];
  memory: number[];
  disk: number[];
}

export default function SimpleBarChart(prop: Prop) {
  const theme = useTheme();
  let label: string[] = [];
  if (prop.label) {
    prop.label.map((date) => label.push(moment(date * 1000).format("MM/DD")));
  }

  const data = {
    datasets: [
      {
        backgroundColor: colors.indigo[500],
        barPercentage: 0.5,
        barThickness: 12,
        borderRadius: 4,
        categoryPercentage: 0.5,
        data: prop.cpu,
        label: "cpu",
        maxBarThickness: 10,
      },
      {
        backgroundColor: colors.grey[200],
        barPercentage: 0.5,
        barThickness: 12,
        borderRadius: 4,
        categoryPercentage: 0.5,
        data: prop.memory,
        label: "memory",
        maxBarThickness: 10,
      },
      {
        backgroundColor: colors.blue[200],
        barPercentage: 0.5,
        barThickness: 12,
        borderRadius: 4,
        categoryPercentage: 0.5,
        data: prop.disk,
        label: "disk",
        maxBarThickness: 10,
      },
    ],
    labels: label,
  };

  return (
    <Card sx={{ height: "100%" }}>
      <CardHeader
        action={
          <Button endIcon={<BsFillCaretDownFill />} size="small" variant="text">
            Last 7 days
          </Button>
        }
        title="Consumption History"
      />
      <Divider />
      <CardContent>
        <Box
          sx={{
            height: 400,
            position: "relative",
          }}
        >
          <Bar data={data} />
        </Box>
      </CardContent>
      <Divider />
      <Box
        sx={{
          display: "flex",
          justifyContent: "flex-end",
          p: 2,
        }}
      >
        <Button
          color="primary"
          endIcon={<BsFillCaretRightFill />}
          size="small"
          variant="text"
        >
          Overview
        </Button>
      </Box>
    </Card>
  );
}
