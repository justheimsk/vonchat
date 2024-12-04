import type Command from "@/lib/core/Command";
import {createSlice, type PayloadAction} from "@reduxjs/toolkit";

export interface CommandRegistryState {
  commands: Command[];
}

const initialState: CommandRegistryState = {
  commands: []
}

const commandRegistry = createSlice({
  name: 'commandRegistry',
  initialState,
  reducers: {
    registerCommand: (state: CommandRegistryState, action: PayloadAction<Command>) => {
      state.commands.push(action.payload);
    }
  }
})

export const { registerCommand } = commandRegistry.actions

export default commandRegistry.reducer
