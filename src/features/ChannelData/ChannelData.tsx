import "./ChannelData.scss";
import ChatInput from "./components/ChatInput/ChatInput";
import Message from "./components/Message/Message";

export default function ChannelData() {
  return (
    <>
      <div id="channel-data">
        <div id="channel-data__messages">
          {"h".repeat(20).split('').map(() => (
            <Message />
          ))}
        </div>
        <ChatInput />
      </div>
    </>
  )
}
