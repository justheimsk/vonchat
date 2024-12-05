import {Observable} from "../core/Observable";

export class InputEvents {
  public clearChatInput = new Observable<null>;
  public setChatInput = new Observable<string>;
  public appendChatInput = new Observable<string>;
}
