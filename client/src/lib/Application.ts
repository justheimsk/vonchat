import { CommandRegistryState } from '@/shared/state/commandRegistry';
import { ProfileState } from '@/shared/state/profiles';
import { UiState } from '@/shared/state/uiState';
import { LocalStorageMemoryAdapter } from './adapters/LocalStorageMemoryAdapter';
import { HTTPAdapter } from './adapters/backend/HTTPAdapter';
import { LogManager } from './core/LogManager';
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
	public logs: LogManager;

	public constructor() {
		this.state = new StateManager({
			cmdRegistry: new CommandRegistryState(),
			ui: new UiState(),
			profiles: new ProfileState(),
		});

		this.logs = new LogManager(
			(log) => console.log(log.message),
			'Application',
		);
		this.ui = new UIManager(this);
		this.cmdRegistry = new CommandRegistry(
			this,
			this.logs.withTag('Command Registry'),
		);
		this.input = new InputManager(this);
		this.profiles = new ProfileManager(
			this,
			new LocalStorageMemoryAdapter(),
			this.logs.withTag('Profile Manager'),
		);

		this.loadProfiles();
		this.loadClientCommands();
	}

	private async loadClientCommands() {
		const commands = (await import('../shared/commands')).default;
		commands();
	}

	private loadProfiles() {
		const profiles = this.profiles.readInMemoryProfiles();

		if (profiles && profiles.length) {
			for (const profile of profiles) {
				const adapter = new HTTPAdapter({
					host: 'localhost',
					secure: false,
					port: 8080,
				});

				this.profiles.createProfile(
					profile.name,
					profile.email,
					profile.password,
					adapter,
				);
			}
		}
	}
}

export const vonchat = new Application();
