import { GlobalAdapter } from './GlobalAdapter';

export abstract class MemoryAdapter extends GlobalAdapter {
	abstract get<T>(key: string): T | null;
	abstract set<T>(key: string, value: T): void;
}
