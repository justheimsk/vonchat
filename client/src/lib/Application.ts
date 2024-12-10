import { CommandRegistryState } from '@/shared/state/commandRegistry';
import { ProfileState } from '@/shared/state/profiles';
import { UiState } from '@/shared/state/uiState';
import UIManager from './core/UIManager';
import CommandRegistry from './core/command/CommandRegistry';
import { InputManager } from './core/input/InputManager';
import { ProfileManager } from './core/profile/ProfileManager';
import { StateManager } from './state/StateManager';

export interface States {
	cmdRegistry: CommandRegistryState;
	ui: UiState;
	profiles: ProfileState;
}

export class Application {
	public ui: UIManager;
	public cmdRegistry: CommandRegistry;
	public input: InputManager;
	public state: StateManager<States>;
	public profiles: ProfileManager;

	public constructor() {
		this.state = new StateManager({
			cmdRegistry: new CommandRegistryState(),
			ui: new UiState(),
			profiles: new ProfileState(),
		});

		this.ui = new UIManager(this);
		this.cmdRegistry = new CommandRegistry(this);
		this.input = new InputManager(this);
		this.profiles = new ProfileManager(this);

		this.loadClientCommands();
	}

	private async loadClientCommands() {
		const commands = (await import('../shared/commands')).default;
		commands();
	}
}

export const vonchat = new Application();
