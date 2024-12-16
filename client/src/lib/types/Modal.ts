export type ModalButtonType = 'danger' | 'success' | 'default';

export interface ModalButton {
	type: ModalButtonType;
	label: string;
	// biome-ignore lint/suspicious/noConfusingVoidType: <explanation>
	callback?: () => boolean | void;
}

export interface Modal {
	title: string;
	description: string;
	buttons: ModalButton[];
	// biome-ignore lint/suspicious/noConfusingVoidType: <explanation>
	onClose: () => boolean | void;
}
