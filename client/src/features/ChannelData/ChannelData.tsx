import {CommandList} from "../CommandList/CommandList";
import "./ChannelData.scss";
import ChatInput from "./components/ChatInput/ChatInput";
import Message from "@components/Message/Message";

export default function ChannelData() {
  return (
    <>
      <div id="channel-data">
        <div id="channel-data__messages">
          {"h".repeat(20).split('').map(() => (
            <Message key={Math.random() * 99999} />
          ))}
        </div>
        <CommandList />
        <ChatInput />
      </div>
    </>
  )
}
