import axios from "axios";
import https from "https";
import fs from "fs";

export default function init() {
  axios.defaults.baseURL = process.env.REACT_APP_BASE_URL;
  // axios.defaults.headers.post["Content-Type"] = "application/json";
}

const Version = "v1";
export const API = {
  user: {
    login: Version + "/user/login",
  },
  dashboard: {
    info: Version + "/dashboard/system/info",
    infoHistory: Version + "/dashboard/system/info/history",
    clientType: Version + "/dashboard/system/client/type/percentage",
    log: Version + "/dashboard/system/log",
  },
};
