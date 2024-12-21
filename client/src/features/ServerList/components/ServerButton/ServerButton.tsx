import { vonchat } from '@/lib/Application';
import './ServerButton.scss';
import type { Server } from '@/lib/core/server/Server';

export interface ServerButtonProps {
	self?: Server;
	disableDot?: boolean;
}

export default function ServerButton(props: ServerButtonProps) {
	return (
		<>
			<div
				onClick={() => {
					props.self?.host &&
						vonchat.profiles
							.getActiveProfile()
							?.servers.setActiveServer(props.self);
					props.self?.connect();
				}}
				className={`server-button ${props.self?.active ? 'server-button--active' : ''} ${props.disableDot ? 'server-button--nodot' : ''}`}
			/>
		</>
	);
}
