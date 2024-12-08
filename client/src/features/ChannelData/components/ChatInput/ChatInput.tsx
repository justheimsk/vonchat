import { CommandList } from '@/features/CommandList/CommandList';
import { vonchat } from '@/lib/Application';
import { useEffect } from 'react';
import { BsEmojiSmileFill } from 'react-icons/bs';
import { FaCirclePlus } from 'react-icons/fa6';
import { RiFileGifFill } from 'react-icons/ri';
import './ChatInput.scss';

export default function ChatInput() {
	useEffect(() => {
		const editor = document.getElementById('chat-input__editor');

		const event = vonchat.input.events.setChatInput.subscribe((text) => {
			if (editor) {
				editor.innerText = text;
				parseCommand(text.replace(/\//gi, ''));

				const range = document.createRange();
				const selection = window.getSelection();

				range.selectNodeContents(editor);
				range.collapse(false);

				if (selection) {
					selection.removeAllRanges();
					selection.addRange(range);
				}
			}
		});

		return () => {
			event.unsubscribe();
		};
	}, []);

	function handleEditorInput(e: React.FormEvent<HTMLDivElement>) {
		const target = e.target as HTMLDivElement;
		vonchat.input.resetHistoryIndex();
		vonchat.input.value = target.innerText;

		if (target.innerText.startsWith('/')) {
			vonchat.ui.openCommandList();
			parseCommand(target.innerText.replace(/\//gi, '').split(' ')[0]);
		} else {
			vonchat.ui.closeCommandList();
		}
	}

	function parseCommand(name: string) {
		const cmd = vonchat.cmdRegistry
			.getState()
			.commands.find((cmd) => cmd.name.startsWith(name));
		if (cmd) vonchat.ui.selectCommand(cmd.name);
		else vonchat.ui.selectCommand('');
	}

	function handleEnter(e: React.KeyboardEvent<HTMLDivElement>) {
		const target = e.target as HTMLDivElement;

		if (e.key === 'Enter') {
			e.preventDefault();

			vonchat.input.send(target.innerText);
		} else if (e.key === 'ArrowUp') {
			e.preventDefault();

			const lastEntry = vonchat.input.cycleHistory();
			if (lastEntry !== undefined)
				vonchat.input.events.setChatInput.notify(lastEntry);
		} else if (e.key === 'ArrowDown') {
			e.preventDefault();

			const firstEntry = vonchat.input.cycleHistory(true);
			if (firstEntry !== undefined)
				vonchat.input.events.setChatInput.notify(firstEntry);
		}
	}

	return (
		<>
			<div id="chat-input">
				<CommandList />
				<div id="chat-input__attachments">
					<i>
						<FaCirclePlus />
					</i>
				</div>
				<div
					contentEditable
					suppressContentEditableWarning
					id="chat-input__editor"
					data-name="general"
					onInput={(e) => handleEditorInput(e)}
					onKeyDown={(e) => handleEnter(e)}
				/>
				<div id="chat-input__actions">
					<i>
						<RiFileGifFill />
					</i>
					<i>
						<BsEmojiSmileFill />
					</i>
				</div>
			</div>
		</>
	);
}
