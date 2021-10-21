import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_SUCCESS,
  FETCH_LOGIN_FAILURE,
} from "./actionTypes";

export interface Token {
  access_token: string;
  refresh_token: string;
}

export interface User {
  ID: number;
  UUID: string;
  name: string;
}

export interface LoginPayload {
  pending: boolean;
  message: string;
  token: Token;
  user: User;
  error: string | null;
}

export interface UserState {
  login: LoginPayload;
}

export interface FetchLoginRequestPayload {
  email: string;
  password: string;
}
export interface FetchLoginSuccessPayload {
  data: LoginPayload;
}

export interface FetchLoginFailurePayload {
  error: string;
}

export interface FetchLoginRequest {
  type: typeof FETCH_LOGIN_REQUEST;
  payload: FetchLoginRequestPayload;
}

export type FetchLoginSuccess = {
  type: typeof FETCH_LOGIN_SUCCESS;
  payload: FetchLoginSuccessPayload;
};

export type FetchLoginFailure = {
  type: typeof FETCH_LOGIN_FAILURE;
  payload: FetchLoginFailurePayload;
};

export type UserActions =
  | FetchLoginRequest
  | FetchLoginSuccess
  | FetchLoginFailure;
