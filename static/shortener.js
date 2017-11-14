
class Results extends React.Component {
  constructor(props) {
    super(props);
    
    this.state = { value: '' };
  }

getPath(){
  alert(window.location.href);  
}

  render() {
    return (

      <div>
        <a href={this.props.data.ShortUrl}>Short link!</a>
        <table class="table table-hover">
          <tbody>
            <tr>
              <td>Long URL</td>
              <td>{this.props.data.LongUrl}</td>
            </tr>
            <tr>
            <td>Short URL</td>
            <td>{this.props.data.ShortUrl}</td>
            </tr>
            <tr>
              
              <td>Info url</td>
              <td>/info/{this.props.data.ShortUrl}</td>
            </tr>
          </tbody>
        </table>
      </div>
    );
  }
}

class NameForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = { value: '', showResults: false, data: "", previous: [] };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({ value: event.target.value });
  }

  handleSubmit(event) {
    this.setState({ showResults: true });
    
    fetch('/', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
      },
      body: this.state.value,
    }).then(res => res.json())
      .then(data => this.setState({ data, previous: [data].concat(this.state.previous) }));

    //this.setState({data: JSON.parse(xmlhttp.responseText)});
    event.preventDefault();
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          <input className="input" type="url" required placeholder="http://www.example.com" value={this.state.value} onChange={this.handleChange} />
        </label>
        <input type="submit" value="Submit" />
        
        {this.state.showResults ? <Results data={this.state.data} /> : null}

        {
          this.state.previous.map(prev => (
            <div style={{ display: 'flex' }}>
              <div style={{ margin: '5px' }}>{prev.ShortUrl}</div>
              <div style={{ margin: '5px' }}>{prev.LongUrl}</div>
            </div>
          ))
        }
      </form>

    );
  }
}


ReactDOM.render(
  <NameForm />,
  document.getElementById('root')
);