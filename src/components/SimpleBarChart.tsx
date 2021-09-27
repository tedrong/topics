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

export default function SimpleBarChart() {
  const theme = useTheme();

  const data = {
    datasets: [
      {
        backgroundColor: colors.indigo[500],
        barPercentage: 0.5,
        barThickness: 12,
        borderRadius: 4,
        categoryPercentage: 0.5,
        data: [18, 5, 19, 27, 29, 19, 20],
        label: "This year",
        maxBarThickness: 10,
      },
      {
        backgroundColor: colors.grey[200],
        barPercentage: 0.5,
        barThickness: 12,
        borderRadius: 4,
        categoryPercentage: 0.5,
        data: [11, 20, 12, 29, 30, 25, 13],
        label: "Last year",
        maxBarThickness: 10,
      },
    ],
    labels: ["1 Aug", "2 Aug", "3 Aug", "4 Aug", "5 Aug", "6 Aug"],
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
        title="API Fetch Count"
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
