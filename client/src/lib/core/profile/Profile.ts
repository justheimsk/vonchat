import type { BackendAdapter } from '../BackendAdapter';

export class Profile {
	public name: string;
	public email: string;
	public password: string;
	public id;
	public adapter: BackendAdapter;

	public constructor(
		name: string,
		email: string,
		password: string,
		adapter: BackendAdapter,
	) {
		this.id = crypto.randomUUID();
		this.password = password;
		this.name = name;
		this.email = email;

		this.adapter = adapter;
		this.adapter.attachProfile(this);
	}
}
