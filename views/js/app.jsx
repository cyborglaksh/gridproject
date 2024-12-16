import React from 'react';

// Testing with "this is testing"
class App extends React.Component {
    render() {
        return (<div className="test">this is testing</div>);
    }
}

class Grid extends React.Component {
  render() {
    const rows = 20;
    const cols = 20;
    const grid = [];

    for (let i = 0; i < rows; i++) {
      const row = [];
      for (let j = 0; j < cols; j++) {
        row.push(<div key={`${i}-${j}`} className="grid-cell" />);
      }
      grid.push(<div key={i} className="grid-row">{row}</div>);
    }

    return (
      <div className="grid">
        {grid}
      </div>
    );
  }
}

// Testing Home as a component
class Home extends React.Component {
    render() {
      return (
        <div className="container">
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>Jokeish</h1>
            <p>A load of Dad jokes XD</p>
            <p>Sign in to get access </p>
            
          </div>
        </div>
      )
    }
}
  
ReactDOM.render(<App />, document.getElementById('app'));