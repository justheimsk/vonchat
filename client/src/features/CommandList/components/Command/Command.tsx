import {vonchat} from "@/lib/Application";
import "./Command.scss"
import {useSelector} from "react-redux";
import type {RootState} from "@/store/store";
import type {Arg} from "@/lib/core/Command";

export interface CommandProps {
  name: string;
  description: string;
  args: Arg[]
}

export function Command(props: CommandProps) {
  const active = useSelector((state: RootState) => state.ui.selectedCommand);

  return (
    <>
      {/* biome-ignore lint/a11y/useKeyWithClickEvents: <explanation> */}
      <div onClick={() => vonchat.input.events.setChatInput.notify(`/${props.name}`)} className={`command ${active === props.name ? 'command--active' : ''}`}>
        <div className="command__header">
          <span>/{props.name}</span>
          <div className="command__args">
            {props.args.map((arg) => (
              <span key={arg.name} className="command__arg">{arg.name}</span>
            ))}
          </div>
        </div>
        <span className="command__desc">{props.description}</span>
      </div>
    </>
  )
}
