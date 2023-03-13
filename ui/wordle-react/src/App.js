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
  
  useEffect(() => {
    postGame(userId, setData);
  }, []);

  const guessOnClick = () => {
    let guessWord = guess;
    postGuess(userId, guessWord).then(result => {
      setData(result);
    });
  };

  return (
    <div className="App">
	<WordleGrid data={data} />
	<input type="text" id="guessInput" onChange={handleGuessInputChange} value={guess}></input>
	<button onClick={guessOnClick}>Guess</button>
	<button onClick={() => setModalOpen(true)}>TEST MODAL</button>
	{modalOpen && <GameOutcomeOverlay isOpen={modalOpen} />}
    </div>
  );
}

export default App;
