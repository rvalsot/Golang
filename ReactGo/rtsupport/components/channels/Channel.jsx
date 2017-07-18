import React from 'react'

// Accepted properties for this component
Channel.propTypes = {
  // name: React.Proptypes.(the type).(is required?)
  channel: React.PropTypes.object.isRequired,
  setChannel: React.PropTypes.func.isRequired
}


// Channel class renders ...

class Channel extends React.Component{
  // onClick(event) method ...
  onClick(event){
    //
    event.preventDefault()
    // destructuring of the props, to sorthen the names
    const {setChannel, channel} = this.props

    // Action of setting the channel (1)
    setChannel(channel)
  }

  render(){
    // destructuring of the props, to shorten the name
    const {channel} = this.props
    return(
      <li>
        <a onClick={this.onClick.bind(this)}>
          {channel.name}
        </a>
      </li>
    )
  }
}

export default Channel


/* 1.- This action will pass a channel name to render the clicked channel to
 *
 */
