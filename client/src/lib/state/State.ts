import { Observable } from '../core/observable/Observable';

export abstract class State<T> extends Observable<T> {
	abstract get data(): T;
}
