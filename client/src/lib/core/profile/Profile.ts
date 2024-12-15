export interface JSONProfile {
	name: string;
	email: string;
	password: string;
	id: string;
}

export class Profile {
	public name: string;
	public email: string;
	public password: string;
	public id;

	public constructor(name: string, email: string, password: string) {
		this.id = crypto.randomUUID();
		this.password = password;
		this.name = name;
		this.email = email;
	}

	public toJSON(): JSONProfile {
		return {
			id: this.id,
			name: this.name,
			email: this.email,
			password: this.password,
		};
	}
}
