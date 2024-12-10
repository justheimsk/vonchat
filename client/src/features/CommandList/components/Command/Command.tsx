import { vonchat } from '@/lib/Application';
import './Command.scss';
import type CommandLib from '@/lib/core/command/Command';
import { useLibState } from '@/lib/state/Hook';

export interface CommandProps {
	self: CommandLib;
}

export function Command(props: CommandProps) {
	const active = useLibState(vonchat.state.reducers.ui).selectedCommand;

	return (
		<>
			{/* biome-ignore lint/a11y/useKeyWithClickEvents: <explanation> */}
			<div
				onClick={() => vonchat.input.formatCommandInChatInput(props.self)}
				className={`command ${active === props.self.name ? 'command--active' : ''}`}
			>
				<div className="command__header">
					<span>/{props.self.name}</span>
					<div className="command__args">
						{props.self.args.map((arg) => (
							<span key={arg.name} className="command__arg">
								{arg.name}
							</span>
						))}
					</div>
				</div>
				<span className="command__desc">{props.self.description}</span>
			</div>
		</>
	);
}
