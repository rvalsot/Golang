import React from 'react'

import ChannelForm from './ChannelForm.jsx'
import ChannelList from './ChannelList.jsx'

ChannelSection.propTypes = {
  channels: React.PropTypes.array.isRequired,
  setChannel: React.PropTypes.func.isRequired,
  addChannel: React.PropTypes.func.isRequired
}

class ChannelSection extends React.Component{
  render (){
    return (
      <div>
        <ChannelList
          {...this.props}
        />
        <ChannelForm
          {...this.props}
        />
      </div>
    )
  }
}

export default ChannelSections
