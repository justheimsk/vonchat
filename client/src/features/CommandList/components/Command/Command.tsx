import {vonchat} from "@/lib/Application";
import "./Command.scss"
import {useSelector} from "react-redux";
import type {RootState} from "@/store/store";

export interface CommandProps {
  name: string;
  description: string;
}

export function Command(props: CommandProps) {
  const active = useSelector((state: RootState) => state.uiSlice.selectedCommand);

  return (
    <>
      {/* biome-ignore lint/a11y/useKeyWithClickEvents: <explanation> */}
      <div onClick={() => vonchat.input.events.setChatInput.notify(`/${props.name}`)} className={`command ${active === props.name ? 'command--active' : ''}`}>
        <span>/{props.name}</span>
        <span className="command__desc">{props.description}</span>
      </div>
    </>
  )
}
