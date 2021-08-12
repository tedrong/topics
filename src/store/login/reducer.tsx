import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_SUCCESS,
  FETCH_LOGIN_FAILURE,
} from "./actionTypes";

import { LoginActions, LoginState } from "./types";

const initialState: LoginState = {
  pending: false,
  tokens: {
    accessToken: "",
    refreshToken: "",
    tokenType: "",
    uid: "",
  },
  error: null,
};

export default (state = initialState, action: LoginActions) => {
  switch (action.type) {
    case FETCH_LOGIN_REQUEST:
      return {
        ...state,
        pending: true,
      };
    case FETCH_LOGIN_SUCCESS:
      return {
        ...state,
        pending: false,
        tokens: action.payload.tokens,
        error: null,
      };
    case FETCH_LOGIN_FAILURE:
      return {
        ...state,
        pending: false,
        login: [],
        error: action.payload.error,
      };
    default:
      return {
        ...state,
      };
  }
};
