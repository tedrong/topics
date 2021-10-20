import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_FAILURE,
  FETCH_LOGIN_SUCCESS,
} from "./actionTypes";
import {
  FetchLoginRequest,
  FetchLoginSuccess,
  FetchLoginSuccessPayload,
  FetchLoginFailure,
  FetchLoginFailurePayload,
  FetchLoginRequestPayload,
} from "./types";

export const fetchLoginRequest = (
  payload: FetchLoginRequestPayload
): FetchLoginRequest => ({
  type: FETCH_LOGIN_REQUEST,
  payload,
});

export const fetchLoginSuccess = (
  payload: FetchLoginSuccessPayload
): FetchLoginSuccess => ({
  type: FETCH_LOGIN_SUCCESS,
  payload,
});

export const fetchLoginFailure = (
  payload: FetchLoginFailurePayload
): FetchLoginFailure => ({
  type: FETCH_LOGIN_FAILURE,
  payload,
});
