import {FaCirclePlus} from "react-icons/fa6";
import "./ChatInput.scss";
import {BsEmojiSmileFill} from "react-icons/bs";
import {RiFileGifFill} from "react-icons/ri";
import {useDispatch} from "react-redux";
import {toggleCommandList} from "@/store/slices/ui";

export default function ChatInput() {
  const dispatch = useDispatch()

  function onChange(e: React.FormEvent<HTMLDivElement>) {
    const target = e.target as HTMLDivElement;
    if(target.innerText.startsWith("/")) {
      dispatch(toggleCommandList(true))
    } else {
      dispatch(toggleCommandList(false))
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
