import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import createSagaMiddleware from "redux-saga";
import reducer from "./store/user";
import rootSaga from "./sagas";
import { Provider } from "react-redux";
import { logger } from "redux-logger";
import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";

let sagaMiddleware = createSagaMiddleware();
const middleware = [
  // ...getDefaultMiddleware({ thunk: false }),
  sagaMiddleware,
  logger,
];

// const store = configureStore(
//   { reducer,  applyMiddleware(sagaMiddleware, logger)},
// );

const store = configureStore({
  reducer: reducer,

  middleware,
});

// `middleware: (gDM) => gDM().concat(logger, apiMiddleware, yourCustomMiddleware)`;

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
