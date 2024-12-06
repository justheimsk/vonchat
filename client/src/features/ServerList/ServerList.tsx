import ServerButton from './components/ServerButton/ServerButton';
import './ServerList.scss';

export default function ServerList() {
	return (
		<>
			<div id="server-list">
				<div className="server-list__panel server-list__panel--small">
					<ServerButton />
					<ServerButton />
				</div>
				<div className="server-list__panel">
					{'h'
						.repeat(20)
						.split('')
						.map(() => (
							<ServerButton key={Math.random() * 99999} />
						))}
				</div>
			</div>
		</>
	);
}
