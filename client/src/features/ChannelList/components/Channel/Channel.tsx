import { FaHashtag, FaUserPlus } from 'react-icons/fa6';
import './Channel.scss';
import { useState } from 'react';
import { FaCog } from 'react-icons/fa';

export function Channel() {
	const [name, _] = useState('general-chat');

	return (
		<>
			<div className="channel">
				<div className="channel__infos">
					<i>
						<FaHashtag />
					</i>
					<span>{name}</span>
				</div>
				<div className="channel__actions">
					<i>
						<FaUserPlus />
					</i>
					<i>
						<FaCog />
					</i>
				</div>
			</div>
		</>
	);
}
