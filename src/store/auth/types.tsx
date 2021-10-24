import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_SUCCESS,
  FETCH_LOGIN_FAILURE,
  FETCH_RENEW_REQUEST,
  FETCH_RENEW_SUCCESS,
  FETCH_RENEW_FAILURE,
} from "./actionTypes";

export interface Token {
  access_token: string;
  refresh_token: string;
}

export interface User {
  ID: number;
  UUID: string;
  first_name: string;
  last_name: string;
}

export interface LoginPayload {
  pending: boolean;
  message: string;
  token: Token;
  user: User;
  error: string | null;
}

export interface RenewPayload {
  pending: boolean;
  user: User;
  error: string | null;
}

export interface AuthState {
  login: LoginPayload;
  renew: RenewPayload;
}

export interface FetchLoginRequestPayload {
  email: string;
  password: string;
}

export interface FetchLoginSuccessPayload {
  data: LoginPayload;
}

export interface FetchRenewSuccessPayload {
  data: RenewPayload;
}

export interface FetchLoginFailurePayload {
  error: string;
}

export interface FetchRenewFailurePayload {
  error: string;
}

export interface FetchLoginRequest {
  type: typeof FETCH_LOGIN_REQUEST;
  payload: FetchLoginRequestPayload;
}

export interface FetchRenewRequest {
  type: typeof FETCH_RENEW_REQUEST;
  payload: User;
}

export type FetchLoginSuccess = {
  type: typeof FETCH_LOGIN_SUCCESS;
  payload: FetchLoginSuccessPayload;
};

export type FetchRenewSuccess = {
  type: typeof FETCH_RENEW_SUCCESS;
  payload: FetchRenewSuccessPayload;
};

export type FetchLoginFailure = {
  type: typeof FETCH_LOGIN_FAILURE;
  payload: FetchLoginFailurePayload;
};

export type FetchRenewFailure = {
  type: typeof FETCH_RENEW_FAILURE;
  payload: FetchRenewFailurePayload;
};

export type AuthActions =
  | FetchLoginRequest
  | FetchLoginSuccess
  | FetchLoginFailure
  | FetchRenewRequest
  | FetchRenewSuccess
  | FetchRenewFailure;
