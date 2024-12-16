import { CommandRegistryState } from '@/shared/state/commandRegistry';
import { ProfileState } from '@/shared/state/profiles';
import { UiState } from '@/shared/state/uiState';
import { to } from '@/utils/to';
import { LocalStorageMemoryAdapter } from './adapters/LocalStorageMemoryAdapter';
import { HTTPAdapter } from './adapters/backend/HTTPAdapter';
import { LogManager } from './core/LogManager';
import { Server } from './core/Server';
import CommandRegistry from './core/command/CommandRegistry';
import { InputManager } from './core/input/InputManager';
import { ProfileManager } from './core/profile/ProfileManager';
import UIManager from './core/ui/UIManager';
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
			(logs) => console.log(logs.map((log) => log.message).join('\n')),
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
		const [commands, err] = await to(import('../shared/commands'));
		if (err) return this.logs.send('error', 'Failed to load client commands.');

		if (commands) commands.default();
	}

	private loadProfiles() {
		try {
			const profiles = this.profiles.readInMemoryProfiles();

			if (profiles && profiles.length) {
				for (const profile of profiles) {
					const servers: Server[] = [];
					for (const server of profile.servers) {
						servers.push(
							new Server(
								server.ip,
								server.port,
								new HTTPAdapter({
									host: server.ip,
									port: server.port,
									secure: false,
								}),
							),
						);
					}

					this.profiles.createProfile(
						profile.name,
						profile.email,
						profile.password,
						profile.active,
						profile.id,
						servers,
					);
				}
			}
		} catch {
			return this.logs.send('error', 'Failed to read in-memory profiles');
		}
	}
}

export const vonchat = new Application();
