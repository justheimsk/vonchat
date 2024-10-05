import {FaCirclePlus} from "react-icons/fa6";
import "./ChatInput.scss";
import {BsEmojiSmileFill} from "react-icons/bs";
import {RiFileGifFill} from "react-icons/ri";

export default function ChatInput() {
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
