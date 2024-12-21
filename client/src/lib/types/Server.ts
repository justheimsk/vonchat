export interface JSONServer {
	ip: string;
	port: string;
	adapter: string;
	active: boolean;
	accountCreated: boolean;
}

export type ServerStatus =
	| 'connecting'
	| 'connected'
	| 'failed'
	| 'disconnected';
