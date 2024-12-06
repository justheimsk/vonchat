import {vonchat} from "@/lib/Application";
import "./Command.scss"
import {useSelector} from "react-redux";
import type {RootState} from "@/store/store";
import type CommandLib from "@/lib/core/Command";

export interface CommandProps {
  self: CommandLib
}

export function Command(props: CommandProps) {
  const active = useSelector((state: RootState) => state.ui.selectedCommand);

  return (
    <>
      {/* biome-ignore lint/a11y/useKeyWithClickEvents: <explanation> */}
      <div onClick={() => vonchat.input.formatCommandInChatInput(props.self)} className={`command ${active === props.self.name ? 'command--active' : ''}`}>
        <div className="command__header">
          <span>/{props.self.name}</span>
          <div className="command__args">
            {props.self.args.map((arg) => (
              <span key={arg.name} className="command__arg">{arg.name}</span>
            ))}
          </div>
        </div>
        <span className="command__desc">{props.self.description}</span>
      </div>
    </>
  )
}
