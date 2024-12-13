import type { Application } from '@/lib/Application';
import type { LogManager } from '../LogManager';
import Command, { type Arg } from './Command';

export type CommandCallback = (ctx: RecvContext) => void;

export interface RecvArg {
	name: string;
	value: unknown;
}

export interface RecvContext {
	args: Map<string, RecvArg>;
}

export default class CommandRegistry {
	private app: Application;
	private logs: LogManager;

	public constructor(app: Application, logs: LogManager) {
		this.app = app;
		this.logs = logs;
	}

	public register(
		name: string,
		description: string,
		args: Arg[],
		execv: CommandCallback,
	) {
		if (this.getState().data.commands.get(name)) {
			this.logs.send(
				'warn',
				`Trying to register a duplicated command: ${name}`,
			);
			return;
		}

		const command = new Command(name, description, args, execv);
		this.app.state.dispatch(
			this.app.state.reducers.cmdRegistry.appendCommand(command),
		);

		this.logs.send('info', `Registering new command: ${command.name}`);
	}

	public fetch(name: string) {
		return this.app.state.reducers.cmdRegistry.data.commands.get(name);
	}

	public exec(name: string, args: RecvArg[]) {
		const cmd = this.fetch(name);
		const ctx: RecvContext = {
			args: new Map(),
		};

		for (const arg of args) {
			ctx.args.set(arg.name, arg);
		}

		if (cmd) {
			try {
				cmd.execv(ctx);
			} catch {
				this.logs.send('error', `Failed to execute command: ${cmd.name}`);
			}
		}
	}

	public getState() {
		return this.app.state.reducers.cmdRegistry;
	}
}
