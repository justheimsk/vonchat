import { MemoryAdapter } from '../core/adapter/MemoryAdapter';

export class LocalStorageMemoryAdapter extends MemoryAdapter {
	public get<T>(key: string): T | null {
		const value = localStorage.getItem(key);
		if (!value) return null;

		return JSON.parse(value) as T;
	}

	public set<T>(key: string, value: T): void {
		const stringified = JSON.stringify(value);
		localStorage.setItem(key, stringified);
	}

	public get adapterName() {
		return 'Local Storage';
	}
}
