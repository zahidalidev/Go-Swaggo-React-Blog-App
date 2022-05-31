import { put, takeLatest, all } from "redux-saga/effects";
import { url } from "../config/Api";
import { GET_POSTS, POSTS_RECIEVED } from "../store/post";
import { GET_USERS, USERS_RECIEVED } from "../store/user";

function* fetchUsers() {
  const json = yield fetch(`${url}/allusers`).then((response) =>
    response.json()
  );

  yield put({ type: USERS_RECIEVED, json });
}

function* actionWatcherUser() {
  yield takeLatest(GET_USERS, fetchUsers);
}

function* fetchPosts() {
  const json = yield fetch(`${url}/allposts`).then((response) =>
    response.json()
  );
  yield put({ type: POSTS_RECIEVED, json });
}

function* actionWatcherPost() {
  yield takeLatest(GET_POSTS, fetchPosts);
}

export default function* rootSaga() {
  yield all([actionWatcherUser(), actionWatcherPost()]);
}
