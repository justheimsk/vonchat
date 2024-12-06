import type {Application} from "../Application";
import {InputEvents} from "../events/InputEvents";
import type Command from "./Command";
import type {RecvArg} from "./CommandRegistry";

export class Input {
  public events: InputEvents;
  private app: Application;

  public constructor(app: Application) {
    this.events = new InputEvents();
    this.app = app;
  }

  public send(entry: string) {
    if(entry.startsWith("/")) {
      const selectedCommand = this.app.ui.getState().selectedCommand;
      const cmd = this.app.cmdRegistry.fetch(selectedCommand);
      const args: RecvArg[] = [];

      const regex = /([a-zA-Z0-9_-]+)="([^"]*)"/g;
      let match: RegExpExecArray | null;
      const test = entry.split(" ").slice(1).join(" ");

      // biome-ignore lint/suspicious/noAssignInExpressions: <explanation>
      while ((match = regex.exec(test)) !== null) {
        const key = match[1];
        const value = match[2].trim();

        if (!value) continue;
        args.push({ name: key, value });
      }

      if(cmd) {
        const requiredArgs = cmd.args.filter((arg) => arg.required === true)
        if(requiredArgs.length > args.length) {
          for(const _arg of requiredArgs) {
            const arg = args.find((arg) => arg.name === _arg.name);
            if(!arg || arg.value) return this.formatCommandInChatInput(cmd);
          }

          return this.formatCommandInChatInput(cmd);
        }

        this.events.clearChatInput.notify(null);
        this.app.cmdRegistry.exec(selectedCommand, args);
        this.app.ui.closeCommandList();
      }
    }
  }

  public formatCommandInChatInput(cmd: Command) {
    return this.app.input.events.setChatInput.notify(`/${cmd.name} ${cmd.args.map((arg) => `${arg.name}=""`).join(" ")}`);
  }
}
