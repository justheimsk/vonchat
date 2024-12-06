import { vonchat } from '@/lib/Application';
import type { RecvContext } from '@/lib/core/CommandRegistry';

export default () => {
	const hello_world = () => {
		alert('Hello World!');
	};

	const reverseText = (ctx: RecvContext) => {
		const text = ctx.args.get('text');
		if (text) {
			const reversed = (text.value as string).split('').reverse();
			vonchat.input.events.setChatInput.notify(reversed.join(''));
		}
	};

	vonchat.cmdRegistry.register(
		'hello_world',
		'Simple hello world command.',
		[],
		hello_world,
	);
	vonchat.cmdRegistry.register(
		'reverse_text',
		'Reverse some text',
		[{ type: 'text', name: 'text', required: true }],
		reverseText,
	);
};
