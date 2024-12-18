import type { CommandRegistryState } from '@/shared/state/commandRegistry';
import type { ProfileState } from '@/shared/state/profiles';
import type { UiState } from '@/shared/state/uiState';
import type { State } from '../state/State';

// biome-ignore lint/suspicious/noExplicitAny: <explanation>
export type Reducers = { [key: string]: State<any> };

export interface ApplicationState {
	cmdRegistry: CommandRegistryState;
	ui: UiState;
	profiles: ProfileState;
}
