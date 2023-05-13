import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import client from "../../../../apis";

const FaculitySlice = createSlice({
  name: "faculity",
  initialState: {
    status: "idle",
    faculties: [],
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchFaculity.pending, (state, action) => {
        state.status = "pending";
      })
      .addCase(fetchFaculity.fulfilled, (state, action) => {
        state.status = "idle";
        if (action.payload.faculties !== null) {
          state.faculties = action.payload.faculties;
        }
      });
  },
});

export const fetchFaculity = createAsyncThunk("getFaculities", async () => {
  const response = await client.get("/api/faculities");

  return response.data;
});

export const FaculityReducer = FaculitySlice.reducer;
export default FaculitySlice;
