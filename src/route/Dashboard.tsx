import React from "react";
import { useDispatch, useSelector } from "react-redux";

import { Info, InfoHistory, ClientType, Log } from "../store/dashboard/types";
import {
  fetchInfoRequest,
  fetchInfoHistoryRequest,
  fetchClientTypeRequest,
  fetchLogRequest,
} from "../store/dashboard/actions";
import {
  getInfoSelector,
  getInfoPendingSelector,
  getInfoHistorySelector,
  getInfoHistoryPendingSelector,
  getClientTypeSelector,
  getClientTypePendingSelector,
  getLogSelector,
  getLogPendingSelector,
} from "../store/dashboard/selectors";

import { Helmet } from "react-helmet";
import { Box, Container, Grid } from "@mui/material";
import { red, orange, yellow, indigo } from "@mui/material/colors";
import { GiCpu } from "react-icons/gi";
import { FaMemory, FaClock } from "react-icons/fa";
import { FiHardDrive } from "react-icons/fi";

import Pending from "../components/Pending";
import ProgressCard from "../components/ProgressCard";
import InfoCard from "../components/InfoCard";
import SimpleBarChart from "../components/SimpleBarChart";
import SimpleDoughnutChart from "../components/SimpleDoughnutChart";
import SimpleTable from "../components/SimpleTable";
import { Epoch2Duration } from "../util/common";

export default function Dashboard() {
  const dispatch = useDispatch();

  var info: Info = useSelector(getInfoSelector);
  var infoPending = useSelector(getInfoPendingSelector);
  var infoHistory: InfoHistory = useSelector(getInfoHistorySelector);
  var infoHistoryPending = useSelector(getInfoHistoryPendingSelector);
  var clientType: ClientType = useSelector(getClientTypeSelector);
  var clientTypePending = useSelector(getClientTypePendingSelector);
  var log: Log[] = useSelector(getLogSelector);
  var logPending = useSelector(getLogPendingSelector);

  React.useEffect(() => {
    dispatch(fetchInfoRequest());
    dispatch(fetchInfoHistoryRequest({ amount: 7 }));
    dispatch(fetchClientTypeRequest());
    dispatch(fetchLogRequest({ amount: 10 }));
  }, [dispatch]);

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
              {(infoPending && <Pending />) || (
                <ProgressCard
                  title={"CPU Usage"}
                  value={+info.cpu}
                  icon={GiCpu}
                  color={red[600]}
                />
              )}
            </Grid>
            <Grid item lg={3} sm={6} xl={3} xs={12}>
              {(infoPending && <Pending />) || (
                <ProgressCard
                  title={"Memory Info"}
                  value={+info.memory}
                  icon={FaMemory}
                  color={orange[600]}
                />
              )}
            </Grid>
            <Grid item lg={3} sm={6} xl={3} xs={12}>
              {(infoPending && <Pending />) || (
                <ProgressCard
                  title={"Disk Status"}
                  value={+info.disk}
                  icon={FiHardDrive}
                  color={yellow[600]}
                />
              )}
            </Grid>
            <Grid item lg={3} sm={6} xl={3} xs={12}>
              {(infoPending && <Pending />) || (
                <InfoCard
                  title={"Boot Time"}
                  value={Epoch2Duration(info.bootTime)}
                  icon={FaClock}
                  color={indigo[600]}
                />
              )}
            </Grid>
            <Grid item lg={8} md={12} xl={9} xs={12}>
              {(infoHistoryPending && <Pending />) || (
                <SimpleBarChart
                  label={infoHistory.label}
                  cpu={infoHistory.cpu}
                  memory={infoHistory.memory}
                  disk={infoHistory.disk}
                />
              )}
            </Grid>
            <Grid item lg={4} md={6} xl={3} xs={12}>
              {(clientTypePending && <Pending />) || (
                <SimpleDoughnutChart
                  desktop={clientType.desktop}
                  mobile={clientType.mobile}
                />
              )}
            </Grid>
            <Grid item lg={12} sm={12} xl={12} xs={12}>
              {(logPending && <Pending />) || <SimpleTable list={log} />}
            </Grid>
          </Grid>
        </Container>
      </Box>
    </>
  );
}
