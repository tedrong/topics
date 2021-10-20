import axios from "axios";
import MockAdapter from "axios-mock-adapter";
import { API } from "./api";

export default function mockup() {
  var mock = new MockAdapter(axios, { delayResponse: 2000 });
  mock.onPost(API.user.login).reply(200, {
    message: "Successfully logged in",
    token: {
      access_token:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjBjMzVmMmE5LTQzZWMtNDc4My04OWIzLWE5NTQ5ZjJkNzIzOSIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYzMTU5OTUxOSwidXNlcl9pZCI6M30.SisvOqAw37PCS0AdTTJPFVa2x5ObwLFAx4oWMt96BKo",
      refresh_token:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzIyMDM0MTksInJlZnJlc2hfdXVpZCI6IjcxNThjZjVjLTFiZDAtNDNlNS1hNzA5LTcwOGE3MjNlOWI1OSIsInVzZXJfaWQiOjN9.Th_CD3wmS-QgJgEDAMnjfA1iccBjQuKnXjhy9JODYgg",
    },
    user: {
      ID: 3,
      CreatedAt: "2021-09-01T19:47:54.098209+08:00",
      UpdatedAt: "2021-09-01T19:47:54.098209+08:00",
      DeletedAt: null,
      UUID: "6835df53-7e4b-476e-bd22-e068a32f1956",
      name: "testing",
      email: "test@test.com",
    },
  });
  mock.onGet(API.dashboard.info).reply(200, {
    cpu: "45.69",
    memory: "75.84",
    disk: "15.54",
    bootTime: 12312373,
  });
  mock.onGet(new RegExp(`${API.dashboard.infoHistory}/*`)).reply(200, {
    cpu: ["42.23", "53.28"],
    memory: ["65.38", "76.03"],
    disk: ["15.52", "15.54"],
    label: [
      "2021-10-17 09:20:49.329041 +0800 CST",
      "2021-10-18 09:31:56.641926 +0800 CST",
    ],
  });
}
