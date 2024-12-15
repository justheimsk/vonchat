import type { Application } from '@/lib/Application';
import type { LogManager } from '../LogManager';
import type { MemoryAdapter } from '../MemoryAdapter';
import type { Server } from '../Server';
import { type JSONProfile, Profile } from './Profile';

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
			this.app.state.reducers.profiles.appendProfile(profile),
		);

		if (active) this.setActiveProfile(profile);
		return profile;
	}

	public addServer(profile: Profile, server: Server) {
		profile.servers.push(server);
		this.app.state.dispatch(
			this.app.state.reducers.profiles.appendProfile(profile),
		);
	}

	public getProfiles() {
		return this.app.state.reducers.profiles.data.profiles;
	}

	public getActiveProfile(
		profiles?: Map<string, Profile>,
	): Profile | undefined {
		return Array.from(profiles?.values() || this.getProfiles().values()).find(
			(p) => p.active,
		);
	}

	public setActiveProfile(profile: Profile) {
		for (const _profile of this.getProfiles().values()) {
			_profile.active = _profile.id === profile.id;
			this.app.state.dispatch(
				this.app.state.reducers.profiles.appendProfile(_profile),
			);
		}
	}

	public readInMemoryProfiles() {
		return this.memory.get<JSONProfile[]>('profiles');
	}

	public saveToMemory() {
		try {
			const profiles = this.app.state.reducers.profiles.data.profiles;
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
