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

  postNewGame("1");

  return (
    <div className="App">
	<h1>Wordle</h1>
	<table><tbody>
		<tr><td><div class="correct-letter">a</div></td><td><div class="wrong-location-letter">b</div></td><td></td><td></td><td></td></tr>
		<tr><td></td><td></td><td></td><td></td><td></td></tr>
		<tr><td></td><td></td><td></td><td></td><td></td></tr>
		<tr><td></td><td></td><td></td><td></td><td></td></tr>
		<tr><td></td><td></td><td></td><td></td><td></td></tr>
	</tbody></table>
	<input></input>
	<button>Guess</button>
    </div>
  );
}

export default App;
