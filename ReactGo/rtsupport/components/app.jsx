import React from 'react'

import ChannelSection from './channels/ChannelSection.jsx'

class App extends React.Component {
  constructor(props){
    super(props)
    this.state = {
      channels : []
    }
  }

  addChannel(name){
    let {channels} = this.state
    // New channel
    channels.push({
      id: channels.length, name
    })
    // New modified channels array
    this.setState(
      {channels}
    )

    // Todo: sent t oserver

  }

  setChannel(activeChannel){
    this.setState({activeChannel})
    // Todo: get the channel's messages
  }

  render(){
    return(
      <div className='app'>
        <div className='nav'>
          <ChannelSection
            channels={this.state.channels}
            addChannel={this.addChannel.bind(this)}
            setChannel={this.setChannel.bind(this)}
          />
        </div>
      </div>
    )
  }
}

export default App
