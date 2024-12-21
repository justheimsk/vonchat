import type { Application } from '@/lib/Application';
import type { JSONServer, ServerStatus } from '../../types/Server';
import type { LogManager } from '../LogManager';
import type { BackendAdapter } from '../adapter/backend/BackendAdapter';
import type { Profile } from '../profile/Profile';

export class Server {
	public host: string;
	public port: string;
	public adapter: BackendAdapter;
	public active: boolean;
	public profile?: Profile;
	public status: ServerStatus = 'disconnected';
	public accountCreated: boolean;

	public constructor(
		host: string,
		port: string,
		active: boolean,
		accountCreated: boolean,
		adapter: BackendAdapter,
	) {
		this.host = host;
		this.port = port;
		this.active = active;
		this.accountCreated = accountCreated;
		this.adapter = adapter;
	}

	public attach(profile: Profile, app: Application, logger: LogManager) {
		this.profile = profile;
		this.adapter.attach(profile, this, app, logger);
	}

	public connect() {
		if (!this.profile)
			throw new Error(
				`Profile is not attached to server: ${this.host}:${this.port}`,
			);
		this.adapter.init();
	}

	public toJSON(): JSONServer {
		return {
			ip: this.host,
			port: this.port,
			active: this.active,
			adapter: this.adapter.adapterName,
			accountCreated: this.accountCreated,
		};
	}
}
