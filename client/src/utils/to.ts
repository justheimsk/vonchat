export async function to<T>(
	prm: Promise<T>,
): Promise<[T | null, unknown | null]> {
	return prm
		.then((res) => [res, null] as [T, null])
		.catch((err) => [null, err] as [null, unknown]);
}
