import axios, { AxiosResponse } from "axios";
import { all, call, put, takeLatest } from "redux-saga/effects";

import { API } from "../api";
import { fetchLoginFailure, fetchLoginSuccess } from "./actions";
import { FETCH_LOGIN_REQUEST } from "./actionTypes";
import { FetchLoginRequest, FetchLoginRequestPayload, ITokens } from "./types";

const postLogin = (payload: FetchLoginRequestPayload) =>
  axios.post<ITokens>(API.user.login, payload);
/*
  Worker Saga: Fired on FETCH_LOGIN_REQUEST action
*/
function* fetchLoginSaga(req: FetchLoginRequest) {
  try {
    const response: AxiosResponse<ITokens> = yield call(postLogin, req.payload);
    yield put(
      fetchLoginSuccess({
        data: response.data,
      })
    );
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
