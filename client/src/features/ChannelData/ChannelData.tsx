import './ChannelData.scss';
import Message from '@components/Message/Message';
import ChatInput from './components/ChatInput/ChatInput';

export default function ChannelData() {
	return (
		<>
			<div id="channel-data">
				<div id="channel-data__messages">
					{'h'
						.repeat(20)
						.split('')
						.map(() => (
							<Message key={Math.random() * 99999} />
						))}
				</div>
				<ChatInput />
			</div>
		</>
	);
}
