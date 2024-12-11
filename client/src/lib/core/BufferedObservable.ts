import { Observable } from './Observable';

export interface BufferedObservableOptions {
	flushInterval?: number;
	bufferMaxSize?: number;
}

export class BufferedObservable<T> extends Observable<T[]> {
	public FLUSH_INTERVAL = 500;
	public BUFFER_MAX_SIZE = 1000;
	public buffer: T[];
	public options: BufferedObservableOptions;

	public constructor(options?: BufferedObservableOptions) {
		super();

		this.options = this.validateOptions(options);
		this.buffer = [];
		setInterval(() => {
			this.flush();
		}, this.FLUSH_INTERVAL);
	}

	private validateOptions(_options?: BufferedObservableOptions) {
		const options: BufferedObservableOptions = _options || {};

		if (!options.bufferMaxSize || typeof options.bufferMaxSize !== 'number')
			options.bufferMaxSize = 1000;
		if (!options.flushInterval || typeof options.flushInterval !== 'number')
			options.flushInterval = 500;

		return options;
	}

	public notify(arg: T[]) {
		this.buffer.push(...arg);

		if (this.buffer.length >= this.BUFFER_MAX_SIZE) this.flush();
	}

	public flush() {
		if (this.buffer.length) super.notify(this.buffer);
		this.buffer = [];
	}
}
