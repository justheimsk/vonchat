import type { Application } from '@/lib/Application';
import type { BackendAdapter } from '../adapter/backend/BackendAdapter';
import type { Profile } from '../profile/Profile';
import { Server } from './Server';

export class ServerManager extends Map<string, Server> {
	private app: Application;
	public profile: Profile;

	public constructor(
		app: Application,
		profile: Profile,
		servers: Server[] = [],
	) {
		super();

		this.app = app;
		this.profile = profile;

		if (servers && Array.isArray(servers)) {
			for (const server of servers) {
				this.set(server.host, server);
			}
		}
	}

	public createServer(
		host: string,
		port: string,
		adapter: BackendAdapter,
		active = false,
		accountCreated = false,
	): Server {
		const server = new Server(
			this.app,
			this.profile,
			host,
			port,
			active,
			accountCreated,
			adapter,
		);
		this.set(host, server);
		if (active) this.setActiveServer(server);

		this.dispatch();
		return server;
	}

	public getAll(): Server[] {
		return Array.from(this.values());
	}

	public getActive(): Server | undefined {
		return this.getAll().find((s) => s.active);
	}

	private dispatch() {
		this.app.state.dispatch(
			this.app.state.reducers.profiles.addProfile(this.profile),
		);
	}

	public setActiveServer(server: Server) {
		for (const _server of this.getAll()) {
			if (server.host === _server.host) _server.active = true;
			else _server.active = false;

			this.set(_server.host, _server);
		}

		this.dispatch();
	}
}
