import type { Arg, CommandCallback, RecvContext } from '@/lib/types/Command';

export default class Command {
	public name: string;
	public description: string;
	public execv: CommandCallback;
	public args: Arg[];

	public constructor(
		name: string,
		description: string,
		args: Arg[],
		execv: CommandCallback,
	) {
		this.name = name;
		this.description = description;
		this.execv = execv;
		this.args = args;
	}

	public execute(ctx: RecvContext) {
		this.execv(ctx);
	}
}
