import type { Application } from '@/lib/Application';
import type { BackendAdapter } from '../BackendAdapter';
import type { MemoryAdapter } from '../MemoryAdapter';
import { type JSONProfile, Profile } from './Profile';

export class ProfileManager {
	public app: Application;
	public memory: MemoryAdapter;

	public constructor(app: Application, memory: MemoryAdapter) {
		this.app = app;
		this.memory = memory;
	}

	public createProfile(
		name: string,
		email: string,
		password: string,
		adapter: BackendAdapter,
	) {
		const profile = new Profile(name, email, password, adapter);
		this.app.state.dispatch(
			this.app.state.reducers.profiles.appendProfile(profile),
		);

		return profile;
	}

	public getProfiles() {
		return this.app.state.reducers.profiles.data;
	}

	public readInMemoryProfiles() {
		return this.memory.get<JSONProfile[]>('profiles');
	}

	public saveToMemory() {
		const profiles = this.app.state.reducers.profiles.data.profiles;
		this.memory.set(
			'profiles',
			Array.from(profiles.values()).map((profile) => profile.toJSON()),
		);
	}
}
