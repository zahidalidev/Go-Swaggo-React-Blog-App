const reducer = (state = [], action) => {
  switch (action.type) {
    case "GET_USERS":
      return [...state];

    case "USERS_RECIEVED":
      console.log("gggggggggg: ", action);
      return [...action.json];

    default:
      return state;
  }
};

export default reducer;
