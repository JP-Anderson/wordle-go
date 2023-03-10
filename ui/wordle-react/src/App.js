import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';

function App() {
  const id = "1"
  const postNewGame = async (userId) => {
    await fetch('http://localhost:8080/game', {
      method: 'POST',
      body: JSON.stringify({
        user_id: userId,
      }),
      headers: {
        'Content-type': 'application/json; charset=UTF-8',
      },
    })
    .then((response) => response.json())
    .then((data) => console.log(data));
  };

  const handleNewGame= (e) => {
    e.preventDefault();
    postNewGame("1");
  }
  return (
    <div className="App">
	<button onClick={handleNewGame}>New Game</button>
    </div>
  );
}

export default App;
