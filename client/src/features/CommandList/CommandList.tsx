import {BsSlashSquareFill} from "react-icons/bs"
import "./CommandList.scss"
import {FaClock} from "react-icons/fa6"
import {useSelector} from "react-redux";
import type {RootState} from "@/store/store";

export function CommandList() {
  const active = useSelector((state: RootState) => state.uiSlice.commandList);

  return (
    <>
      <div id="command-list" className={`${active ? 'command-list--active' : ''}`}>
        <div id="command-list__sidebar">
          <FaClock />
          <BsSlashSquareFill />
        </div>
        <span id="command-list__title"><BsSlashSquareFill /> Client Commands</span>
      </div>
    </>
  )
}
