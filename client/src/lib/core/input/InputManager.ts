import type { Application } from '@/lib/Application';
import { InputEvents } from '@/lib/events/InputEvents';
import type { Arg, RecvArg } from '@/lib/types/Command';
import type Command from '../command/Command';
import { InputHistory } from './InputHistory';

export class InputManager {
	public events: InputEvents;
	private app: Application;
	public value: string;
	public history: InputHistory;

	public constructor(app: Application) {
		this.events = new InputEvents();
		this.history = new InputHistory({ maxHistory: 100 });
		this.app = app;
		this.value = '';
	}

	public send(entry: string) {
		if (entry.startsWith('/')) {
			const selectedCommand = this.app.ui.getState().selectedCommand;
			const cmd = this.app.cmdRegistry.fetch(selectedCommand);
			const args: RecvArg[] = [];

			const regex = /([a-zA-Z0-9_-]+)="([^"]*)"/g;
			let match: RegExpExecArray | null;
			const test = entry.split(' ').slice(1).join(' ');

			// biome-ignore lint/suspicious/noAssignInExpressions: <explanation>
			while ((match = regex.exec(test)) !== null) {
				const key = match[1];
				const value = match[2].trim();

				if (!value) continue;
				args.push({ name: key, value });
			}

			if (cmd) {
				const requiredArgs = cmd.args.filter(
					(arg: Arg) => arg.required === true,
				);
				if (requiredArgs.length > args.length) {
					for (const _arg of requiredArgs) {
						const arg = args.find((arg) => arg.name === _arg.name);
						if (!arg || arg.value) return this.formatCommandInChatInput(cmd);
					}

					return this.formatCommandInChatInput(cmd);
				}

				this.events.domSetInnerText.notify('');
				this.app.cmdRegistry.exec(selectedCommand, args);
				this.app.ui.closeCommandList();
				this.history.pushHistory(entry);
				this.history.current = '';
				this.value = '';
			}
		}
	}

	public formatCommandInChatInput(cmd: Command) {
		return this.app.input.events.domSetInnerText.notify(
			`/${cmd.name} ${cmd.args.map((arg) => `${arg.name}=""`).join(' ')}`,
		);
	}

	public setValue(value: string, modifyDom = true) {
		// TODO: value is not always registering on change, so... make it work :)
		this.value = value;
		this.history.current = value;
		this.events.onInput.notify(value);
		if (modifyDom) this.events.domSetInnerText.notify(value);
	}
}
