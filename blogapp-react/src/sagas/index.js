import { put, takeLatest, all } from "redux-saga/effects";

function* fetchNews() {
  const json = yield fetch(
    "https://newsapi.org/v2/everything?q=tesla&from=2022-03-25&sortBy=publishedAt&apiKey=b69eec62000447eb8d3d5b50d076920c"
  ).then((response) => response.json());
  yield put({ type: "NEWS_RECEIVED", json: json.articles });
}

function* actionWatcher() {
  yield takeLatest("GET_NEWS", fetchNews);
}

export default function* rootSaga() {
  yield all([actionWatcher()]);
}
