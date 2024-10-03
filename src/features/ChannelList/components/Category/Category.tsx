import {FaAngleDown, FaPlus} from "react-icons/fa6";
import "./Category.scss";
import {Channel} from "../Channel/Channel";
import {useState} from "react";

export default function Category() {
  const [active, setActive] = useState(true);

  return (
    <>
      <div className="category">
        <div className="category__infos">
          <div onClick={() => setActive(!active)} className={`category__infos__name ${active ? 'category__infos__name--active' : ''}`}>
            <i><FaAngleDown /></i>
            <span>General</span>
          </div>
          <div className="category__infos__actions">
            <i><FaPlus /></i>
          </div>
        </div>
        {active && <div className="category__channels">
          {"h".repeat(10).split('').map(() => (
            <Channel />
          ))}
        </div>}
      </div>
    </>
  )
}
