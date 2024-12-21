import type { Application } from '@/lib/Application';
import type { StateManager } from '@/lib/state/StateManager';
import type { LogManager } from '../../LogManager';
import type { Profile } from '../../profile/Profile';
import type { Server } from '../../server/Server';
import { GlobalAdapter } from './../GlobalAdapter';
import type { BackendAdapterState } from './BackendAdapterState';

export abstract class BackendAdapter extends GlobalAdapter {
	abstract get state(): StateManager<BackendAdapterState>;
	abstract get address(): string;
	abstract get version(): string;
	abstract init(): void;
	abstract attach(
		profile: Profile,
		server: Server,
		app: Application,
		logger: LogManager,
	): void;
}
