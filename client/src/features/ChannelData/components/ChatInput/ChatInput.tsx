import {FaCirclePlus} from "react-icons/fa6";
import "./ChatInput.scss";
import {BsEmojiSmileFill} from "react-icons/bs";
import {RiFileGifFill} from "react-icons/ri";
import {vonchat} from "@/lib/Application";
import {useSelector} from "react-redux";
import type {RootState} from "@/store/store";
import {CommandList} from "@/features/CommandList/CommandList";

export default function ChatInput() {
  const value = useSelector((state: RootState) => state.input.chatInput);

  function onChange(e: React.FormEvent<HTMLDivElement>) {
    const target = e.target as HTMLDivElement;
    if(target.innerText.startsWith("/")) {
      vonchat.ui.openCommandList();
    } else {
      vonchat.ui.closeCommandList();
    }
    
    vonchat.input.setChatInputValue(target.innerText);
  }

  function onEnter(e: React.KeyboardEvent<HTMLDivElement>) {
    const target = e.target as HTMLInputElement;
    if(e.key === "Enter") {
      if(target.innerText.startsWith("/")) {
        vonchat.cmdRegistry.execCommand(target.innerText.replace(/\//gi, ""))
        vonchat.ui.closeCommandList()
        vonchat.input.setChatInputValue("")
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
          onInput={(e) => onChange(e)}
          onKeyUp={(e) => onEnter(e)}
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
