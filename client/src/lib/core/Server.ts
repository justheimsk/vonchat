import type { JSONServer } from '../types/Server';
import type { BackendAdapter } from './BackendAdapter';

export class Server {
	public ip: string;
	public port: string;
	public adapter: BackendAdapter;

	public constructor(ip: string, port: string, adapter: BackendAdapter) {
		this.ip = ip;
		this.port = port;
		this.adapter = adapter;
	}

	public toJSON(): JSONServer {
		return {
			ip: this.ip,
			port: this.port,
			adapter: this.adapter.adapterName,
		};
	}
}
