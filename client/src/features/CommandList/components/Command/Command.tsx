import "./Command.scss"

export interface CommandProps {
  name: string;
  description: string;
}

export function Command(props: CommandProps) {
  return (
    <>
      <div className="command" key={props.name}>
        <span>/{props.name}</span>
        <span className="command__desc">{props.description}</span>
      </div>
    </>
  )
}
