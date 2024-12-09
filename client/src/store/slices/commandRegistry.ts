import type Command from '@/lib/core/command/Command';
import { type PayloadAction, createSlice } from '@reduxjs/toolkit';

export interface CommandRegistryState {
	commands: Command[];
}

const initialState: CommandRegistryState = {
	commands: [],
};

const commandRegistry = createSlice({
	name: 'commandRegistry',
	initialState,
	reducers: {
		registerCommand: (
			state: CommandRegistryState,
			action: PayloadAction<Command>,
		) => {
			state.commands.push(action.payload);
		},
	},
});

export const { registerCommand } = commandRegistry.actions;

export default commandRegistry.reducer;
