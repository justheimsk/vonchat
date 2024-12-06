import './ChannelList.scss';
import Category from './components/Category/Category';

export default function ChannelList() {
	return (
		<>
			<div id="channel-list">
				<Category />
				<Category />
			</div>
		</>
	);
}
