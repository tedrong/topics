import { AnyAction, combineReducers, Reducer } from "redux";

import todoReducer from "./todo/reducer";
import authReducer from "./auth/reducer";
import dashboardReducer from "./dashboard/reducer";

const combinedReducer = combineReducers({
  todo: todoReducer,
  auth: authReducer,
  dashboard: dashboardReducer,
});

const rootReducer: Reducer = (state: AppState, action: AnyAction) => {
  if (action.type === "FETCH_LOGOUT_REQUEST") {
    state = {} as AppState;
  }
  return combinedReducer(state, action);
};

export type AppState = ReturnType<typeof rootReducer>;

export default rootReducer;
