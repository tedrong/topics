import axios, { AxiosResponse } from "axios";
import { all, call, put, takeLatest } from "redux-saga/effects";

import { API } from "../api";
import {
  fetchInfoSuccess,
  fetchInfoFailure,
  fetchInfoHistorySuccess,
  fetchInfoHistoryFailure,
  fetchClientTypeSuccess,
  fetchClientTypeFailure,
  fetchLogSuccess,
  fetchLogFailure,
} from "./actions";
import {
  FETCH_INFO_REQUEST,
  FETCH_INFO_HISTORY_REQUEST,
  FETCH_CLIENT_TYPE_REQUEST,
  FETCH_LOG_REQUEST,
} from "./actionTypes";
import {
  FetchInfoHistoryRequestPayload,
  FetchInfoHistoryRequest,
  FetchLogRequestPayload,
  FetchLogRequest,
  Info,
  InfoHistory,
  ClientType,
  Log,
} from "./types";

const getInfo = () => axios.get<Info>(API.dashboard.info);
const getInfoHistory = (payload: FetchInfoHistoryRequestPayload) =>
  axios.get<InfoHistory>(API.dashboard.infoHistory + "/" + payload.amount);
const getClientType = () => axios.get<ClientType>(API.dashboard.clientType);
const getLog = (payload: FetchLogRequestPayload) =>
  axios.get<Log>(API.dashboard.log + "/" + payload.amount);

/*
  Worker Saga: Fired on FETCH_INFO_REQUEST action
*/
function* fetchInfoSaga() {
  try {
    const response: AxiosResponse<Info> = yield call(getInfo);
    yield put(
      fetchInfoSuccess({
        data: response.data,
      })
    );
  } catch (e) {
    if (axios.isAxiosError(e)) {
      yield put(
        fetchInfoFailure({
          error: e.message,
        })
      );
    }
  }
}

/*
  Worker Saga: Fired on FETCH_INFO_HISTORY_REQUEST action
*/
function* fetchInfoHistorySaga(req: FetchInfoHistoryRequest) {
  try {
    const response: AxiosResponse<InfoHistory> = yield call(
      getInfoHistory,
      req.payload
    );
    yield put(
      fetchInfoHistorySuccess({
        data: response.data,
      })
    );
  } catch (e) {
    if (axios.isAxiosError(e)) {
      yield put(
        fetchInfoHistoryFailure({
          error: e.message,
        })
      );
    }
  }
}

/*
  Worker Saga: Fired on FETCH_CLIENT_TYPE_REQUEST action
*/
function* fetchClientTypeSaga() {
  try {
    const response: AxiosResponse<ClientType> = yield call(getClientType);
    yield put(
      fetchClientTypeSuccess({
        data: response.data,
      })
    );
  } catch (e) {
    if (axios.isAxiosError(e)) {
      yield put(
        fetchClientTypeFailure({
          error: e.message,
        })
      );
    }
  }
}

/*
  Worker Saga: Fired on FETCH_LOG_REQUEST action
*/
function* fetchLogSaga(req: FetchLogRequest) {
  try {
    const response: AxiosResponse<Log> = yield call(getLog, req.payload);
    yield put(
      fetchLogSuccess({
        data: response.data,
      })
    );
  } catch (e) {
    if (axios.isAxiosError(e)) {
      yield put(
        fetchLogFailure({
          error: e.message,
        })
      );
    }
  }
}

export default function* dashboardSaga() {
  yield all([
    takeLatest(FETCH_INFO_REQUEST, fetchInfoSaga),
    takeLatest(FETCH_INFO_HISTORY_REQUEST, fetchInfoHistorySaga),
    takeLatest(FETCH_CLIENT_TYPE_REQUEST, fetchClientTypeSaga),
    takeLatest(FETCH_LOG_REQUEST, fetchLogSaga),
  ]);
}
