import { User } from '@components/User/User';
import './UserList.scss';
import { vonchat } from '@/lib/Application';
import { useLibState } from '@/lib/state/Hook';

export default function UserList() {
	const users = useLibState(
		vonchat.context.getActiveServer()?.adapter.state.reducers.users,
	);

	return (
		<>
			<div id="user-list">
				{users?.values.map((user) => (
					<User username={user.username} status={user.status} key={user.id} />
				))}
			</div>
		</>
	);
}
