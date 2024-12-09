import { CommandList } from '@/features/CommandList/CommandList';
import { vonchat } from '@/lib/Application';
import { useEffect } from 'react';
import { BsEmojiSmileFill } from 'react-icons/bs';
import { FaCirclePlus } from 'react-icons/fa6';
import { RiFileGifFill } from 'react-icons/ri';
import './ChatInput.scss';

export default function ChatInput() {
	useEffect(() => {
		const editor = document.getElementById(
			'chat-input__editor',
		) as HTMLDivElement;

		const event = vonchat.input.events.domSetInnerText.subscribe((text) => {
			if (editor) {
				editor.innerText = text;
				vonchat.input.value = text;
				parseCommand(text);
				moveCaretToEnd(editor);
			}
		});

		return () => {
			event.unsubscribe();
		};
	}, []);

	function moveCaretToEnd(editor: HTMLDivElement) {
		const range = document.createRange();
		const selection = window.getSelection();

		range.selectNodeContents(editor);
		range.collapse(false);

		if (selection) {
			selection.removeAllRanges();
			selection.addRange(range);
		}
	}

	function parseCommand(name: string) {
		if (name.startsWith('/')) {
			vonchat.ui.openCommandList();
			const cmd = vonchat.cmdRegistry
				.getState()
				.commands.find((cmd) =>
					cmd.name.startsWith(name.trim().replace(/\//gi, '').split(' ')[0]),
				);

			if (cmd) vonchat.ui.selectCommand(cmd.name);
			else vonchat.ui.selectCommand('');
		} else {
			vonchat.ui.closeCommandList();
		}
	}

	function handleInput(e: React.FormEvent<HTMLDivElement>) {
		const target = e.target as HTMLDivElement;
		if (!target) return;

		vonchat.input.history.resetIdx();
		parseCommand(target.innerText);

		vonchat.input.setValue(target.innerText, false);
	}

	function handleKeyDown(e: React.KeyboardEvent<HTMLDivElement>) {
		const target = e.target as HTMLDivElement;
		if (!target) return;

		vonchat.input.events.onKeyDown.notify(target.innerText);
		if (e.key === 'Enter') {
			e.preventDefault();
			vonchat.input.send(vonchat.input.value);
		} else if (e.key === 'ArrowUp') {
			e.preventDefault();

			const previous = vonchat.input.history.getPrevious();
			if (previous !== undefined)
				vonchat.input.events.domSetInnerText.notify(previous);
		} else if (e.key === 'ArrowDown') {
			e.preventDefault();
			vonchat.input.events.domSetInnerText.notify(
				vonchat.input.history.getNext(),
			);
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
					onInput={(e) => handleInput(e)}
					onKeyDown={(e) => handleKeyDown(e)}
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
