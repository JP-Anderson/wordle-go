import './App.css';
import { useState, useEffect } from 'react';
import { postGame, postGuess } from './wordleRest';
import WordleGrid from './WordleGrid';
import GameOutcomeOverlay from './GameOutcomeOverlay';
import { setEventToPropsCB, setEnterEventFunction,  onKeyPress } from './keyPress';

function App({userId}) {
  const [data, setData] = useState({});
  const [buffer, setBuffer] = useState("");
  const [modalOpen, setModalOpen] = useState(false);
  
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
    window.addEventListener("keydown", onKeyPress);
    return () => {
      window.removeEventListener("keydown", onKeyPress);
    }
  }, [userId]);

  const makeGuess = (guessWord) => {
    console.log("Posting guess..." + guessWord);
    postGuess(userId, guessWord).then((response) => handleWordleApiResponse(response));
  }
  const eventToPropsCallbackFunction = (newBuffer) => {
    setBuffer(newBuffer);
  }

  setEnterEventFunction(makeGuess);
  setEventToPropsCB(eventToPropsCallbackFunction);
  return (
    <div className="App">
        {modalOpen && <GameOutcomeOverlay isOpen={modalOpen} data={data} />}
	<WordleGrid data={data} buffer={buffer} />
	<button onClick={() => setModalOpen(true)}>TEST MODAL</button>
    </div>
  );
}

export default App;
