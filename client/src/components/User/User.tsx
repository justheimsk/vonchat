import type { UserStatus } from '@/lib/types/api/common/APIUser';
import './User.scss';

export interface UserProps {
	username: string;
	status: UserStatus;
}

export function User(props: UserProps) {
	return (
		<>
			<div className="user">
				<div className={`user__avatar user__avatar--${props.status}`} />
				<div className="user__infos">
					<span>{props.username}</span>
					<small>{props.status}</small>
				</div>
			</div>
		</>
	);
}
