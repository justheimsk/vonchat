import { useLibState } from '@/lib/state/Hook';
import ServerButton from './components/ServerButton/ServerButton';
import './ServerList.scss';
import { vonchat } from '@/lib/Application';
import type { ProfileStateData } from '@/shared/state/profiles';

export default function ServerList() {
	const profile = useLibState<ProfileStateData>(
		vonchat.profiles.getState(),
	).activeProfile;

	return (
		<>
			<div id="server-list">
				<div className="server-list__panel server-list__panel--small">
					<ServerButton />
					<ServerButton />
				</div>
				<div className="server-list__panel">
					{profile?.servers.getAll().map((s) => (
						<ServerButton key={s.host} />
					))}
				</div>
			</div>
		</>
	);
}
