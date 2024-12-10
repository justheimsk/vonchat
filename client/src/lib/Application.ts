import { CommandRegistryState } from '@/shared/state/commandRegistry';
import { UiState } from '@/shared/state/uiState';
import UIManager from './core/UIManager';
import CommandRegistry from './core/command/CommandRegistry';
import { InputManager } from './core/input/InputManager';
import { StateManager } from './state/StateManager';

export interface States {
	cmdRegistry: CommandRegistryState;
	ui: UiState;
}

export class Application {
	public ui: UIManager;
	public cmdRegistry: CommandRegistry;
	public input: InputManager;
	public state: StateManager<States>;

	public constructor() {
		this.ui = new UIManager();
		this.cmdRegistry = new CommandRegistry();
		this.input = new InputManager(this);
		this.state = new StateManager({
			cmdRegistry: new CommandRegistryState(),
			ui: new UiState(),
		});

		this.loadClientCommands();
	}

	private async loadClientCommands() {
		const commands = (await import('../shared/commands')).default;
		commands();
	}
}

export const vonchat = new Application();
