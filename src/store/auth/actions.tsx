import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_SUCCESS,
  FETCH_LOGIN_FAILURE,
  FETCH_RENEW_REQUEST,
  FETCH_RENEW_SUCCESS,
  FETCH_RENEW_FAILURE,
  FETCH_LOGOUT_REQUEST,
  FETCH_LOGOUT_SUCCESS,
  FETCH_LOGOUT_FAILURE,
} from "./actionTypes";
import {
  User,
  FetchLoginRequestPayload,
  FetchLoginSuccessPayload,
  FetchLoginFailurePayload,
  FetchLoginRequest,
  FetchLoginSuccess,
  FetchLoginFailure,
  FetchRenewSuccessPayload,
  FetchRenewFailurePayload,
  FetchRenewRequest,
  FetchRenewSuccess,
  FetchRenewFailure,
  FetchLogoutFailurePayload,
  FetchLogoutRequest,
  FetchLogoutSuccess,
  FetchLogoutFailure,
} from "./types";

export const fetchLoginRequest = (
  payload: FetchLoginRequestPayload
): FetchLoginRequest => ({
  type: FETCH_LOGIN_REQUEST,
  payload,
});

export const fetchRenewRequest = (payload: User): FetchRenewRequest => ({
  type: FETCH_RENEW_REQUEST,
  payload,
});

export const fetchLogoutRequest = (): FetchLogoutRequest => ({
  type: FETCH_LOGOUT_REQUEST,
});

export const fetchLoginSuccess = (
  payload: FetchLoginSuccessPayload
): FetchLoginSuccess => ({
  type: FETCH_LOGIN_SUCCESS,
  payload,
});

export const fetchRenewSuccess = (
  payload: FetchRenewSuccessPayload
): FetchRenewSuccess => ({
  type: FETCH_RENEW_SUCCESS,
  payload,
});

export const fetchLogoutSuccess = (): FetchLogoutSuccess => ({
  type: FETCH_LOGOUT_SUCCESS,
});

export const fetchLoginFailure = (
  payload: FetchLoginFailurePayload
): FetchLoginFailure => ({
  type: FETCH_LOGIN_FAILURE,
  payload,
});

export const fetchRenewFailure = (
  payload: FetchRenewFailurePayload
): FetchRenewFailure => ({
  type: FETCH_RENEW_FAILURE,
  payload,
});

export const fetchLogoutFailure = (
  payload: FetchLogoutFailurePayload
): FetchLogoutFailure => ({
  type: FETCH_LOGOUT_FAILURE,
  payload,
});
