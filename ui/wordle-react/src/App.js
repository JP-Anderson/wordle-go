import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import WordleGrid from './WordleGrid';

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

  return (
    <div className="App">
	<h1>Wordle</h1>
	<WordleGrid data={data} />
	<table><tbody>
		<tr><td><div class="correct letter">C</div></td><td><div class="wrong letter">R</div></td><td><div class="misplaced letter">A</div></td><td><div class="wrong letter">N</div></td><td><div class="misplaced letter">E</div></td></tr>
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
