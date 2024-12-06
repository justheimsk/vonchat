import store from "@/store/store";
import Command, {type Arg} from "./Command";
import {registerCommand} from "@/store/slices/commandRegistry";

export type CommandCallback = (ctx: RecvContext) => void;

export interface RecvArg {
  name: string;
  value: unknown;
}
export interface RecvContext {
  args: Map<string, RecvArg>
}

export default class CommandRegistry {
  public register(name: string, description: string, args: Arg[], execv: CommandCallback) {
    const command = new Command(name, description, args, execv);
    store.dispatch(registerCommand(command));
  }

  public fetch(name: string) {
    return this.getState().commands.find((cmd) => cmd.name === name);
  }

  public exec(name: string, args: RecvArg[]) {
    const cmd = this.getState().commands.find((cmd) => cmd.name === name);
    const ctx: RecvContext = {
      args: new Map()
    }

    for(const arg of args) {
      ctx.args.set(arg.name, arg);
    }

    if(cmd) {
      cmd.execv(ctx);
    }
  }

  public getState() {
    return store.getState().commandRegistry;
  }
}
