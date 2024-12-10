import { vonchat } from '@/lib/Application';
import type { RecvContext } from '@/lib/core/command/CommandRegistry';
import { Profile } from '@/lib/core/profile/Profile';
import { HTTPAdapter } from './adapters/HTTPAdapter';

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

		const profile = new Profile(
			name,
			email,
			password,
			new HTTPAdapter({ secure: false, host: 'localhost', port: 8080 }),
		);
		vonchat.state.reducers.profiles.appendProfile(profile);
		profile.adapter.init();
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
};
