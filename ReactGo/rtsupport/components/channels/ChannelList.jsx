import React from 'react'
import Channel from './Channel.jsx'

ChannelList.propTypes = {
  channels: React.PropTypes.array.isRequired,
  setChannel: React.PropTypes.func.isRequired
}

class ChannelList extends React.Component{
  render(){
    return(
      <ul>
        {this.props.channels.map(function(channel){
          <Channel
            channel={channel}
            key={channel.id}
            setChannel={this.props.setChannel}
          />
        })}
      </ul>
    )
  }
}

export default ChannelList
