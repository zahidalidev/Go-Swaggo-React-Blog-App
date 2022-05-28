import { put, takeLatest, all } from "redux-saga/effects";

function* fetchUsers() {
  const json = yield fetch("http://localhost:8080/api/v1/allusers").then(
    (response) => response.json()
  );

  console.log("kkkkkkkkkkkkkkkkkk");
  console.log("json: ", json);

  yield put({ type: "USERS_RECIEVED", json });
}

function* actionWatcher() {
  yield takeLatest("GET_USERS", fetchUsers);
}

export default function* rootSaga() {
  yield all([actionWatcher()]);
}
