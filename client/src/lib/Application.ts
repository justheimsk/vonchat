import CommandRegistry from './core/CommandRegistry';
import UIManager from './core/UIManager';
import { InputManager } from './core/input/InputManager';

export class Application {
	public ui: UIManager;
	public cmdRegistry: CommandRegistry;
	public input: InputManager;

	public constructor() {
		this.ui = new UIManager();
		this.cmdRegistry = new CommandRegistry();
		this.input = new InputManager(this);

		this.loadClientCommands();
	}

	private async loadClientCommands() {
		const commands = (await import('../shared/commands')).default;
		commands();
	}
}

export const vonchat = new Application();
