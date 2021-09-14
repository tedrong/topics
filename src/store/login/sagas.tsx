import axios, { AxiosResponse } from "axios";
import { all, call, put, takeLatest } from "redux-saga/effects";

import { fetchLoginFailure, fetchLoginSuccess } from "./actions";
import { FETCH_LOGIN_REQUEST } from "./actionTypes";
import { ITokens } from "./types";

const getLogin = () => axios.get<ITokens>("/test");
/*
  Worker Saga: Fired on FETCH_LOGIN_REQUEST action
*/
function* fetchLoginSaga() {
  try {
    const response: AxiosResponse<ITokens> = yield call(getLogin);
    yield put(
      fetchLoginSuccess({
        tokens: response.data,
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
