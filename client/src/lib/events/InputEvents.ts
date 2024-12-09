import { Observable } from '../core/Observable';

export class InputEvents {
	public domSetInnerText = new Observable<string>();
	public onInput = new Observable<string>();
	public onKeyDown = new Observable<string>();
}
