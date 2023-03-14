import './App.css';
import { useState, useEffect } from 'react';
import { postGame, postGuess } from './wordleRest';
import WordleGrid from './WordleGrid';
import GameOutcomeOverlay from './GameOutcomeOverlay';
import {onKeyPress} from './keyPress';

function App({userId}) {
  const [data, setData] = useState({});
  const [guess, setGuess] = useState("");
  const [modalOpen, setModalOpen] = useState(false);
  document.onkeydown = onKeyPress;
  
  const handleGuessInputChange = event => {
    setGuess(event.target.value);
  }

  const handleWordleApiResponse = (response) => {
    if ( response === undefined || (response.skip !== undefined && response.skip) ) {
      return;
    }
    setData(response)
    if (response.game_state === 1 || response.game_state === 2) {
      setModalOpen(true);
    }
  }
  
  useEffect(() => {
    postGame(userId).then((response) => handleWordleApiResponse(response));
  }, [userId]);

  const guessOnClick = () => {
    let guessWord = guess;
    postGuess(userId, guessWord).then((response) => handleWordleApiResponse(response));
  };

  return (
    <div className="App">
        {modalOpen && <GameOutcomeOverlay isOpen={modalOpen} data={data} />}
	<WordleGrid data={data} />
	<input type="text" id="guessInput" onChange={handleGuessInputChange} value={guess}></input>
	<button onClick={guessOnClick}>Guess</button>
	<button onClick={() => setModalOpen(true)}>TEST MODAL</button>
    </div>
  );
}

export default App;
