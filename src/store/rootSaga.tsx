import { all, fork } from "redux-saga/effects";

import todoSaga from "./todo/sagas";
import loginSaga from "./login/sagas";

export function* rootSaga() {
  yield all([fork(todoSaga), fork(loginSaga)]);
}
