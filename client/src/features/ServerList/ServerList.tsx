import { useLibState } from '@/lib/state/Hook';
import ServerButton from './components/ServerButton/ServerButton';
import './ServerList.scss';
import { vonchat } from '@/lib/Application';

export default function ServerList() {
	const profiles = useLibState(vonchat.state.reducers.profiles).profiles;

	return (
		<>
			<div id="server-list">
				<div className="server-list__panel server-list__panel--small">
					<ServerButton />
					<ServerButton />
				</div>
				<div className="server-list__panel">
					{vonchat.profiles.getActiveProfile(profiles)?.servers.map((s) => (
						<ServerButton key={s.ip} />
					))}
				</div>
			</div>
		</>
	);
}
