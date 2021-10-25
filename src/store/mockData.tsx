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
      first_name: "demo_first",
      last_name: "demo_last",
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
    label: [1635044801, 1635044801],
  });
  mock.onGet(API.dashboard.clientType).reply(200, {
    desktop: 70,
    mobile: 30,
  });
  mock.onGet(new RegExp(`${API.dashboard.log}/*`)).reply(200, [
    {
      level: "debug",
      time: 1635044993,
      message: "Table migrate successfully in DB:internal",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "debug",
      time: 1635044993,
      message: "Table migrate successfully in DB:content",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "debug",
      time: 1635044993,
      message: "Table migrate successfully in DB:googleTrends",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "debug",
      time: 1635044993,
      message: "Table migrate successfully in DB:stock",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "info",
      time: 1635044993,
      message: "logging configured",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "debug",
      time: 1635044945,
      message: "Table migrate successfully in DB:internal",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "debug",
      time: 1635044945,
      message: "Table migrate successfully in DB:content",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "debug",
      time: 1635044945,
      message: "Table migrate successfully in DB:googleTrends",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "debug",
      time: 1635044945,
      message: "Table migrate successfully in DB:stock",
      type: "",
      server: "",
      cronjob: "",
    },
    {
      level: "info",
      time: 1635044944,
      message: "logging configured",
      type: "",
      server: "",
      cronjob: "",
    },
  ]);
  mock.onPut(new RegExp(`${API.user.renew}/*`)).reply(200, {
    user: {
      ID: 2,
      CreatedAt: "2021-10-24T09:28:02.097789877+08:00",
      UpdatedAt: "2021-10-24T09:28:02.097789877+08:00",
      DeletedAt: null,
      UUID: "",
      first_name: "renew_first_demo",
      last_name: "renew_last_demo",
      email: "",
    },
  });
  mock.onPost(API.auth.refresh).reply(200, {
    access_token:
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6ImNmYzM3ZDE0LWRkNjgtNGM1Zi04ZjRlLTIyM2U4NzEyZDQ2ZCIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYzNTA1MjQxOSwidXNlcl9pZCI6MX0.bINqV3SkYFysM9XnqlusEaqSCXEM_JSfSPiFlz1-sjw",
    refresh_token:
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzU2NTYzMTksInJlZnJlc2hfdXVpZCI6IjczNWEwZjY4LTNmOTUtNDU2Zi04MDAzLTFhNmUzYzY0YTUwOCIsInVzZXJfaWQiOjF9.SwOhj9qXnSOE_t4s-HlLKJruzfhjBhwM8mNLWbfZPt0",
  });
}
