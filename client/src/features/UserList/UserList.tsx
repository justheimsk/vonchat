import {User} from "@components/User/User";
import "./UserList.scss";

export default function UserList() {
  return (
    <>
      <div id="user-list">
        {"h".repeat(20).split('').map(() => (
          <User />
        ))}
      </div>
    </>
  )
}
