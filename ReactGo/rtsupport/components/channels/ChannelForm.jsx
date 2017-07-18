import React from 'react'

ChannelForm.propTypes = {
  addChannel: React.PropTypes.func.isRequired
}

class ChannelForm extends React.Component{

  // onSubmit method (1)
  onSubmit(event){
    event.preventDefault()

    const node = this.refs.channel
    const channelName =node.value

    this.props.addChannel(channelName)

    // to empty string clear
    node.value = ''
  }

  render() {
    return (
      <form onSubmit={this.onSubmit.bind(this)}>
        <input type='text'
          ref='channel'
        />
      </form>
    )
  }
}

export default ChannelForm

/*
 * 1.- onSubmit will create a new channel with the input of the form
 *
 */
