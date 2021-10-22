import { createSelector } from "reselect";

import { AppState } from "../rootReducer";

const getLoginPending = (state: AppState) => state.auth.login.pending;

const getLogin = (state: AppState) => state.auth.login;

const getLoginError = (state: AppState) => state.auth.login.error;

export const getLoginSelector = createSelector(getLogin, (login) => login);

export const getLoginPendingSelector = createSelector(
  getLoginPending,
  (pending) => pending
);

export const getLoginErrorSelector = createSelector(
  getLoginError,
  (error) => error
);
