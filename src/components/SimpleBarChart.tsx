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
  prop.label.map((date) => {
    label.push(moment(date * 1000).format("MM/DD"));
  });
  
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

  const options: ChartOptions = {
    animation: false,
    // cornerRadius: 20,
    layout: { padding: 0 },
    // legend: { display: false },
    maintainAspectRatio: false,
    responsive: true,
    scales: {
      x: {
        ticks: {
          color: theme.palette.text.secondary,
        },
        grid: {
          display: false,
          drawBorder: false,
        },
      },
      y: {
        beginAtZero: true,
        min: 0,
        ticks: {
          color: theme.palette.text.secondary,
        },
        grid: {
          borderDash: [2],
          //   borderDashOffset: [2],
          color: theme.palette.divider,
          drawBorder: false,
          //   zeroLineBorderDash: [2],
          //   zeroLineBorderDashOffset: [2],
          //   zeroLineColor: theme.palette.divider,
        },
      },
    },
    plugins: {
      tooltip: {
        backgroundColor: theme.palette.background.paper,
        // bodyFontColor: theme.palette.text.secondary,
        borderColor: theme.palette.divider,
        borderWidth: 1,
        enabled: true,
        // footerFontColor: theme.palette.text.secondary,
        intersect: false,
        mode: "index",
        // titleFontColor: theme.palette.text.primary,
      },
    },
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
          <Bar data={data} options={options} />
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
