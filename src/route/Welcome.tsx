import styled from "styled-components";
import { Helmet } from "react-helmet";

const Content = styled.div`
  text-align: center;
  margin-top: 20%;
`;

const MainTitle = styled.div`
  color: #a9dfbf;
  font-size: 8vh;
`;

const SubTitle = styled.div`
  color: #145a32;
  font-size: 6vh;
  font-weight: bold;
`;

export default function Welcome() {
  return (
    <>
      <Helmet>
        <title>Welcome | Topics</title>
      </Helmet>
      <Content>
        <MainTitle>{"Welcome to Topics system"}</MainTitle>
        <SubTitle>{"Administration Web Server"}</SubTitle>
      </Content>
    </>
  );
}
