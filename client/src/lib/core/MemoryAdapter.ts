export abstract class MemoryAdapter {
	abstract get<T>(key: string): T | null;
	abstract set<T>(key: string, value: T): void;
}
