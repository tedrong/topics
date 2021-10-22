import axios, { AxiosResponse } from "axios";
import { all, call, put, takeLatest } from "redux-saga/effects";

import { API, attachAuthToken } from "../api";
import { fetchLoginFailure, fetchLoginSuccess } from "./actions";
import { FETCH_LOGIN_REQUEST } from "./actionTypes";
import {
  FetchLoginRequest,
  FetchLoginRequestPayload,
  LoginPayload,
} from "./types";

const postLogin = (payload: FetchLoginRequestPayload) =>
  axios.post<LoginPayload>(API.user.login, payload);
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
  Starts worker saga on latest dispatched `FETCH_LOGIN_REQUEST` action.
  Allows concurrent increments.
*/
export default function* loginSaga() {
  yield all([takeLatest(FETCH_LOGIN_REQUEST, fetchLoginSaga)]);
}
