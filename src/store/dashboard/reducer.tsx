import {
  FETCH_INFO_REQUEST,
  FETCH_INFO_SUCCESS,
  FETCH_INFO_FAILURE,
  FETCH_INFO_HISTORY_REQUEST,
  FETCH_INFO_HISTORY_SUCCESS,
  FETCH_INFO_HISTORY_FAILURE,
  FETCH_CLIENT_TYPE_REQUEST,
  FETCH_CLIENT_TYPE_SUCCESS,
  FETCH_CLIENT_TYPE_FAILURE,
  FETCH_LOG_REQUEST,
  FETCH_LOG_SUCCESS,
  FETCH_LOG_FAILURE,
} from "./actionTypes";

import { DashboardActions, DashboardState } from "./types";

const initialState: DashboardState = {
  info: {
    pending: false,
    data: { cpu: "", memory: "", disk: "", bootTime: 0 },
    error: null,
  },
  infoHistory: {
    pending: false,
    data: { cpu: [], memory: [], disk: [], label: [] },
    error: null,
  },
  clientType: {
    pending: false,
    data: { desktop: 0, mobile: 0 },
    error: null,
  },
  log: {
    pending: false,
    data: [],
    error: null,
  },
};

const reducer = (state = initialState, action: DashboardActions) => {
  switch (action.type) {
    case FETCH_INFO_REQUEST:
      return {
        ...state,
        info: {
          pending: true,
        },
      };
    case FETCH_INFO_SUCCESS:
      return {
        ...state,
        info: {
          pending: false,
          data: action.payload.data,
          error: null,
        },
      };
    case FETCH_INFO_FAILURE:
      return {
        ...state,
        info: {
          pending: false,
          data: {},
          error: action.payload.error,
        },
      };
    case FETCH_INFO_HISTORY_REQUEST:
      return {
        ...state,
        infoHistory: {
          pending: true,
        },
      };
    case FETCH_INFO_HISTORY_SUCCESS:
      return {
        ...state,
        infoHistory: {
          pending: false,
          data: action.payload.data,
          error: null,
        },
      };
    case FETCH_INFO_HISTORY_FAILURE:
      return {
        ...state,
        infoHistory: {
          pending: false,
          data: {},
          error: action.payload.error,
        },
      };
    case FETCH_CLIENT_TYPE_REQUEST:
      return {
        ...state,
        clientType: {
          pending: true,
        },
      };
    case FETCH_CLIENT_TYPE_SUCCESS:
      return {
        ...state,
        clientType: {
          pending: false,
          data: action.payload.data,
          error: null,
        },
      };
    case FETCH_CLIENT_TYPE_FAILURE:
      return {
        ...state,
        clientType: {
          pending: false,
          data: {},
          error: action.payload.error,
        },
      };
    case FETCH_LOG_REQUEST:
      return {
        ...state,
        log: {
          pending: true,
        },
      };
    case FETCH_LOG_SUCCESS:
      return {
        ...state,
        log: {
          pending: false,
          data: action.payload.data,
          error: null,
        },
      };
    case FETCH_LOG_FAILURE:
      return {
        ...state,
        log: {
          pending: false,
          data: {},
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
