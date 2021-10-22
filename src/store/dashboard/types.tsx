import { FETCH_LOGIN_FAILURE } from "../user/actionTypes";
import {
  FETCH_INFO_REQUEST,
  FETCH_INFO_SUCCESS,
  FETCH_INFO_FAILURE,
  FETCH_INFO_HISTORY_REQUEST,
  FETCH_INFO_HISTORY_SUCCESS,
  FETCH_INFO_HISTORY_FAILURE,
  FETCH_CLIENT_TYPE_REQUEST,
  FETCH_CLIENT_TYPE_SUCCESS,
  FETCH_CLIENT_TYPE_FAILURE,
  FETCH_LOG_REQUEST,
  FETCH_LOG_SUCCESS,
  FETCH_LOG_FAILURE,
} from "./actionTypes";

export interface Info {
  cpu: string;
  memory: string;
  disk: string;
  bootTime: number;
}

export interface InfoHistory {
  cpu: number[];
  memory: number[];
  disk: number[];
  label: number[];
}

export interface ClientType {
  desktop: number;
  mobile: number;
}

export interface Log {
  level: string;
  time: number;
  message: string;
  type: string;
}

export interface InfoPayload {
  pending: boolean;
  data: Info;
  error: string | null;
}

export interface InfoHistoryPayload {
  pending: boolean;
  data: InfoHistory;
  error: string | null;
}

export interface ClientTypePayload {
  pending: boolean;
  data: ClientType;
  error: string | null;
}

export interface LogPayload {
  pending: boolean;
  data: Log[];
  error: string | null;
}

export interface DashboardState {
  info: InfoPayload;
  infoHistory: InfoHistoryPayload;
  clientType: ClientTypePayload;
  log: LogPayload;
}

export interface FetchInfoHistoryRequestPayload {
  amount: number;
}

export interface FetchLogRequestPayload {
  amount: number;
}

export interface FetchInfoSuccessPayload {
  data: Info;
}

export interface FetchInfoHistorySuccessPayload {
  data: InfoHistory;
}

export interface FetchClientTypeSuccessPayload {
  data: ClientType;
}

export interface FetchLogSuccessPayload {
  data: Log;
}

export interface FetchInfoFailurePayload {
  error: string;
}

export interface FetchInfoHistoryFailurePayload {
  error: string;
}

export interface FetchClientTypeFailurePayload {
  error: string;
}

export interface FetchLogFailurePayload {
  error: string;
}

export interface FetchInfoRequest {
  type: typeof FETCH_INFO_REQUEST;
}

export interface FetchInfoHistoryRequest {
  type: typeof FETCH_INFO_HISTORY_REQUEST;
  payload: FetchInfoHistoryRequestPayload;
}

export interface FetchClientTypeRequest {
  type: typeof FETCH_CLIENT_TYPE_REQUEST;
}

export interface FetchLogRequest {
  type: typeof FETCH_LOG_REQUEST;
  payload: FetchLogRequestPayload;
}

export type FetchInfoSuccess = {
  type: typeof FETCH_INFO_SUCCESS;
  payload: FetchInfoSuccessPayload;
};

export type FetchInfoHistorySuccess = {
  type: typeof FETCH_INFO_HISTORY_SUCCESS;
  payload: FetchInfoHistorySuccessPayload;
};

export type FetchClientTypeSuccess = {
  type: typeof FETCH_CLIENT_TYPE_SUCCESS;
  payload: FetchClientTypeSuccessPayload;
};

export type FetchLogSuccess = {
  type: typeof FETCH_LOG_SUCCESS;
  payload: FetchLogSuccessPayload;
};

export type FetchInfoFailure = {
  type: typeof FETCH_INFO_FAILURE;
  payload: FetchInfoFailurePayload;
};

export type FetchInfoHistoryFailure = {
  type: typeof FETCH_INFO_HISTORY_FAILURE;
  payload: FetchInfoHistoryFailurePayload;
};

export type FetchClientTypeFailure = {
  type: typeof FETCH_CLIENT_TYPE_FAILURE;
  payload: FetchClientTypeFailurePayload;
};

export type FetchLogFailure = {
  type: typeof FETCH_LOG_FAILURE;
  payload: FetchLogFailurePayload;
};

export type DashboardActions =
  | FetchInfoRequest
  | FetchInfoSuccess
  | FetchInfoFailure
  | FetchInfoHistoryRequest
  | FetchInfoHistorySuccess
  | FetchInfoHistoryFailure
  | FetchClientTypeRequest
  | FetchClientTypeSuccess
  | FetchClientTypeFailure
  | FetchLogRequest
  | FetchLogSuccess
  | FetchLogFailure;
