import {
	BufferedObservable,
	type BufferedObservableOptions,
} from './observable/BufferedObservable';

export type LogLevel = 'info' | 'error' | 'warn' | 'debug';
export interface Log {
	level: LogLevel;
	message: string;
}

export type Output = (logs: Log[]) => void;

export interface LogManagerOptions {
	maxLogSize?: number;
	buffer?: BufferedObservableOptions;
}

// EXPERIMENTAL CODE!!!!!!!!!!!!!!!!!!!!!!!!!!
// This will probably have a huge overhead if there are many logs, or many instances.
// Another possible problem is data duplication, since each instance will store its logs, and the main instance will store the logs of all other instances.
export class LogManager extends BufferedObservable<Log> {
	public logs: Log[];
	public tag?: string;
	private output?: Output;
	private instances: LogManager[];
	private _options: LogManagerOptions;
	private startOffset = 0;

	public constructor(
		output?: Output,
		tag?: string,
		options?: LogManagerOptions,
	) {
		super(options?.buffer);

		this.instances = [];
		this.output = output;
		this.tag = tag;
		this._options = this.validateLoggerOptions(options);
		this.logs = [];

		this.subscribe((log) => this.output?.(log));
	}

	private validateLoggerOptions(_options?: LogManagerOptions) {
		const options = _options || {};
		if (!options.maxLogSize || typeof options.maxLogSize !== 'number')
			options.maxLogSize = 100000;

		return options;
	}

	public get logOptions(): LogManagerOptions {
		return this._options;
	}

	public updateOptions(options: LogManagerOptions) {
		this._options = options;
		this.iterate((inst) => inst.updateOptions(options));
	}

	private pushLog(log: Log) {
		if (
			this.logs.length - this.startOffset >=
			(this.logOptions.maxLogSize || 1000)
		) {
			//@ts-ignore
			this.logs[this.startOffset] = null;
			this.startOffset++;
			return;
		}

		this.logs[this.logs.length] = log;
	}

	public send(level: LogLevel, message: string) {
		const log = {
			level,
			message: `${this.tag ? `[${this.tag}]:` : ''} ${message}`,
		};

		this.pushLog(log);
		this.notify([log]);
	}

	public withTag(tag: string) {
		const instance = new LogManager(undefined, tag);
		this.instances.push(instance);

		instance.subscribe(async (log) => {
			for (const _log of log) {
				this.pushLog(_log);
			}

			this.notify(log);
		});

		return instance;
	}

	public setOutput(output: Output) {
		this.output = output;
	}

	public clear() {
		this.logs = [];
		this.startOffset = 0;

		this.iterate((inst) => inst.clear());
	}

	private iterate(cb: (instance: LogManager) => void) {
		for (const instance of this.instances) {
			cb(instance);
		}
	}
}
