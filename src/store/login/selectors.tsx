import { createSelector } from "reselect";

import { AppState } from "../rootReducer";

const getPending = (state: AppState) => state.login.pending;

const getLogin = (state: AppState) => state.login.tokens;

const getError = (state: AppState) => state.login.error;

export const getLoginSelector = createSelector(getLogin, (login) => login);

export const getPendingSelector = createSelector(
  getPending,
  (pending) => pending
);

export const getErrorSelector = createSelector(getError, (error) => error);
