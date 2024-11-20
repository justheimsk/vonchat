import ServerList from '@features/ServerList/ServerList'
import './app.scss'
import ChannelList from '@features/ChannelList/ChannelList'
import ServerName from '@features/ServerName/ServerName'
import UserInfo from '@features/UserInfo/UserInfo'
import ChannelInfo from '@features/ChannelInfo/ChannelInfo'
import ChannelData from '@features/ChannelData/ChannelData'
import UserList from '@features/UserList/UserList'

export default function App() {
    return (
    <>
      <div id="layout">
        <ServerList />
        <div className="layout--flex-col channels-side">
          <ServerName />
          <ChannelList />
          <UserInfo />
        </div>
        <div className="layout--flex-col layout--full">
          <ChannelInfo />
          <div className="layout--flex">
            <ChannelData />
            <UserList />
          </div>
        </div>
      </div>
    </>
  )
}
