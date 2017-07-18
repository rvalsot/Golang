// Auxiliaries ________________________________________________________________


// let channelComponents = channels.map(function(channel){
//   return <Channel name={channel.name} />
// })

// Class Channel component ____________________________________________________

class Channel extends React.Component {
  onClick(){
    console.log("I was clicked!", this.props.name)
  }
  render(){
    return(
      <li onClick={this.onClick.bind(this)}> {this.props.name} </li>
    )
  }
}

// Class ChannelList component ________________________________________________

class ChannelList extends React.Component{
  render(){
    return(
      <ul>
        {this.props.channels.map(function(channel){
          return (
            <Channel name={channel.name}/>
          )
        })}
      </ul>
    )
  }
}

// Class ChannelForm component ________________________________________________

class ChannelForm extends React.Component{
  constructor(props){
    super(props)
    this.state = {}
  }

  onChange(event){
    this.setState({
      channelName: event.target.value
    })
  }

  onSubmit(event){
    let {channelName} = this.state
    console.log(channelName)
    // Clear the input
    this.setState({
      channelName: ''
    })

    // Adds a channel, see definition at P/ChannelSection
    this.props.addChannel(channelName)

    // Avoids trying to submit via HTTP
    event.preventDefault()
  }

  render(){
    return(
      <form onSubmit={this.onSubmit.bind(this)}>
        <input type='text'
          onChange={this.onChange.bind(this)}

          value={this.state.channelName}
        />
      <input type='submit' value='Submit'/>
      </form>
    )
  }
}

// Class ChannelSection component _____________________________________________

class ChannelSection extends React.Component{
  constructor(props){
    super(props)
    this.state = {
      channels : [
        {name: 'Numenor'},
        {name: 'Gondor'},
        {name: 'Arnor'},
        {name: 'Rohan'},
        {name: 'Rovanion'},
        {name: 'Rivendell'}
      ]
    }
  }

  // addChannel is the method that onSubmit of ChannelForm, creates a new list of channels
  addChannel(name){
    let {channels} = this.state
    channels.push({
      name: name
    })
    this.setState({
      channels: channels
    })
  }

  render(){
    return (
      <div>
        <ChannelList channels={this.state.channels} />
        <ChannelForm addChannel={this.addChannel.bind(this)} />
      </div>
    )
  }
}


// React DOM __________________________________________________________________

ReactDOM.render(<ChannelSection />, document.getElementById('app'))
