import React from "react";
import { connect, useDispatch } from "react-redux";
import { getNews } from "../actions";

let Button = () => {
  const dispatch = useDispatch();

  return <button onClick={() => dispatch(getNews())}>Press to see news</button>;
};

export default Button;
