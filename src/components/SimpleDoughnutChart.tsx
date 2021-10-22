import { Doughnut } from "react-chartjs-2";
import { ChartOptions } from "chart.js";
import {
  Box,
  Card,
  CardContent,
  CardHeader,
  Divider,
  Typography,
  colors,
  useTheme,
} from "@mui/material";
import { BsLaptop, BsTablet, BsPhone } from "react-icons/bs";
// import LaptopMacIcon from '@material-ui/icons/LaptopMac';
// import PhoneIcon from '@material-ui/icons/Phone';
// import TabletIcon from '@material-ui/icons/Tablet';

interface Prop {
  desktop: number;
  mobile: number;
}

export default function SimpleDoughnutChart(prop: Prop) {
  const theme = useTheme();

  const data = {
    datasets: [
      {
        data: [prop.desktop, prop.mobile],
        backgroundColor: [colors.indigo[500], colors.orange[600]],
        borderWidth: 8,
        borderColor: colors.common.white,
        hoverBorderColor: colors.common.white,
      },
    ],
    labels: ["Desktop", "Mobile"],
  };

  const options: ChartOptions = {
    animation: false,
    // cutoutPercentage: 80,
    layout: { padding: 0 },
    maintainAspectRatio: false,
    responsive: true,
    plugins: {
      legend: {
        display: false,
      },
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

  const devices = [
    {
      title: "Desktop",
      value: prop.desktop,
      icon: BsLaptop,
      color: colors.indigo[500],
    },
    {
      title: "Mobile",
      value: prop.mobile,
      icon: BsPhone,
      color: colors.orange[600],
    },
  ];

  return (
    <Card sx={{ height: "100%" }}>
      <CardHeader title="Traffic by Device" />
      <Divider />
      <CardContent>
        <Box
          sx={{
            height: 300,
            position: "relative",
          }}
        >
          <Doughnut data={data} options={options} />
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "center",
            pt: 2,
          }}
        >
          {devices.map(({ color, icon: Icon, title, value }) => (
            <Box
              key={title}
              sx={{
                p: 1,
                textAlign: "center",
              }}
            >
              <Icon color="action" />
              <Typography color="textPrimary" variant="body1">
                {title}
              </Typography>
              <Typography style={{ color }} variant="h2">
                {value}%
              </Typography>
            </Box>
          ))}
        </Box>
      </CardContent>
    </Card>
  );
}
