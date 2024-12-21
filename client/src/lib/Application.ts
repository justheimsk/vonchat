import { CommandRegistryState } from '@/shared/state/commandRegistry';
import { ProfileState } from '@/shared/state/profiles';
import { UiState } from '@/shared/state/uiState';
import { to } from '@/utils/to';
import { LocalStorageMemoryAdapter } from './adapters/LocalStorageMemoryAdapter';
import { HTTPAdapter } from './adapters/backend/http/HTTPAdapter';
import { Context } from './core/Context';
import { LogManager } from './core/LogManager';
import CommandRegistry from './core/command/CommandRegistry';
import { InputManager } from './core/input/InputManager';
import { ProfileManager } from './core/profile/ProfileManager';
import { Server } from './core/server/Server';
import UIManager from './core/ui/UIManager';
import { StateManager } from './state/StateManager';
import type { ApplicationState } from './types/State';

export class Application {
	public ui: UIManager;
	public cmdRegistry: CommandRegistry;
	public input: InputManager;
	public state: StateManager<ApplicationState>;
	public profiles: ProfileManager;
	public logs: LogManager;
	public context: Context;

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
		this.context = new Context(this);

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
				for (const _profile of profiles) {
					const profile = this.profiles.createProfile(
						_profile.name,
						_profile.email,
						_profile.password,
						_profile.active,
						_profile.id,
					);

					for (const server of _profile.servers) {
						profile.servers.createServer(
							server.ip,
							server.port,
							new HTTPAdapter({
								host: server.ip,
								port: server.port,
								secure: false,
							}),
							server.active,
							server.accountCreated,
						);
					}
				}
			}
		} catch {
			return this.logs.send('error', 'Failed to read in-memory profiles');
		}
	}
}

export const vonchat = new Application();
