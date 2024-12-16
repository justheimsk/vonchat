import type { BufferedObservableOptions } from '@/lib/types/Observable';
import { Observable } from './Observable';

export class BufferedObservable<T> extends Observable<T[]> {
	public buffer: T[];
	public options: BufferedObservableOptions;

	public constructor(options?: Partial<BufferedObservableOptions>) {
		super();

		this.options = this.validateOptions(options);
		this.buffer = [];
		setInterval(() => {
			this.flush();
		}, this.options.flushInterval);
	}

	private validateOptions(_options?: Partial<BufferedObservableOptions>) {
		const options = {
			bufferMaxSize: 10000,
			flushInterval: 500,
		};

		if (_options?.bufferMaxSize && typeof _options.bufferMaxSize === 'number')
			options.bufferMaxSize = _options.bufferMaxSize;

		if (_options?.flushInterval && typeof _options.flushInterval === 'number')
			options.flushInterval = _options.flushInterval;

		return options;
	}

	public notify(arg: T[]) {
		this.buffer.push(...arg);

		if (this.buffer.length >= this.options.bufferMaxSize) this.flush();
	}

	public flush() {
		if (this.buffer.length) super.notify(this.buffer);
		this.buffer = [];
	}
}
