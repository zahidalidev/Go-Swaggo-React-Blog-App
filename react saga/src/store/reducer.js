import { combineReducers } from "@reduxjs/toolkit";

import userReducer from "../store/user";

export default combineReducers({
  user: userReducer,
});
