import { BackendAdapter } from '@/lib/core/BackendAdapter';
import type { Profile } from '@/lib/core/profile/Profile';
import { to } from '@/utils/to';

export interface HTTPAdapterOptions {
	host: string;
	port: number | string;
	secure: boolean;
}

export class HTTPAdapter extends BackendAdapter {
	public options: HTTPAdapterOptions;
	public profile?: Profile;

	public constructor(options: HTTPAdapterOptions) {
		super();

		this.options = options;
	}

	public attachProfile(profile: Profile) {
		this.profile = profile;
	}

	public get address() {
		return `${this.options.secure ? 'https' : 'http'}://${this.options.host}:${this.options.port}`;
	}

	public async init() {
		const [ok, err] = await to(this.checkHealth());
		if (err) return;

		if (ok) await this.identify();
	}

	private async checkHealth() {
		const [_, err] = await to(fetch(`${this.address}/`));
		if (err) return false;

		return true;
	}

	private async identify() {
		if (!this.profile) return false;

		const [login, err] = await to(
			this.login(this.profile.email, this.profile.password),
		);
		if (err || !login) {
			const [register, err] = await to(
				this.register(
					this.profile.name,
					this.profile.email,
					this.profile.password,
				),
			);
			if (err || !register) return console.log(err);
		}

		return true;
	}

	private async login(email: string, password: string): Promise<string> {
		if (!email || !password) throw new Error('Missing email or password.');

		const res = await fetch(`${this.address}/v1/auth/login`, {
			method: 'POST',
			body: JSON.stringify({ email, password }),
		});

		return (await res.json()).token;
	}

	private async register(
		username: string,
		email: string,
		password: string,
	): Promise<string> {
		if (!username || !email || !password)
			throw new Error('Missing username, email or password.');

		const res = await fetch(`${this.address}/v1/auth/register`, {
			method: 'POST',
			body: JSON.stringify({
				username: username,
				password: password,
				email: email,
			}),
		});

		return (await res.json()).token;
	}
}
