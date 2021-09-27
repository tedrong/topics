import { AnyAction, combineReducers, Reducer } from "redux";

import todoReducer from "./todo/reducer";
import loginReducer from "./login/reducer";

const combinedReducer = combineReducers({
  todo: todoReducer,
  login: loginReducer,
});

const rootReducer: Reducer = (state: AppState, action: AnyAction) => {
  if (action.type === "FETCH_LOGOUT_REQUEST") {
    state = {} as AppState;
  }
  return combinedReducer(state, action);
};

export type AppState = ReturnType<typeof rootReducer>;

export default rootReducer;
