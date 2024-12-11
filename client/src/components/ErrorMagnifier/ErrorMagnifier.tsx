import { vonchat } from '@/lib/Application';
import './ErrorMagnifier.scss';
import { useEffect, useState } from 'react';

export function ErrorMagnifier() {
	const [errors, setErrorCount] = useState(0);

	useEffect(() => {
		vonchat.logs.subscribe(() => {
			const errors = vonchat.logs.logs.filter((log) => log.level === 'error');
			setErrorCount(errors.length);
		});
	}, []);

	return (
		<>
			<div
				className={`error-magnifier ${errors ? 'error-magnifier--active' : ''}`}
			>
				<span>{errors} errors</span>
			</div>
		</>
	);
}
