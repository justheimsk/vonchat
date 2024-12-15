import { FaHeadphones, FaMicrophone } from 'react-icons/fa6';
import './UserInfo.scss';
import { vonchat } from '@/lib/Application';
import { useLibState } from '@/lib/state/Hook';
import type { IProfilesState } from '@/shared/state/profiles';
import { FaCog } from 'react-icons/fa';

export default function UserInfo() {
	const profiles = useLibState<IProfilesState>(
		vonchat.state.reducers.profiles,
	).profiles;

	return (
		<>
			<div id="user-info">
				<div id="user-info__hoverable">
					<div id="user-info__avatar" />
					<div id="user-info__infos">
						<span>{vonchat.profiles.getActiveProfile(profiles)?.name}</span>
						<small>Online</small>
					</div>
				</div>
				<div id="user-info__actions">
					<i>
						<FaMicrophone />
					</i>
					<i>
						<FaHeadphones />
					</i>
					<i>
						<FaCog />
					</i>
				</div>
			</div>
		</>
	);
}
