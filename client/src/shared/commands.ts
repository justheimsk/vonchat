import { vonchat } from '@/lib/Application';
import { HTTPAdapter } from '@/lib/adapters/backend/http/HTTPAdapter';
import type { RecvContext } from '@/lib/types/Command';

export default () => {
	const hello_world = () => {
		alert('Hello World!');
	};

	vonchat.cmdRegistry.register(
		'hello_world',
		'Simple hello world command.',
		[],
		hello_world,
	);

	const reverseText = (ctx: RecvContext) => {
		const text = ctx.args.get('text');
		if (text) {
			const reversed = (text.value as string).split('').reverse();
			vonchat.input.setValue(reversed.join(''));
		}
	};

	vonchat.cmdRegistry.register(
		'reverse_text',
		'Reverse some text',
		[{ type: 'text', name: 'text', required: true }],
		reverseText,
	);

	const createProfile = async (ctx: RecvContext) => {
		const name = ctx.args.get('username')?.value as string;
		const email = ctx.args.get('email')?.value as string;
		const password = ctx.args.get('password')?.value as string;
		if (!name || !email || !password) return;

		vonchat.profiles.createProfile(name, email, password, true);
		vonchat.profiles.saveToMemory();
	};

	vonchat.cmdRegistry.register(
		'create_profile',
		'Create a new user profile',
		[
			{ type: 'text', name: 'username', required: true },
			{ type: 'text', name: 'email', required: true },
			{ type: 'text', name: 'password', required: true },
		],
		createProfile,
	);

	const connect = (ctx: RecvContext) => {
		const host = ctx.args.get('host')?.value as string;
		const port = ctx.args.get('port')?.value as string;
		if (!port || !host) return;

		const profile = vonchat.profiles.getActiveProfile();
		if (!profile) return;
		let server = profile.servers.get(host);

		if (!server) {
			server = profile.servers.createServer(
				host,
				port,
				new HTTPAdapter({ host, port, secure: false }),
				true,
				false,
			);
		}

		server.connect();
		vonchat.profiles.saveToMemory();
	};

	vonchat.cmdRegistry.register(
		'connect',
		'Connect to a server',
		[
			{ type: 'text', name: 'host', required: true },
			{ type: 'text', name: 'port', required: true },
		],
		connect,
	);
};
