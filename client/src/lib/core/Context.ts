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
}
