export interface InputHistoryOptions {
	maxHistory: number;
	allowDuplicatedHistory?: boolean;
}

export class InputHistory {
	private historyIdx: number;
	private history: string[];
	public options: InputHistoryOptions;
	public current: string;

	public constructor(options: InputHistoryOptions) {
		this.historyIdx = 0;
		this.history = [];
		this.options = this.validateOptions(options);
		this.current = '';
	}

	public pushHistory(value: string) {
		if (
			this.options.allowDuplicatedHistory === false &&
			this.history[this.history.length - 1] === value
		)
			return;

		this.history.push(value);
		this.resetIdx();
	}

	public resetIdx() {
		this.historyIdx = this.history.length;
	}

	public getPrevious() {
		if (this.historyIdx > 0) {
			this.historyIdx--;
			const value = this.history[this.historyIdx];

			return value;
		}
	}

	public getNext() {
		if (this.historyIdx < this.history.length - 1) {
			this.historyIdx++;
			const value = this.history[this.historyIdx];

			return value;
			// biome-ignore lint/style/noUselessElse: <explanation>
		} else {
			this.resetIdx();
			return this.current;
		}
	}

	private validateOptions(options: InputHistoryOptions) {
		if (
			!options.maxHistory ||
			typeof options.maxHistory !== 'number' ||
			options.maxHistory < 0
		)
			options.maxHistory = 100;

		if (
			options.allowDuplicatedHistory === undefined ||
			typeof options.allowDuplicatedHistory !== 'boolean'
		)
			options.allowDuplicatedHistory = false;
		return options;
	}
}
