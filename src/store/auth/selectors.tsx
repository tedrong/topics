import { createSelector } from "reselect";

import { AppState } from "../rootReducer";

const getToken = (state: AppState) => state.auth.login.token;
const getUser = (state: AppState) => state.auth.login.user;
const getLoginPending = (state: AppState) => state.auth.login.pending;
const getLoginError = (state: AppState) => state.auth.login.error;

export const getTokenSelector = createSelector(getToken, (token) => token);
export const getUserSelector = createSelector(getUser, (user) => user);
export const getLoginPendingSelector = createSelector(
  getLoginPending,
  (pending) => pending
);
export const getLoginErrorSelector = createSelector(
  getLoginError,
  (error) => error
);
