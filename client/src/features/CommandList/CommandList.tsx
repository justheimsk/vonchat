import {BsSlashSquareFill} from "react-icons/bs"
import "./CommandList.scss"
import {FaClock} from "react-icons/fa6"
import {useSelector} from "react-redux";
import type {RootState} from "@/store/store";
import {Command} from "./components/Command/Command"

export function CommandList() {
  const active = useSelector((state: RootState) => state.ui.commandList);
  const commands = useSelector((state: RootState) => state.commandRegistry.commands);

  return (
    <>
      <div id="command-list" className={`${active ? 'command-list--active' : ''}`}>
        <div id="command-list__sidebar">
          <FaClock />
          <BsSlashSquareFill />
        </div>
        <div id="command-list__context">
          <span id="command-list__title"><BsSlashSquareFill /> Client Commands</span>
          <div id="command-list__commands">
            {Array.from(commands.values()).map((cmd) => (
              <Command key={cmd.name} self={cmd} />
            ))}
          </div>
        </div>
      </div>
    </>
  )
}
