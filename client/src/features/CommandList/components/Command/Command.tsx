import {vonchat} from "@/lib/Application";
import "./Command.scss"

export interface CommandProps {
  name: string;
  description: string;
}

export function Command(props: CommandProps) {
  return (
    <>
      {/* biome-ignore lint/a11y/useKeyWithClickEvents: <explanation> */}
      <div onClick={() => vonchat.input.setChatInputValue(`/${props.name}`)} className="command" key={props.name}>
        <span>/{props.name}</span>
        <span className="command__desc">{props.description}</span>
      </div>
    </>
  )
}
