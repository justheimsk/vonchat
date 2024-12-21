import { BsSlashSquareFill } from 'react-icons/bs';
import './CommandList.scss';
import { vonchat } from '@/lib/Application';
import { useLibState } from '@/lib/state/Hook';
import type { ICommandRegistryState } from '@/shared/state/commandRegistry';
import type { IUiState } from '@/shared/state/uiState';
import { useEffect } from 'react';
import { FaClock } from 'react-icons/fa6';
import { Command } from './components/Command/Command';

export function CommandList() {
	const active = useLibState<IUiState>(
		vonchat.state.reducers.ui,
	)?.commandListOpen;
	const registry = useLibState<ICommandRegistryState>(
		vonchat.state.reducers.cmdRegistry,
	);

	useEffect(() => {
		vonchat.state.reducers.ui.subscribe((state) => {
			const element = document.getElementById(state.selectedCommand);
			if (element) element.scrollIntoView({ behavior: 'smooth', block: 'end' });
		});
	}, []);

	return (
		<>
			<div
				id="command-list"
				className={`${active ? 'command-list--active' : ''}`}
			>
				<div id="command-list__sidebar">
					<FaClock />
					<BsSlashSquareFill />
				</div>
				<div id="command-list__context">
					<span id="command-list__title">
						<BsSlashSquareFill /> Client Commands
					</span>
					<div id="command-list__commands">
						{registry &&
							Array.from(registry.commands.values()).map((cmd) => (
								<Command key={cmd.name} self={cmd} />
							))}
					</div>
				</div>
			</div>
		</>
	);
}
