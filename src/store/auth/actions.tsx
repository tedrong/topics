import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_FAILURE,
  FETCH_LOGIN_SUCCESS,
  FETCH_RENEW_REQUEST,
  FETCH_RENEW_FAILURE,
  FETCH_RENEW_SUCCESS,
} from "./actionTypes";
import {
  User,
  FetchLoginRequest,
  FetchLoginSuccess,
  FetchLoginFailure,
  FetchLoginRequestPayload,
  FetchLoginSuccessPayload,
  FetchLoginFailurePayload,
  FetchRenewRequest,
  FetchRenewSuccess,
  FetchRenewFailure,
  FetchRenewSuccessPayload,
  FetchRenewFailurePayload,
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
