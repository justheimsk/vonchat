import { UserState } from '@/lib/core/adapter/backend/BackendAdapterState';
import type { APIUser } from '@/lib/types/api/common/APIUser';

export class HTTPUserState extends UserState {
	private values: APIUser[] = [];

	public get data() {
		return {
			values: this.values,
		};
	}

	public pushUser(user: APIUser) {
		this.values.push(user);
		return this;
	}
}
