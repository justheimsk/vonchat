import {selectCommand, toggleCommandList} from "@/store/slices/ui";
import store from "@/store/store";

export default class UIManager {
  public openCommandList() {
    store.dispatch(toggleCommandList(true));
  }

  public closeCommandList() {
    store.dispatch(toggleCommandList(false));
  }

  public selectCommand(name: string) {
    store.dispatch(selectCommand(name));
  }
}
