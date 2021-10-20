import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_SUCCESS,
  FETCH_LOGIN_FAILURE,
} from "./actionTypes";

export interface ITokens {
  access_token: string;
  refresh_token: string;
}

export interface LoginState {
  pending: boolean;
  tokens: ITokens;
  error: string | null;
}

export interface FetchLoginRequestPayload {
  email: string;
  password: string;
}
export interface FetchLoginSuccessPayload {
  data: ITokens;
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

export type LoginActions =
  | FetchLoginRequest
  | FetchLoginSuccess
  | FetchLoginFailure;
