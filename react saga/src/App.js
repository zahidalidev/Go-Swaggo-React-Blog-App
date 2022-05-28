import axios from "axios";
import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { getUsers } from "./actions";

const App = () => {
  // const dispatch = useDispatch();
  // const users = useSelector((state) => state);
  // console.log("users: ", users);

  const getAllUsers = async () => {
    const data = await axios.get("http://localhost:8080/api/v1/allusers");
    console.log("data: ", data);
  };

  useEffect(() => {
    getAllUsers();
    // dispatch(getUsers());
  }, []);

  return (
    <div>
      <h1></h1>
    </div>
  );
};
export default App;
