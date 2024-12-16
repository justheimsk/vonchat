import ServerList from '@features/ServerList/ServerList';
import './app.scss';
import { ErrorMagnifier } from '@/components/ErrorMagnifier/ErrorMagnifier';
import { vonchat } from '@/lib/Application';
import { ModalBuilder } from '@/lib/core/ui/ModalBuilder';
import ChannelData from '@features/ChannelData/ChannelData';
import ChannelInfo from '@features/ChannelInfo/ChannelInfo';
import ChannelList from '@features/ChannelList/ChannelList';
import ServerName from '@features/ServerName/ServerName';
import UserInfo from '@features/UserInfo/UserInfo';
import UserList from '@features/UserList/UserList';
import { useEffect } from 'react';

export default function App() {
	useEffect(() => {
		const read = localStorage.getItem('noticeRead');

		if (!read) {
			const modal = new ModalBuilder()
				.setTitle('Important Notice')
				.setDescription(`This project is still in its early stages of development, so you may encounter bugs, issues, or features that are incomplete or under construction. We are actively working to improve the experience and make the project more robust.
          \nIf you find any problems or have suggestions, we’d greatly appreciate it if you could report them on our GitHub repository. Your feedback and collaboration are essential for the growth and success of this open source project!
          \nAlso, if you enjoy coding and would like to contribute, feel free to explore the codebase and submit pull requests. Every contribution helps!
          \nLastly, if you like the project, consider leaving a star on our GitHub repository. It’s a small gesture, but it goes a long way in supporting the project and motivating us to keep improving.
          \nThank you for being part of this journey with us.`)
				.addButton('default', 'Github', () => {
					window.open('https://github.com/justheimsk/vonchat', '_blank');
					return false;
				})
				.addButton('success', 'Ok', () => {})
				.setOnCloseCallback(() => {
					localStorage.setItem('noticeRead', 'true');
					return true;
				});
			vonchat.ui.createModal(modal);
		}
	}, []);

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
