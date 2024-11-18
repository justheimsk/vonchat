import {FaAngleDown, FaPlus} from "react-icons/fa6";
import "./Category.scss";
import {Channel} from "../Channel/Channel";
import {useState} from "react";
import onEnterPress from "@/utils/onKeyPress";

export default function Category() {
  const [active, setActive] = useState(true);
  const [name, _] = useState('GENERAL');

  function handleInteraction() {
    setActive(!active);
  }

  return (
    <>
      <div className="category">
        <div className="category__infos">
          <div onKeyUp={(e) => onEnterPress(e, () => handleInteraction())} onClick={() => handleInteraction()} className={`category__infos__name ${active ? 'category__infos__name--active' : ''}`}>
            <i><FaAngleDown /></i>
            <span>{name}</span>
          </div>
          <div className="category__infos__actions">
            <i><FaPlus /></i>
          </div>
        </div>
        {active && <div className="category__channels">
          {"h".repeat(10).split('').map(() => (
            <Channel key={Math.random() * 99999} />
          ))}
        </div>}
      </div>
    </>
  )
}
