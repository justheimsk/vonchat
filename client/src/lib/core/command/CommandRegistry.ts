import type { Application } from '@/lib/Application';
import type {
	Arg,
	CommandCallback,
	RecvArg,
	RecvContext,
} from '@/lib/types/Command';
import type { LogManager } from '../LogManager';
import Command from './Command';

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
			} catch (err) {
				this.logs.send(
					'error',
					`Failed to execute command: ${cmd.name}: ${err}`,
				);
			}
		}
	}

	public getState() {
		return this.app.state.reducers.cmdRegistry;
	}
}
