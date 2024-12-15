import { FaHeadphones, FaMicrophone } from 'react-icons/fa6';
import './UserInfo.scss';
import { vonchat } from '@/lib/Application';
import { useLibState } from '@/lib/state/Hook';
import type { ProfileStateData } from '@/shared/state/profiles';
import { FaCog } from 'react-icons/fa';

export default function UserInfo() {
	const profile = useLibState<ProfileStateData>(
		vonchat.profiles.getState(),
	).activeProfile;

	return (
		<>
			<div id="user-info">
				<div id="user-info__hoverable">
					<div id="user-info__avatar" />
					<div id="user-info__infos">
						<span>{profile?.name}</span>
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
