import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import WordleGrid from './WordleGrid';

async function guess(userID, guessWord) {
  // param is a highlighted word from the user before it clicked the button
  return await fetch("http://localhost:8080/guess", {
    method: 'POST',
    body: JSON.stringify({
      user_id: userID,
      guess: guessWord,
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8',
    },
  })
  .then((response) => response.json())
  .catch((error) => console.log(error))
}


function App({userId}) {
  const [data, setData] = useState({});
  console.log("App start!");
  useEffect(() => {
    console.log("Hit useEffect...");
    const postNewGame = async (id) => {
      console.log("postNewGame---------------------");
      await fetch('http://localhost:8080/game', {
        method: 'POST',
        body: JSON.stringify({
          user_id: id,
        }),
        headers: {
          'Content-type': 'application/json; charset=UTF-8',
        },
      })
      .then((response) => response.json())
      .then((js) => setData(js))
      .catch((error) => console.log(error))
    }
    postNewGame(userId);
  }, []);

  const guessOnClick = () => {
    // todo: get this from input
    let guessWord = "crane";
    guess(userId, guessWord).then(result => {
      setData(result);
    });
  };

  return (
    <div className="App">
	<h1>Wordle</h1>
	<WordleGrid data={data} />
	<input></input>
	<button onClick={guessOnClick}>Guess</button>
    </div>
  );
}

export default App;
