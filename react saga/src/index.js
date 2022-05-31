import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import createSagaMiddleware from "redux-saga";
import rootSaga from "./sagas";
import { Provider } from "react-redux";
import { logger } from "redux-logger";
import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";

import userReducer from "./store/user";
import postReducer from "./store/post";

let sagaMiddleware = createSagaMiddleware();
const middleware = [
  // ...getDefaultMiddleware({ thunk: false }),
  sagaMiddleware,
  // logger,
];

const store = configureStore({
  reducer: {
    user: userReducer,
    post: postReducer,
  },
  middleware,
});

sagaMiddleware.run(rootSaga);

const root = ReactDOM.createRoot(document.getElementById("root"));

root.render(
  <Provider store={store}>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </Provider>
);

reportWebVitals();
