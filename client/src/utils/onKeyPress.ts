export default function onEnterPress(
	e: React.KeyboardEvent<HTMLDivElement>,
	cb: () => void,
) {
	if (e.key === 'Enter') {
		cb();
	}
}
