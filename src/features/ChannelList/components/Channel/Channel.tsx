import {FaHashtag, FaUserPlus} from "react-icons/fa6";
import "./Channel.scss";
import {FaCog} from "react-icons/fa";

export function Channel() {
  return (
    <>
      <div className="channel">
        <div className="channel__infos">
          <i><FaHashtag /></i>
          <span>General chat</span>
        </div>
        <div className="channel__actions">
          <i><FaUserPlus /></i>
          <i><FaCog /></i>
        </div>
      </div>
    </>
  )
}
