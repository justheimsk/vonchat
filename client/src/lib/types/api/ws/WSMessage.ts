export interface WSMessage {
	op: number;
	d?: unknown | undefined | null;
	t?: string | undefined | null;
}
