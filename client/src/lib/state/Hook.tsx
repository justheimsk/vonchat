import { useEffect, useState } from 'react';
import type { State } from './State';

export function useLibState<T>(state: State<T>) {
	const [data, setData] = useState<T>(state.data);

	useEffect(() => {
		const sub = state.subscribe((data) => {
			setData(data);
		});

		return () => {
			sub.unsubscribe();
		};
	}, [state]);

	return data;
}
