import type { LogManager } from '@/lib/core/LogManager';
import type { WSMessage } from '@/lib/types/api/ws/WSMessage';
import { OPCODES } from '@/lib/types/api/ws/opcodes';
import type { Client } from './Client';

export class WebSocketClient {
	private token?: string;
	private client: Client;
	public address: string;
	public logger: LogManager;
	public socket?: WebSocket;

	public constructor(client: Client) {
		this.address = client.adapter.address;
		this.client = client;
		this.logger = client.logger;
	}

	public connect(token: string) {
		this.token = token;
		this.socket = new WebSocket(`${this.address}/ws`);
		this.socket.onopen = () => this.onOpen();
		this.socket.onclose = () => this.onClose();
	}

	private onClose() {
		setTimeout(() => {
			if (this.token) this.connect(this.token);
		}, 5000);

		this.client.server.status = 'disconnected';
		this.logger.send(
			'info',
			`Connection closed with ${this.address}, trying to connect in 5 seconds...`,
		);
	}

	private onOpen() {
		this.client.server.status = 'connected';
		this.logger.send('info', `Connection opened with ${this.address}`);
		this.identify();
	}

	public send(data: WSMessage) {
		this.socket?.send(JSON.stringify(data));
	}

	private identify() {
		this.logger.send('info', 'Sending identify packet...');
		this.send({ op: OPCODES.IDENTIFY, d: { token: this.token } });
	}
}
