import {createSlice, type PayloadAction} from "@reduxjs/toolkit";

export interface UiState {
  commandList: boolean;
  selectedCommand: string;
}

const initialState: UiState = {
  commandList: false,
  selectedCommand: ""
}

export const uiSlice = createSlice({
  name: 'ui',
  initialState,
  reducers: {
    toggleCommandList: (state: UiState, action: PayloadAction<boolean>) => {
      state.commandList = action.payload;
    },
    selectCommand: (state: UiState, action: PayloadAction<string>) => {
      state.selectedCommand = action.payload;
    }
  }
})

export const { toggleCommandList, selectCommand } = uiSlice.actions

export default uiSlice.reducer
