import type { Modal, ModalButton, ModalButtonType } from '@/lib/types/Modal';

export interface ModalConstructor {
	title: string;
	description: string;
	buttons: ModalButton[];
}

export class ModalBuilder {
	public title = '';
	public description = '';
	public buttons: ModalButton[] = [];
	public onCloseCb: Modal['onClose'] = () => true;

	public constructor(modal?: ModalConstructor) {
		if (modal) {
			this.title = modal.title;
			this.description = modal.description;
			this.buttons = modal.buttons;
		}
	}

	public setTitle(title: string) {
		this.title = title;
		return this;
	}

	public setDescription(description: string) {
		this.description = description;
		return this;
	}

	addButton(button: ModalButton): this;
	addButton(
		type: ModalButtonType,
		label: string,
		cb: ModalButton['callback'],
	): this;

	public addButton(
		button: ModalButton | ModalButtonType,
		label?: string,
		cb?: ModalButton['callback'],
	) {
		if (typeof button === 'string') {
			this.buttons.push({
				type: button,
				label: label || 'Ok',
				callback: cb,
			});
		} else if ('type' in button) {
			this.buttons.push(button);
		}

		return this;
	}

	public setOnCloseCallback(cb: () => boolean) {
		this.onCloseCb = cb;
		return this;
	}

	public toJSON(): Modal {
		return {
			title: this.title,
			description: this.description,
			buttons: this.buttons.map((bn) => bn),
			onClose: this.onCloseCb,
		};
	}
}
