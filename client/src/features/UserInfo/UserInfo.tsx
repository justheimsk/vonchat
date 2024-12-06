import { FaHeadphones, FaMicrophone } from 'react-icons/fa6';
import './UserInfo.scss';
import { FaCog } from 'react-icons/fa';

export default function UserInfo() {
	return (
		<>
			<div id="user-info">
				<div id="user-info__hoverable">
					<div id="user-info__avatar" />
					<div id="user-info__infos">
						<span>vonderheimsk</span>
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
