import type { JSONServer } from '../../types/Server';
import type { BackendAdapter } from '../BackendAdapter';
import type { Profile } from '../profile/Profile';

export class Server {
	public host: string;
	public port: string;
	public adapter: BackendAdapter;
	public active: boolean;
	public profile?: Profile;

	public constructor(
		host: string,
		port: string,
		active: boolean,
		adapter: BackendAdapter,
	) {
		this.host = host;
		this.port = port;
		this.active = active;
		this.adapter = adapter;
	}

	public attach(profile: Profile) {
		this.profile = profile;
		this.adapter.attach(profile);
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
			adapter: this.adapter.adapterName,
			active: this.active,
		};
	}
}
