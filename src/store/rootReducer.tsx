import { combineReducers } from "redux";

import todoReducer from "./todo/reducer";
import loginReducer from "./login/reducer";

const rootReducer = combineReducers({
  todo: todoReducer,
  login: loginReducer,
});

export type AppState = ReturnType<typeof rootReducer>;

export default rootReducer;
