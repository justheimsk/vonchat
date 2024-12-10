import type { State } from './State';

// biome-ignore lint/suspicious/noExplicitAny: <explanation>
export type Reducers = { [key: string]: State<any> };

export class StateManager<T = Reducers> {
	public reducers: T;

	public constructor(reducers: T) {
		this.reducers = reducers;
	}

	// biome-ignore lint/suspicious/noExplicitAny: <explanation>
	public dispatch(state: State<any>) {
		state.notify(state.data);
	}
}
