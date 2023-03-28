import { useState, useEffect } from 'react';
import App from './App.js';

export default function LogIn() {
  const [userID, setUserID] = useState("");  
  const [input, setInput] = useState("");
  useEffect(() => {
  }, []);

  const logInClicked = () => {
    alert(input);
    setUserID(input);
  }

  const inputChanged = (event) => {
    setInput(event.target.value);
  }

  if (userID === "") {
    return <div className="App"><input id="userIDInput" type="text" onChange={inputChanged} value={input}/><button onClick={logInClicked} >Log In</button></div>;
  }
  else {
    return <App userID={userID}/>;
  }
}
