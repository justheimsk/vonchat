import type { Server } from './Server';

export class ServerManager extends Map<string, Server> {
	public constructor(servers: Server[] = []) {
		super();

		if (servers && Array.isArray(servers)) {
			for (const server of servers) {
				this.set(server.host, server);
			}
		}
	}

	public getAll(): Server[] {
		return Array.from(this.values());
	}

	public getActive(): Server | undefined {
		return this.getAll().find((s) => s.active);
	}

	public setActiveServer(server: Server) {
		for (const _server of this.getAll()) {
			if (server.host === _server.host) _server.active = true;
			else _server.active = false;

			this.set(_server.host, _server);
		}
	}
}
