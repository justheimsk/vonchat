import { vonchat } from '@/lib/Application';
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
	public register(
		name: string,
		description: string,
		args: Arg[],
		execv: CommandCallback,
	) {
		const command = new Command(name, description, args, execv);
		//store.dispatch(registerCommand(command));
		vonchat.state.dispatch(
			vonchat.state.reducers.cmdRegistry.appendCommand(command),
		);
	}

	public fetch(name: string) {
		//return this.getState().commands.find((cmd) => cmd.name === name);
		return vonchat.state.reducers.cmdRegistry.data.commands.get(name);
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
			cmd.execv(ctx);
		}
	}

	public getState() {
		return vonchat.state.reducers.cmdRegistry;
	}
}
