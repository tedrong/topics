import {
  FETCH_INFO_REQUEST,
  FETCH_INFO_FAILURE,
  FETCH_INFO_SUCCESS,
  FETCH_INFO_HISTORY_REQUEST,
  FETCH_INFO_HISTORY_FAILURE,
  FETCH_INFO_HISTORY_SUCCESS,
  FETCH_CLIENT_TYPE_REQUEST,
  FETCH_CLIENT_TYPE_FAILURE,
  FETCH_CLIENT_TYPE_SUCCESS,
  FETCH_LOG_REQUEST,
  FETCH_LOG_FAILURE,
  FETCH_LOG_SUCCESS,
} from "./actionTypes";
import {
  FetchInfoRequest,
  FetchInfoSuccessPayload,
  FetchInfoSuccess,
  FetchInfoFailurePayload,
  FetchInfoFailure,
  FetchInfoHistoryRequestPayload,
  FetchInfoHistoryRequest,
  FetchInfoHistorySuccessPayload,
  FetchInfoHistorySuccess,
  FetchInfoHistoryFailurePayload,
  FetchInfoHistoryFailure,
  FetchClientTypeRequest,
  FetchClientTypeSuccessPayload,
  FetchClientTypeSuccess,
  FetchClientTypeFailurePayload,
  FetchClientTypeFailure,
  FetchLogRequestPayload,
  FetchLogRequest,
  FetchLogSuccessPayload,
  FetchLogSuccess,
  FetchLogFailurePayload,
  FetchLogFailure,
} from "./types";

export const fetchInfoRequest = (): FetchInfoRequest => ({
  type: FETCH_INFO_REQUEST,
});

export const fetchInfoHistoryRequest = (
  payload: FetchInfoHistoryRequestPayload
): FetchInfoHistoryRequest => ({
  type: FETCH_INFO_HISTORY_REQUEST,
  payload,
});

export const fetchClientTypeRequest = (): FetchClientTypeRequest => ({
  type: FETCH_CLIENT_TYPE_REQUEST,
});

export const fetchLogRequest = (
  payload: FetchLogRequestPayload
): FetchLogRequest => ({
  type: FETCH_LOG_REQUEST,
  payload,
});

export const fetchInfoSuccess = (
  payload: FetchInfoSuccessPayload
): FetchInfoSuccess => ({
  type: FETCH_INFO_SUCCESS,
  payload,
});

export const fetchInfoHistorySuccess = (
  payload: FetchInfoHistorySuccessPayload
): FetchInfoHistorySuccess => ({
  type: FETCH_INFO_HISTORY_SUCCESS,
  payload,
});

export const fetchClientTypeSuccess = (
  payload: FetchClientTypeSuccessPayload
): FetchClientTypeSuccess => ({
  type: FETCH_CLIENT_TYPE_SUCCESS,
  payload,
});

export const fetchLogSuccess = (
  payload: FetchLogSuccessPayload
): FetchLogSuccess => ({
  type: FETCH_LOG_SUCCESS,
  payload,
});

export const fetchInfoFailure = (
  payload: FetchInfoFailurePayload
): FetchInfoFailure => ({
  type: FETCH_INFO_FAILURE,
  payload,
});

export const fetchInfoHistoryFailure = (
  payload: FetchInfoHistoryFailurePayload
): FetchInfoHistoryFailure => ({
  type: FETCH_INFO_HISTORY_FAILURE,
  payload,
});

export const fetchClientTypeFailure = (
  payload: FetchClientTypeFailurePayload
): FetchClientTypeFailure => ({
  type: FETCH_CLIENT_TYPE_FAILURE,
  payload,
});

export const fetchLogFailure = (
  payload: FetchLogFailurePayload
): FetchLogFailure => ({
  type: FETCH_LOG_FAILURE,
  payload,
});
