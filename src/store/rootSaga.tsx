import { all, fork } from "redux-saga/effects";

import todoSaga from "./todo/sagas";
import authSaga from "./auth/sagas";
import dashboardSaga from "./dashboard/sagas";

export function* rootSaga() {
  yield all([fork(todoSaga), fork(authSaga), fork(dashboardSaga)]);
}
