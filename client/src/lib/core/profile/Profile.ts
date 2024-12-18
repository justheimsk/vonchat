import type { Server } from '@/lib/core/server/Server';
import type { JSONProfile } from '@/lib/types/Profile';
import { ServerManager } from '../server/ServerManager';

export class Profile {
	public name: string;
	public email: string;
	public password: string;
	public id;
	public active: boolean;
	public servers: ServerManager;

	public constructor(
		name: string,
		email: string,
		password: string,
		active?: boolean,
		id?: string,
		servers: Server[] = [],
	) {
		this.id = id || crypto.randomUUID();
		this.password = password;
		this.name = name;
		this.email = email;
		this.active = active || false;
		this.servers = new ServerManager(servers);
	}

	public toJSON(): JSONProfile {
		return {
			servers: this.servers.getAll().map((s) => s.toJSON()),
			active: this.active,
			id: this.id,
			name: this.name,
			email: this.email,
			password: this.password,
		};
	}
}
