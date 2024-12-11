import { Observable } from './Observable';

export type LogLevel = 'info' | 'error' | 'debug';
export interface Log {
	level: LogLevel;
	message: string;
}

export type Output = (log: Log) => void;

// EXPERIMENTAL CODE!!!!!!!!!!!!!!!!!!!!!!!!!!
// This will probably have a huge overhead if there are many logs, or many instances.
// Another possible problem is data duplication, since each instance will store its logs, and the main instance will store the logs of all other instances.
export class LogManager extends Observable<Log> {
	public logs: Log[];
	public tag?: string;
	private output?: Output;
	private instances: LogManager[];

	public constructor(output?: Output, tag?: string) {
		super();

		this.instances = [];
		this.output = output;
		this.tag = tag;
		this.logs = [];
	}

	public send(level: LogLevel, message: string) {
		const log = {
			level,
			message: `${this.tag ? `[${this.tag}]:` : ''} ${message}`,
		};
		this.logs.push(log);
		this.notify(log);

		if (this.output) this.output(log);
	}

	public withTag(tag: string) {
		const instance = new LogManager(this.output, tag);
		this.instances.push(instance);

		instance.subscribe((log) => {
			this.logs.push(log);
			this.notify(log);
		});

		return instance;
	}

	public setOutput(output: Output) {
		this.iterate((inst) => inst.setOutput(output));
	}

	public clear() {
		this.logs = [];

		this.iterate((inst) => inst.clear());
	}

	private iterate(cb: (instance: LogManager) => void) {
		for (const instance of this.instances) {
			cb(instance);
		}
	}
}
