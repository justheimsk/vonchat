export type UserStatus = 'online' | 'offline' | 'dnd' | 'idle';

export interface APIUser {
	id: string;
	username: string;
	status: UserStatus;
}
