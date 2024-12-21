import type { Application } from '../Application';

export class Context {
	private app: Application;

	public constructor(app: Application) {
		this.app = app;
	}

	public getActiveProfile() {
		return this.app.profiles.getActiveProfile();
	}

	public getActiveServer() {
		return this.getActiveProfile()?.servers.getActive();
	}

	public setActiveServer(host: string) {
		const profile = this.getActiveProfile();
		if (profile) {
			const found = profile.servers.get(host);
			if (found) {
				found.active = true;
				this.app.profiles
					.getActiveProfile()
					?.servers.createServer(found.host, found.port, found.adapter);
				found.connect();

				this.app.state.dispatch(
					this.app.state.reducers.profiles.addProfile(profile),
				);
			}
		}
	}
}
