import {
  FETCH_LOGIN_REQUEST,
  FETCH_LOGIN_SUCCESS,
  FETCH_LOGIN_FAILURE,
} from "./actionTypes";

import { UserActions, UserState } from "./types";

const initialState: UserState = {
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
      name: "",
    },
    error: null,
  },
};

export default (state = initialState, action: UserActions) => {
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
