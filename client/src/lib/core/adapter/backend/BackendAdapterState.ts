import { State } from '@/lib/state/State';
import type { APIUser } from '@/lib/types/api/common/APIUser';

export interface UserStateData {
	values: APIUser[];
}

export abstract class UserState extends State<UserStateData> {
	abstract pushUser(user: APIUser): this;
}

export interface BackendAdapterState {
	users: UserState;
}
