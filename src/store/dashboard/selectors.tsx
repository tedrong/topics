import { createSelector } from "reselect";
import { AppState } from "../rootReducer";

const getInfo = (state: AppState) => state.dashboard.info.data;
const getInfoPending = (state: AppState) => state.dashboard.info.pending;
const getInfoError = (state: AppState) => state.dashboard.info.error;
const getInfoHistory = (state: AppState) => state.dashboard.infoHistory.data;
const getInfoHistoryPending = (state: AppState) =>
  state.dashboard.infoHistory.pending;
const getInfoHistoryError = (state: AppState) =>
  state.dashboard.infoHistory.error;
const getClientType = (state: AppState) => state.dashboard.clientType.data;
const getClientTypePending = (state: AppState) =>
  state.dashboard.clientType.pending;
const getClientTypeError = (state: AppState) =>
  state.dashboard.clientType.error;
const getLog = (state: AppState) => state.dashboard.log.data;
const getLogPending = (state: AppState) => state.dashboard.log.pending;
const getLogError = (state: AppState) => state.dashboard.log.error;

export const getInfoSelector = createSelector(getInfo, (info) => info);
export const getInfoPendingSelector = createSelector(
  getInfoPending,
  (pending) => pending
);
export const getInfoErrorSelector = createSelector(
  getInfoError,
  (error) => error
);
export const getInfoHistorySelector = createSelector(
  getInfoHistory,
  (infoHistory) => infoHistory
);
export const getInfoHistoryPendingSelector = createSelector(
  getInfoHistoryPending,
  (pending) => pending
);
export const getInfoHistoryErrorSelector = createSelector(
  getInfoHistoryError,
  (error) => error
);
export const getClientTypeSelector = createSelector(
  getClientType,
  (clientType) => clientType
);
export const getClientTypePendingSelector = createSelector(
  getClientTypePending,
  (pending) => pending
);
export const getClientTypeErrorSelector = createSelector(
  getClientTypeError,
  (error) => error
);
export const getLogSelector = createSelector(getLog, (log) => log);
export const getLogPendingSelector = createSelector(
  getLogPending,
  (pending) => pending
);
export const getLogErrorSelector = createSelector(
  getLogError,
  (error) => error
);
