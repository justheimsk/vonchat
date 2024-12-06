import {FaCirclePlus} from "react-icons/fa6";
import "./ChatInput.scss";
import {BsEmojiSmileFill} from "react-icons/bs";
import {RiFileGifFill} from "react-icons/ri";
import {CommandList} from "@/features/CommandList/CommandList";
import {useEffect} from "react";
import {vonchat} from "@/lib/Application";
import type {Subscription} from "@/lib/core/Observable";
import type {RecvArg} from "@/lib/core/CommandRegistry";

export default function ChatInput() {
  useEffect(() => {
    const editor = document.getElementById('chat-input__editor');
    
    // biome-ignore lint/suspicious/noExplicitAny: <explanation>
    const events: Subscription<any>[] = []

    events.push(vonchat.input.events.clearChatInput.subscribe(() => {
      if(editor) {
        editor.innerText = "";
      }
    }));

    events.push(vonchat.input.events.setChatInput.subscribe((text) => {
      if(editor) {
        editor.innerText = text;
        parseCommand(text.replace(/\//gi, ""));

        const range = document.createRange();
        const selection = window.getSelection();

        range.selectNodeContents(editor);
        range.collapse(false);

        if(selection) {
          selection.removeAllRanges();
          selection.addRange(range);
        }
      }
    }));

    return () => {
      events.map((ev) => ev.unsubscribe());
    }
  }, []);

  function handleEditorInput(e: React.FormEvent<HTMLDivElement>) {
    const target = e.target as HTMLDivElement;

    if(target.innerText.startsWith("/")) {
      vonchat.ui.openCommandList();
      parseCommand(target.innerText.replace(/\//gi, "").split(" ")[0]);
    } else {
      vonchat.ui.closeCommandList()
    }
  }

  function parseCommand(name: string) {
    const cmd = vonchat.cmdRegistry.getState().commands.find((cmd) => cmd.name.startsWith(name));
    if(cmd) vonchat.ui.selectCommand(cmd.name);
    else vonchat.ui.selectCommand("");
  }

  function handleEnter(e: React.KeyboardEvent<HTMLDivElement>) {
    const target = e.target as HTMLDivElement;
    if(e.key === "Enter") {
      if(target.innerText.startsWith("/")) {
        const entry = target.innerText;

        const selectedCommand = vonchat.ui.getState().selectedCommand;
        const cmd = vonchat.cmdRegistry.fetch(selectedCommand);
        const args: RecvArg[] = [];

        const regex = /([a-zA-Z0-9_-]+)="([^"]*)"/g;
        let match: RegExpExecArray | null;
        const test = entry.split(" ").slice(1).join(" ")

        // biome-ignore lint/suspicious/noAssignInExpressions: <explanation>
        while((match = regex.exec(test)) !== null) {
          const key = match[1];
          const value = match[2].trim();

          if(!value) return;
          args.push({ name: key, value });
        }

        if(cmd) {
          e.preventDefault();

          const formatCommandInChatInput = () => {
            return vonchat.input.events.setChatInput.notify(`/${cmd.name} ${cmd.args.map((arg) => `${arg.name}=""`).join(" ")}`);
          }

          const requiredArgs = cmd.args.filter((arg) => arg.required === true)
          if(requiredArgs.length > args.length) {
            for(const _arg of requiredArgs) {
              const arg = args.find((arg) => arg.name === _arg.name);
              if(!arg || arg.value) return formatCommandInChatInput();
            }

            return formatCommandInChatInput();
          }

          vonchat.input.events.clearChatInput.notify(null);
          vonchat.cmdRegistry.exec(selectedCommand, args);
          vonchat.ui.closeCommandList();
        }
      }
    }
  }

  return (
    <>
      <div id="chat-input">
        <CommandList />
        <div id="chat-input__attachments">
          <i><FaCirclePlus /></i>
        </div>
        <div
          contentEditable
          suppressContentEditableWarning
          id="chat-input__editor"
          data-name="general"
          onInput={(e) => handleEditorInput(e)}
          onKeyDown={(e) => handleEnter(e)}
        >
        </div>
        <div id="chat-input__actions">
          <i><RiFileGifFill /></i>
          <i><BsEmojiSmileFill /></i>
        </div>
      </div>
    </>
  )
}
