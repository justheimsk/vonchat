import ServerList from '@features/ServerList/ServerList';
import './app.scss';
import { ErrorMagnifier } from '@/components/ErrorMagnifier/ErrorMagnifier';
import ChannelData from '@features/ChannelData/ChannelData';
import ChannelInfo from '@features/ChannelInfo/ChannelInfo';
import ChannelList from '@features/ChannelList/ChannelList';
import ServerName from '@features/ServerName/ServerName';
import UserInfo from '@features/UserInfo/UserInfo';
import UserList from '@features/UserList/UserList';

export default function App() {
	return (
		<>
			<div id="layout">
				<ErrorMagnifier />
				<ServerList />
				<div className="layout--flex-col channels-side">
					<ServerName />
					<ChannelList />
					<UserInfo />
				</div>
				<div className="layout--flex-col layout--full">
					<ChannelInfo />
					<div className="layout--flex">
						<ChannelData />
						<UserList />
					</div>
				</div>
			</div>
		</>
	);
}
