import axios from "axios";
import { Token } from "./auth/types";
import { CookieStorage } from "cookie-storage";

export default function init() {
  axios.defaults.baseURL = process.env.REACT_APP_BASE_URL;
}

export function attachAuthToken(token: Token) {
  const cookieStorage = new CookieStorage();
  cookieStorage.setItem("refreshToken", token.refresh_token);
  axios.defaults.headers.common["Authorization"] =
    "Bearer " + token.access_token;
  axios.interceptors.request.use(
    function (response) {
      return response;
    },
    function (error) {
      if (error.response) {
        switch (error.response.status) {
          case 401:
            if (error.config.url !== API.auth.refresh) {
              const originalRequest = error.config;
              let refreshToken = cookieStorage.getItem("refreshToken");
              return axios
                .post(API.auth.refresh, {
                  refresh_token: refreshToken,
                })
                .then((response) => {
                  cookieStorage.setItem(
                    "refreshToken",
                    response.data.refresh_token
                  );
                  axios.defaults.headers.common["Authorization"] =
                    "Bearer " + response.data.access_token;
                  originalRequest.headers.Authorization =
                    "Bearer " + response.data.access_token;
                  return axios(originalRequest);
                })
                .catch(() => {
                  window.location.href = "/login";
                  return Promise.reject(error);
                });
            }
        }
      }
      return Promise.reject(error);
    }
  );
}

const Version = "v1";
export const API = {
  auth: {
    refresh: Version + "/token/refresh",
  },
  user: {
    login: Version + "/user/login",
    renew: Version + "/user/renew",
    logout: Version + "/user/logout",
  },
  dashboard: {
    info: Version + "/dashboard/system/info",
    infoHistory: Version + "/dashboard/system/info/history",
    clientType: Version + "/dashboard/system/client/type/percentage",
    log: Version + "/dashboard/system/log",
  },
};
