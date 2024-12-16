import type { BufferedObservableOptions } from './Observable';

export type LogLevel = 'info' | 'error' | 'warn' | 'debug';
export type Output = (logs: Log[]) => void;

export interface Log {
	level: LogLevel;
	message: string;
}

export interface LogManagerOptions {
	maxLogSize?: number;
	buffer?: BufferedObservableOptions;
}
