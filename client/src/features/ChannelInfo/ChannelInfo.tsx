import { FaBell, FaHashtag, FaInbox, FaUsers } from 'react-icons/fa6';
import './ChannelInfo.scss';
import { TiPin } from 'react-icons/ti';
import Input from './components/Input/Input';

export default function ChannelInfo() {
	return (
		<>
			<div id="channel-info">
				<div id="channel-info__infos">
					<i>
						<FaHashtag />
					</i>
					<span>General</span>
				</div>
				<div id="channel-info__actions">
					<i>
						<FaBell />
					</i>
					<i>
						<TiPin />
					</i>
					<i>
						<FaUsers />
					</i>
					<i>
						<FaInbox />
					</i>
					<Input placeholder="Search" />
				</div>
			</div>
		</>
	);
}
