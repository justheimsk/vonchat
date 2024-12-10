import { State } from '@/lib/state/State';

export interface IUiState {
	commandListOpen: boolean;
	selectedCommand: string;
}

export class UiState extends State<IUiState> {
	private commandListOpen = false;
	private selectedCommand = '';

	public get data() {
		return {
			commandListOpen: this.commandListOpen,
			selectedCommand: this.selectedCommand,
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
}
