import type { Application } from '@/lib/Application';
import type { JSONServer, ServerStatus } from '../../types/Server';
import type { LogManager } from '../LogManager';
import type { BackendAdapter } from '../adapter/backend/BackendAdapter';
import type { Profile } from '../profile/Profile';

export class Server {
  public app: Application;
	public host: string;
	public port: string;
	public adapter: BackendAdapter;
	public active: boolean;
	public profile?: Profile;
	public status: ServerStatus = 'disconnected';
	public accountCreated: boolean;

	public constructor(
    app: Application,
    profile: Profile,
		host: string,
		port: string,
		active: boolean,
		accountCreated: boolean,
		adapter: BackendAdapter,
	) {
    this.app = app;
		this.host = host;
		this.port = port;
		this.active = active;
		this.accountCreated = accountCreated;
    this.profile = profile;

		this.adapter = adapter;
		this.adapter.attach(profile, this, app, this.app.logs.withTag(adapter.adapterName));
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
