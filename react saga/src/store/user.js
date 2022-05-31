import { createSlice } from "@reduxjs/toolkit";

const slice = createSlice({
  name: "user",
  initialState: [],
  reducers: {
    GET_USERS: (user, action) => {
      return user;
    },
    USERS_RECIEVED: (user, action) => {
      console.log("action user: ", action);
      return action.json;
    },
  },
});

export const { GET_USERS, USERS_RECIEVED } = slice.actions;
export default slice.reducer;
