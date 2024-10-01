import {FaAngleDown, FaUserPlus} from "react-icons/fa6";
import "./Category.scss";
import {FaCog} from "react-icons/fa";
import {Channel} from "../Channel/Channel";

export default function Category() {
  return (
    <>
      <div className="category">
        <div className="category__infos">
          <div className="category__infos__name">
            <i><FaAngleDown /></i>
            <span>General</span>
          </div>
          <div className="category__infos__actions">
            <i><FaUserPlus /></i>
            <i><FaCog /></i>
          </div>
        </div>
        <div className="category__channels">
          {"h".repeat(10).split('').map(() => (
            <Channel />
          ))}
        </div>
      </div>
    </>
  )
}
