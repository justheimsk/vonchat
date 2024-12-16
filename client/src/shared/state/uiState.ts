import { State } from '@/lib/state/State';
import type { Modal } from '@/lib/types/Modal';

export interface IUiState {
	commandListOpen: boolean;
	selectedCommand: string;
	modal?: Modal;
}

export class UiState extends State<IUiState> {
	private commandListOpen = false;
	private selectedCommand = '';
	private modal?: Modal;

	public get data() {
		return {
			commandListOpen: this.commandListOpen,
			selectedCommand: this.selectedCommand,
			modal: this.modal,
		};
	}

	public toggleCommandList(open: boolean) {
		this.commandListOpen = open;
		return this;
	}

	public selectCommand(name: string) {
		this.selectedCommand = name;
		return this;
	}

	public setModal(modal: Modal) {
		this.modal = modal;
		return this;
	}

	public removeModal() {
		this.modal = undefined;
		return this;
	}
}
