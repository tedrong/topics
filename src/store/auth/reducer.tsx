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

import { AuthActions, AuthState } from "./types";

const initialState: AuthState = {
  login: {
    pending: false,
    message: "",
    token: {
      access_token: "",
      refresh_token: "",
    },
    user: {
      ID: -1,
      UUID: "",
      first_name: "",
      last_name: "",
    },
    error: null,
  },
  renew: {
    pending: false,
    user: {
      ID: -1,
      UUID: "",
      first_name: "",
      last_name: "",
    },
    error: null,
  },
  logout: {
    pending: false,
    error: null,
  },
};

const reducer = (state = initialState, action: AuthActions) => {
  switch (action.type) {
    case FETCH_LOGIN_REQUEST:
      return {
        ...state,
        login: { pending: true },
      };
    case FETCH_LOGIN_SUCCESS:
      return {
        ...state,
        login: { ...action.payload.data, pending: false, error: null },
      };
    case FETCH_LOGIN_FAILURE:
      return {
        ...state,
        login: { pending: false, error: action.payload.error },
      };
    case FETCH_RENEW_REQUEST:
      return {
        ...state,
        renew: { pending: true },
      };
    case FETCH_RENEW_SUCCESS:
      return {
        ...state,
        renew: { user: action.payload.data.user },
        login: {
          ...state.login,
          user: action.payload.data.user,
        },
      };
    case FETCH_RENEW_FAILURE:
      return {
        ...state,
        renew: { pending: false, error: action.payload.error },
      };
    case FETCH_LOGOUT_REQUEST:
      return {
        ...state,
        logout: {
          pending: true,
        },
      };
    case FETCH_LOGOUT_SUCCESS:
      return { ...state, logout: { pending: false, error: null } };
    case FETCH_LOGOUT_FAILURE:
      return {
        ...state,
        logout: { pending: false, error: action.payload.error },
      };
    default:
      return { ...state };
  }
};
export default reducer;
