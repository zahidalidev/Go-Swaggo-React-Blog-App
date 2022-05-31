import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { allposts, GET_POSTS } from "./store/post";
import { allusers, GET_USERS } from "./store/user";

const App = () => {
  const dispatch = useDispatch();
  const users = useSelector((state) => state.user);
  const posts = useSelector((state) => state.post);

  console.log("users: ", users);
  console.log("posts: ", posts);

  useEffect(() => {
    dispatch(GET_POSTS());
    dispatch(GET_USERS());
  }, []);

  return (
    <div>
      {users.map((item, index) => (
        <h1 key={index.toString()}>{item.name}</h1>
      ))}
    </div>
  );
};
export default App;
