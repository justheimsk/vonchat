import type { InputManager } from './InputManager';

export interface InputHistoryOptions {
	maxHistory: number;
}

export class InputHistory {
	private historyIdx: number;
	private history: string[];
	private input: InputManager;
	public options: InputHistoryOptions;

	public constructor(input: InputManager, options: InputHistoryOptions) {
		this.input = input;
		this.historyIdx = 0;
		this.history = [];
		this.options = this.validateOptions(options);
	}

	public pushHistory(value: string) {
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
			return this.input.value;
		}
	}

	private validateOptions(options: InputHistoryOptions) {
		if (
			!options.maxHistory ||
			typeof options.maxHistory !== 'number' ||
			options.maxHistory < 0
		) {
			options.maxHistory = 100;
		}

		return options;
	}
}
