import React from "react";
import ReactDOM from "react-dom";
 
// コンポーネント
class Test extends React.Component {
  render(){
    return (
      <h3>React.js-App</h3>
    );
  }
}
// レンダリング
ReactDOM.render(
  <Test />,
  document.getElementById('container')
);