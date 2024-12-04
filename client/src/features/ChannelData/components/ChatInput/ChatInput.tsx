import {FaCirclePlus} from "react-icons/fa6";
import "./ChatInput.scss";
import {BsEmojiSmileFill} from "react-icons/bs";
import {RiFileGifFill} from "react-icons/ri";
import {vonchat} from "@/lib/Application";

export default function ChatInput() {
  function onChange(e: React.FormEvent<HTMLDivElement>) {
    const target = e.target as HTMLDivElement;
    if(target.innerText.startsWith("/")) {
      vonchat.ui.openCommandList();
    } else {
      vonchat.ui.closeCommandList()
    }
  }

  return (
    <>
      <div id="chat-input">
        <div id="chat-input__attachments">
          <i><FaCirclePlus /></i>
        </div>
        <div
          suppressContentEditableWarning
          id="chat-input__editor"
          contentEditable
          data-name="general"
          onInput={(e) => onChange(e)}
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
