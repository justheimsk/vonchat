import type { JSONProfile } from '@/lib/types/Profile';
import type { Server } from '../Server';

export class Profile {
	public name: string;
	public email: string;
	public password: string;
	public id;
	public active: boolean;
	public servers: Server[];

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
		this.servers = servers;
	}

	public toJSON(): JSONProfile {
		return {
			servers: this.servers.map((s) => s.toJSON()),
			active: this.active,
			id: this.id,
			name: this.name,
			email: this.email,
			password: this.password,
		};
	}
}
