import type { Application } from '../Application';

export default class UIManager {
	private app: Application;

	public constructor(app: Application) {
		this.app = app;
	}

	public getState() {
		return this.app.state.reducers.ui.data;
	}

	public openCommandList() {
		this.app.state.dispatch(this.app.state.reducers.ui.toggleCommandList(true));
	}

	public closeCommandList() {
		this.app.state.dispatch(
			this.app.state.reducers.ui.toggleCommandList(false),
		);
	}

	public selectCommand(name: string) {
		this.app.state.dispatch(this.app.state.reducers.ui.selectCommand(name));
	}
}
