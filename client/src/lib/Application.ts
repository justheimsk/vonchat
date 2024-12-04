import CommandRegistry from "./core/CommandRegistry";
import UIManager from "./core/UIManager"

class Application {
  public ui: UIManager;
  public cmdRegistry: CommandRegistry;

  public constructor() {
    this.ui = new UIManager();
    this.cmdRegistry = new CommandRegistry();
    this.loadClientCommands()
  }

  public async loadClientCommands() {
    const commands = (await import('../shared/commands')).default;
    commands();
  }
}

export const vonchat = new Application()
