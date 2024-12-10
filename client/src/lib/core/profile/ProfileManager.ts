import type { Application } from '@/lib/Application';
import type { BackendAdapter } from '../BackendAdapter';
import { Profile } from './Profile';

export class ProfileManager {
	public app: Application;

	public constructor(app: Application) {
		this.app = app;
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
}
