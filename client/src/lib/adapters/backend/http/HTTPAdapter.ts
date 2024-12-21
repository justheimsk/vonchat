import type { Application } from '@/lib/Application';
import type { LogManager } from '@/lib/core/LogManager';
import { BackendAdapter } from '@/lib/core/adapter/backend/BackendAdapter';
import type { BackendAdapterState } from '@/lib/core/adapter/backend/BackendAdapterState';
import type { Profile } from '@/lib/core/profile/Profile';
import type { Server } from '@/lib/core/server/Server';
import { StateManager } from '@/lib/state/StateManager';
import type { HealthCheckResponse } from '@/lib/types/api/common/APIResponses';
import { Client } from './core/Client';
import { HTTPUserState } from './state/user';

export interface HTTPAdapterOptions {
	host: string;
	port: number | string;
	secure: boolean;
}

export class HTTPAdapter extends BackendAdapter {
	public options: HTTPAdapterOptions;
	public profile?: Profile;
	public app?: Application;
	public logger?: LogManager;
	public server?: Server;
	private _version?: string;
	public state: StateManager<BackendAdapterState>;

	public constructor(options: HTTPAdapterOptions) {
		super();

		this.state = new StateManager({
			users: new HTTPUserState(),
		});
		this.options = options;
	}

	public get version() {
		return this._version || 'unknown';
	}

	public get adapterName() {
		return 'HTTP';
	}

	public attach(
		profile: Profile,
		server: Server,
		app: Application,
		logger: LogManager,
	) {
		this.profile = profile;
		this.server = server;
		this.logger = logger;
		this.app = app;
	}

	public get address() {
		return `${this.options.secure ? 'https' : 'http'}://${this.options.host}:${this.options.port}`;
	}

	public async init() {
		if (!this.app || !this.profile || !this.server || !this.logger)
			throw new Error('Profile, application and/or logger not attached.');

		this.server.status = 'connecting';
		this.logger.send('info', 'Initializing adapter...');

		try {
			this.logger.send('info', 'Checking server health...');

			const health = await this.checkHealth();
			this._version = health.version;

			this.logger.send(
				'info',
				`Server is healthy, running on version: ${health.version}`,
			);

			const client = new Client(
				this,
				this.logger.withTag(this.address),
				this.server,
				this.app,
				this.profile,
			);
			client.init();
		} catch (err) {
			this.logger.send('info', `Server is unhealthy: ${err}`);
		}
	}

	private async checkHealth(): Promise<HealthCheckResponse> {
		const res = await fetch(`${this.address}/`);
		const json = await res.json();

		return { version: json.version, message: json.message };
	}
}
