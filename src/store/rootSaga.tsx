import { all, fork } from "redux-saga/effects";

import authSaga from "./auth/sagas";
import dashboardSaga from "./dashboard/sagas";

export function* rootSaga() {
  yield all([fork(authSaga), fork(dashboardSaga)]);
}
