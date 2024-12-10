import { Observable } from '../core/Observable';

export abstract class State<T> extends Observable<T> {
	abstract get data(): T;
}
