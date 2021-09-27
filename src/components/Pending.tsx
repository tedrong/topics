import styled from "styled-components";
import { CircularProgress } from "@mui/material";

const Content = styled.div`
  display: table;
  margin: 27% auto 0 auto;
`;

const Wrap = styled.div`
  width: 10vw;
  text-align: center;
`;

const Text = styled.div`
  paddingtop: 10px;
`;

export default function Pending() {
  return (
    <Content>
      <Wrap>
        <CircularProgress />
        <Text>Loading</Text>
      </Wrap>
    </Content>
  );
}
