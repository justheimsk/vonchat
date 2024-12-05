import {FaCirclePlus} from "react-icons/fa6";
import "./ChatInput.scss";
import {BsEmojiSmileFill} from "react-icons/bs";
import {RiFileGifFill} from "react-icons/ri";
import {CommandList} from "@/features/CommandList/CommandList";
import {useEffect} from "react";
import {vonchat} from "@/lib/Application";
import type {Subscription} from "@/lib/core/Observable";

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
        parseCommand(text.replace(/\//gi, ""))
      }
    }));

    events.push(vonchat.input.events.appendChatInput.subscribe((text) => {
      if(editor) {
        editor.innerText += text;
      }
    }))

    return () => {
      events.map((ev) => ev.unsubscribe());
    }
  }, []);

  function handleEditorInput(e: React.FormEvent<HTMLDivElement>) {
    const target = e.target as HTMLDivElement;

    if(target.innerText.startsWith("/")) {
      vonchat.ui.openCommandList();
      parseCommand(target.innerText.replace(/\//gi, ""));
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
        const selectedCommand = vonchat.ui.getState().selectedCommand;

        if(selectedCommand) {
          e.preventDefault();
          vonchat.cmdRegistry.exec(selectedCommand);

          vonchat.ui.closeCommandList();
          vonchat.input.events.clearChatInput.notify(null);
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
