import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_SUCCESS,
  FETCH_LOGIN_FAILURE,
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
};

const reducer = (state = initialState, action: AuthActions) => {
  switch (action.type) {
    case FETCH_LOGIN_REQUEST:
      return {
        ...state,
        login: {
          pending: true,
        },
      };
    case FETCH_LOGIN_SUCCESS:
      return {
        ...state,
        login: {
          ...action.payload.data,
          pending: false,
          error: null,
        },
      };
    case FETCH_LOGIN_FAILURE:
      return {
        ...state,
        login: {
          pending: false,
          error: action.payload.error,
        },
      };
    default:
      return {
        ...state,
      };
  }
};
export default reducer;
