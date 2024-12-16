export type ArgType = 'text';
export type CommandCallback = (ctx: RecvContext) => void;

export interface Arg {
	type: ArgType;
	name: string;
	required?: boolean;
}

export interface RecvArg {
	name: string;
	value: unknown;
}

export interface RecvContext {
	args: Map<string, RecvArg>;
}
