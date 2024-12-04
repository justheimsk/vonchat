export default class Command {
  public name: string;
  public description: string;
  public execv: (...args: unknown[]) => void;

  public constructor(name: string, description: string, execv: (...args: unknown[]) => void) {
    this.name = name;
    this.description = description;
    this.execv = execv;
  }

  public execute() {
    this.execv();
  }
}
