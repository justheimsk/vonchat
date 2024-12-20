import type { Application } from '@/lib/Application';
import { BackendAdapter } from '@/lib/core/BackendAdapter';
import type { LogManager } from '@/lib/core/LogManager';
import type { Profile } from '@/lib/core/profile/Profile';
import { to } from '@/utils/to';

export interface HTTPAdapterOptions {
	host: string;
	port: number | string;
	secure: boolean;
}

export class HTTPAdapter extends BackendAdapter {
	public options: HTTPAdapterOptions;
	private profile?: Profile;
	private app?: Application;
	private logger?: LogManager;

	public constructor(options: HTTPAdapterOptions) {
		super();

		this.options = options;
	}

	public get adapterName() {
		return 'HTTP';
	}

	public attach(profile: Profile, app: Application, logger: LogManager) {
		this.profile = profile;
		this.logger = logger;
		this.app = app;
	}

	public get address() {
		return `${this.options.secure ? 'https' : 'http'}://${this.options.host}:${this.options.port}`;
	}

	public async init() {
		if (!this.app || !this.profile || !this.logger)
			throw new Error('Profile, application and/or logger not attached.');

		this.logger.send('info', 'Initializing adapter...');
		this.checkHealth();
	}

	private async checkHealth() {
		this.logger?.send('info', 'Checking server health...');
		const [res, err] = await to(fetch(`${this.address}/`));
		if (err || !res) {
			this.logger?.send('info', 'Failed to contact server.');
			return false;
		}

		const json = await res.json();
		this.logger?.send(
			'info',
			`Server is healthy, running on version: ${json.version}`,
		);
		return true;
	}
}
