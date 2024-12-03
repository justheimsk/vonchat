import {createSlice, type PayloadAction} from "@reduxjs/toolkit";

export interface UiState {
  commandList: boolean;
}

const initialState: UiState = {
  commandList: false
}

export const uiSlice = createSlice({
  name: 'ui',
  initialState,
  reducers: {
    toggleCommandList: (state: UiState, action: PayloadAction<boolean>) => {
      state.commandList = action.payload;
    }
  }
})

export const { toggleCommandList } = uiSlice.actions

export default uiSlice.reducer
