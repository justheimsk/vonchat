import type { Profile } from './profile/Profile';

export abstract class BackendAdapter {
	abstract get adapterName(): string;
	abstract get address(): string;
	abstract init(): void;
	abstract attachProfile(profile: Profile): void;
}