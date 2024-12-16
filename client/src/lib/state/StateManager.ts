import type { Reducers } from '../types/State';
import type { State } from './State';

export class StateManager<T = Reducers> {
	public reducers: T;

	public constructor(reducers: T) {
		this.reducers = reducers;
	}

	// biome-ignore lint/suspicious/noExplicitAny: <explanation>
	public dispatch(...state: State<any>[]) {
		for (const st of state) {
			st.notify(st.data);
		}
	}
}
