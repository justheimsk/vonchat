import CommandRegistry from "./core/CommandRegistry";
import {Input} from "./core/Input";
import UIManager from "./core/UIManager"

export class Application {
  public ui: UIManager;
  public cmdRegistry: CommandRegistry;
  public input: Input;

  public constructor() {
    this.ui = new UIManager();
    this.cmdRegistry = new CommandRegistry();
    this.input = new Input(this);

    this.loadClientCommands()
  }

  private async loadClientCommands() {
    const commands = (await import('../shared/commands')).default;
    commands();
  }
}

export const vonchat = new Application()
