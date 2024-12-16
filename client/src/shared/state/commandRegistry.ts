import type Command from '@/lib/core/command/Command';
import { State } from '@/lib/state/State';

export interface ICommandRegistryState {
	commands: Map<string, Command>;
}

export class CommandRegistryState extends State<ICommandRegistryState> {
	private commands = new Map<string, Command>();

	public get data() {
		return {
			commands: this.commands,
		};
	}

	public appendCommand(cmd: Command): this {
		this.commands.set(cmd.name, cmd);
		return this;
	}
}
