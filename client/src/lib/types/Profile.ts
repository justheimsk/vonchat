import type { JSONServer } from './Server';

export interface JSONProfile {
	name: string;
	email: string;
	password: string;
	id: string;
	active: boolean;
	servers: JSONServer[];
}
