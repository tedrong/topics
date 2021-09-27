import {
  Avatar,
  Box,
  Card,
  CardContent,
  Grid,
  LinearProgress,
  Typography,
} from "@mui/material";
import { IconType } from "react-icons";

interface Prop {
  title: string;
  value: number;
  icon: IconType;
  color: string;
}

export default function TasksProgress(prop: Prop) {
  let Icon = prop.icon;
  return (
    <Card sx={{ height: "100%" }}>
      <CardContent>
        <Grid container spacing={3} sx={{ justifyContent: "space-between" }}>
          <Grid item>
            <Typography color="textSecondary" gutterBottom variant="h6">
              {prop.title}
            </Typography>
            <Typography color="textPrimary" variant="h3">
              {prop.value + "%"}
            </Typography>
          </Grid>
          <Grid item>
            <Avatar
              sx={{
                backgroundColor: prop.color,
                height: 56,
                width: 56,
              }}
            >
              <Icon />
            </Avatar>
          </Grid>
        </Grid>
        <Box sx={{ pt: 3 }}>
          <LinearProgress value={prop.value} variant="determinate" />
        </Box>
      </CardContent>
    </Card>
  );
}
