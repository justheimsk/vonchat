import type { State } from '../state/State';

// biome-ignore lint/suspicious/noExplicitAny: <explanation>
export type Reducers = { [key: string]: State<any> };
