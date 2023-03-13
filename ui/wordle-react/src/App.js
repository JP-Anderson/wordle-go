import logo from './logo.svg';
import './App.css';
import { useState, useEffect } from 'react';
import { postGame, postGuess } from './wordleRest';
import WordleGrid from './WordleGrid';
import GameOutcomeOverlay from './GameOutcomeOverlay';

function App({userId}) {
  const [data, setData] = useState({});
  const [guess, setGuess] = useState("");
  const [modalOpen, setModalOpen] = useState(false);

  const handleGuessInputChange = event => {
    setGuess(event.target.value);
  }

  const handleWordleApiResponse = (data) => {
    console.log("Wordle response: " + data);
    setData(data)
    if (data.game_state === 1 || data.game_state === 2) {
      setModalOpen(true);
    }
  }
  
  useEffect(() => {
    postGame(userId, handleWordleApiResponse);
  }, []);

  const guessOnClick = () => {
    let guessWord = guess;
    postGuess(userId, guessWord).then(result => {
      handleWordleApiResponse(result);
    });
  };

  return (
    <div className="App">
        {modalOpen && <GameOutcomeOverlay isOpen={modalOpen} />}
	<WordleGrid data={data} />
	<input type="text" id="guessInput" onChange={handleGuessInputChange} value={guess}></input>
	<button onClick={guessOnClick}>Guess</button>
	<button onClick={() => setModalOpen(true)}>TEST MODAL</button>
    </div>
  );
}

export default App;
