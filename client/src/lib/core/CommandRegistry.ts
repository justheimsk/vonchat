import store from "@/store/store";
import Command from "./Command";
import {registerCommand} from "@/store/slices/commandRegistry";

export default class CommandRegistry {
  public register(name: string, description: string, execv: (...args: unknown[]) => void) {
    const command = new Command(name, description, execv);
    store.dispatch(registerCommand(command));
  }
}
