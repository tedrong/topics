import axios from "axios";

export default function init() {
  axios.defaults.baseURL = process.env.REACT_APP_BASE_URL;
}

export function attachAuthToken(token: string) {
  if (token) {
    axios.defaults.headers.common["Authorization"] = "Bearer " + token;
  } else {
    axios.defaults.headers.common["Authorization"] = null;
  }
}

const Version = "v1";
export const API = {
  user: {
    login: Version + "/user/login",
    renew: Version + "/user/renew",
  },
  dashboard: {
    info: Version + "/dashboard/system/info",
    infoHistory: Version + "/dashboard/system/info/history",
    clientType: Version + "/dashboard/system/client/type/percentage",
    log: Version + "/dashboard/system/log",
  },
};
