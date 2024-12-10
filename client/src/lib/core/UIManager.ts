import { vonchat } from '../Application';

export default class UIManager {
	public getState() {
		return vonchat.state.reducers.ui.data;
	}

	public openCommandList() {
		vonchat.state.dispatch(vonchat.state.reducers.ui.toggleCommandList(true));
	}

	public closeCommandList() {
		vonchat.state.dispatch(vonchat.state.reducers.ui.toggleCommandList(false));
	}

	public selectCommand(name: string) {
		vonchat.state.dispatch(vonchat.state.reducers.ui.selectCommand(name));
	}
}
