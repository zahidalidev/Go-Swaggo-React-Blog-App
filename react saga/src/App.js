import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { GET_USERS } from "./store/user";

const App = () => {
  const dispatch = useDispatch();
  const users = useSelector((state) => state);

  console.log("users: ", users);

  useEffect(() => {
    dispatch(GET_USERS());
  }, []);

  return (
    <div>
      {/* {users.map((item, index) => (
        <h1 key={index.toString()}>{item.name}</h1>
      ))} */}
    </div>
  );
};
export default App;
