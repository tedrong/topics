import axios, { AxiosResponse } from "axios";
import { all, call, put, takeLatest } from "redux-saga/effects";

import { API, attachAuthToken } from "../api";
import {
  fetchLoginSuccess,
  fetchLoginFailure,
  fetchRenewSuccess,
  fetchRenewFailure,
} from "./actions";
import { FETCH_LOGIN_REQUEST, FETCH_RENEW_REQUEST } from "./actionTypes";
import {
  User,
  LoginPayload,
  FetchLoginRequestPayload,
  FetchLoginRequest,
  RenewPayload,
  FetchRenewRequest,
} from "./types";

const postLogin = (payload: FetchLoginRequestPayload) =>
  axios.post<LoginPayload>(API.user.login, payload);
const putRenew = (payload: User) =>
  axios.put<RenewPayload>(API.user.renew + "/" + payload.UUID, payload);

/*
  Worker Saga: Fired on FETCH_LOGIN_REQUEST action
*/
function* fetchLoginSaga(req: FetchLoginRequest) {
  try {
    const response: AxiosResponse<LoginPayload> = yield call(
      postLogin,
      req.payload
    );
    yield put(
      fetchLoginSuccess({
        data: response.data,
      })
    );
    attachAuthToken(response.data.token.access_token);
  } catch (e) {
    if (axios.isAxiosError(e)) {
      yield put(
        fetchLoginFailure({
          error: e.message,
        })
      );
    }
  }
}

/*
  Worker Saga: Fired on FETCH_RENEW_REQUEST action
*/
function* fetchRenewSaga(req: FetchRenewRequest) {
  try {
    const response: AxiosResponse<RenewPayload> = yield call(
      putRenew,
      req.payload
    );
    yield put(
      fetchRenewSuccess({
        data: response.data,
      })
    );
  } catch (e) {
    if (axios.isAxiosError(e)) {
      yield put(
        fetchRenewFailure({
          error: e.message,
        })
      );
    }
  }
}

/*
  Starts worker saga on latest dispatched `FETCH_LOGIN_REQUEST` action.
  Allows concurrent increments.
*/
export default function* loginSaga() {
  yield all([
    takeLatest(FETCH_LOGIN_REQUEST, fetchLoginSaga),
    takeLatest(FETCH_RENEW_REQUEST, fetchRenewSaga),
  ]);
}
