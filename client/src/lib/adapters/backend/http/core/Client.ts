import type { Application } from '@/lib/Application';
import type { LogManager } from '@/lib/core/LogManager';
import type { Profile } from '@/lib/core/profile/Profile';
import type { Server } from '@/lib/core/server/Server';
import type { HTTPAdapter } from '../HTTPAdapter';
import { WebSocketClient } from './Websocket';

export class Client {
	public adapter: HTTPAdapter;
	public logger: LogManager;
	public server: Server;
	public app: Application;
	public profile: Profile;
	public ws: WebSocketClient;
	private token = '';

	public constructor(
		adapter: HTTPAdapter,
		logger: LogManager,
		server: Server,
		app: Application,
		profile: Profile,
	) {
		this.adapter = adapter;
		this.server = server;
		this.app = app;
		this.profile = profile;
		this.logger = logger;
		this.ws = new WebSocketClient(this);
	}

	public async init() {
		if (this.server.accountCreated) {
			this.logger.send('info', 'Trying to login...');
			this.token = await this.login();
		} else {
			this.logger.send('info', 'Trying to create a account...');
			this.token = await this.createAccount();
			this.server.accountCreated = true;
			this.app.profiles.saveToMemory();
		}

		if (this.token) {
			this.logger.send('info', 'Authentication : OK');
			this.ws.connect(this.token);
			this.load();
		}
	}

	public async load() {
		this.getUsers();
	}

	public async getUsers() {
		try {
			const res = await fetch(`${this.adapter.address}/v1/users`, {
				headers: {
					Authorization: this.token,
				},
			});

			const json = await res.json();
			for (const user of json.users) {
				this.adapter.state.dispatch(
					this.adapter.state.reducers.users.pushUser(user),
				);
			}
		} catch (err) {
			this.logger.send('error', `Failed to get users: ${err}`);
		}
	}

	public async createAccount() {
		const res = await fetch(
			`${this.adapter.address}/v${this.adapter.version}/auth/register`,
			{
				method: 'post',
				body: JSON.stringify({
					username: this.adapter.profile?.name,
					email: this.adapter.profile?.email,
					password: this.adapter.profile?.password,
				}),
			},
		);
		const json = await res.json();

		return json.token;
	}

	public async login() {
		const res = await fetch(
			`${this.adapter.address}/v${this.adapter.version}/auth/login`,
			{
				method: 'post',
				body: JSON.stringify({
					email: this.adapter.profile?.email,
					password: this.adapter.profile?.password,
				}),
			},
		);

		const json = await res.json();
		return json.token;
	}
}
