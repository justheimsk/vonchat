import type { Application } from '../Application';
import { GlobalAdapter } from './GlobalAdapter';
import type { LogManager } from './LogManager';
import type { Profile } from './profile/Profile';

export abstract class BackendAdapter extends GlobalAdapter {
	abstract get address(): string;
	abstract init(): void;
	abstract attach(profile: Profile, app: Application, logger: LogManager): void;
}
