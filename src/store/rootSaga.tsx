import { all, fork } from "redux-saga/effects";

import todoSaga from "./todo/sagas";
import userSaga from "./user/sagas";
import dashboardSaga from "./dashboard/sagas";

export function* rootSaga() {
  yield all([fork(todoSaga), fork(userSaga), fork(dashboardSaga)]);
}
