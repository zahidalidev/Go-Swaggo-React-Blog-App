import { createSlice } from "@reduxjs/toolkit";

const slice = createSlice({
  name: "post",
  initialState: [],
  reducers: {
    GET_POSTS: (post, action) => {
      return post;
    },
    POSTS_RECIEVED: (post, action) => {
      return action.json;
    },
  },
});

export const { GET_POSTS, POSTS_RECIEVED } = slice.actions;
export default slice.reducer;
