import { Helmet } from "react-helmet";
import { Box, Container, Grid } from "@mui/material";
import { red, orange, yellow, indigo } from "@mui/material/colors";
import { GiCpu } from "react-icons/gi";
import { FaMemory, FaClock } from "react-icons/fa";
import { FiHardDrive } from "react-icons/fi";

import ProgressCard from "../components/ProgressCard";
import InfoCard from "../components/InfoCard";
import SimpleBarChart from "../components/SimpleBarChart";
import SimpleDoughnutChart from "../components/SimpleDoughnutChart";
import SimpleTable from "../components/SimpleTable";

export default function Dashboard() {
  return (
    <>
      <Helmet>
        <title>Dashboard | Topics</title>
      </Helmet>
      <Box
        sx={{
          backgroundColor: "background.default",
          minHeight: "100%",
          py: 3,
        }}
      >
        <Container maxWidth={false}>
          <Grid container spacing={3}>
            <Grid item lg={3} sm={6} xl={3} xs={12}>
              <ProgressCard
                title={"CPU Usage"}
                value={30}
                icon={GiCpu}
                color={red[600]}
              />
            </Grid>
            <Grid item lg={3} sm={6} xl={3} xs={12}>
              <ProgressCard
                title={"Memory Info"}
                value={60}
                icon={FaMemory}
                color={orange[600]}
              />
            </Grid>
            <Grid item lg={3} sm={6} xl={3} xs={12}>
              <ProgressCard
                title={"Disk Status"}
                value={60}
                icon={FiHardDrive}
                color={yellow[600]}
              />
            </Grid>
            <Grid item lg={3} sm={6} xl={3} xs={12}>
              <InfoCard
                title={"Boot Time"}
                value={"7Days"}
                icon={FaClock}
                color={indigo[600]}
              />
            </Grid>
            <Grid item lg={8} md={12} xl={9} xs={12}>
              <SimpleBarChart />
            </Grid>
            <Grid item lg={4} md={6} xl={3} xs={12}>
              <SimpleDoughnutChart />
            </Grid>
            <Grid item lg={12} sm={12} xl={12} xs={12}>
              <SimpleTable />
            </Grid>
          </Grid>
        </Container>
      </Box>
    </>
  );
}
