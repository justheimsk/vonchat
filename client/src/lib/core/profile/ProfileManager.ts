import type { Application } from '@/lib/Application';
import type { JSONProfile } from '@/lib/types/Profile';
import type { LogManager } from '../LogManager';
import type { MemoryAdapter } from '../MemoryAdapter';
import type { Server } from '../server/Server';
import { Profile } from './Profile';

export class ProfileManager {
	public app: Application;
	public memory: MemoryAdapter;
	public logs: LogManager;

	public constructor(
		app: Application,
		memory: MemoryAdapter,
		logs: LogManager,
	) {
		this.app = app;
		this.memory = memory;
		this.logs = logs;
	}

	public createProfile(
		name: string,
		email: string,
		password: string,
		active?: boolean,
		id?: string,
		servers: Server[] = [],
	) {
		const profile = new Profile(name, email, password, active, id, servers);

		this.app.state.dispatch(
			this.app.state.reducers.profiles.addProfile(profile),
		);

		if (active) this.setActiveProfile(profile);
		return profile;
	}

	public addServer(profile: Profile, server: Server) {
		profile.servers.set(server.host, server);
		server.attach(profile);
		if (server.active) profile.servers.setActiveServer(server);

		this.app.state.dispatch(
			this.app.state.reducers.profiles.addProfile(profile),
		);
	}

	public getState() {
		return this.app.state.reducers.profiles;
	}

	public getActiveProfile(): Profile | undefined {
		return this.getState().data.activeProfile;
	}

	public setActiveProfile(profile: Profile) {
		for (const _profile of this.getState().data.profiles.values()) {
			if (_profile.id === profile.id) {
				_profile.active = true;
				this.app.state.dispatch(
					this.app.state.reducers.profiles.setActiveProfile(_profile),
				);
			} else {
				_profile.active = false;
			}

			this.app.state.dispatch(
				this.app.state.reducers.profiles.addProfile(_profile),
			);
		}
	}

	public readInMemoryProfiles() {
		return this.memory.get<JSONProfile[]>('profiles');
	}

	public saveToMemory() {
		try {
			const profiles = this.app.state.reducers.profiles.data.profiles;
			console.log(profiles);
			this.memory.set(
				'profiles',
				Array.from(profiles.values()).map((profile) => profile.toJSON()),
			);
			this.logs.send(
				'info',
				`Saved ${profiles.size} profiles using ${this.memory.adapterName} adapter.`,
			);
		} catch (err) {
			this.logs.send(
				'error',
				`Failed to write profiles in memory, using: ${this.memory.adapterName} adapter: ${err}`,
			);
		}
	}
}
