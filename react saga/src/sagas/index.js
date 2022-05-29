import { put, takeLatest, all } from "redux-saga/effects";
import { url } from "../config/Api";
import { GET_USERS, USERS_RECIEVED } from "../store/user";

function* fetchUsers() {
  const json = yield fetch(`${url}/allusers`).then((response) =>
    response.json()
  );

  yield put({ type: USERS_RECIEVED, json });
}

function* actionWatcher() {
  yield takeLatest(GET_USERS, fetchUsers);
}

export default function* rootSaga() {
  yield all([actionWatcher()]);
}
