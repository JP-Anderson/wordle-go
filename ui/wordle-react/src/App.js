import './App.css';
import { useState, useEffect, useRef } from 'react';
import { postGame, postGuess } from './wordleRest';
import WordleGrid from './WordleGrid';
import GameOutcomeOverlay from './GameOutcomeOverlay';
import { setEnterEventFunction, onKeyPress} from './keyPress';

function App({userId}) {
  const [data, setData] = useState({});
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
  const ref = useRef(null);
  
  useEffect(() => {
    // focus to enable the onKeyPress handler
    ref.current.focus();
    postGame(userId).then((response) => handleWordleApiResponse(response));
  }, [userId]);

  const makeGuess = (guessWord) => {
    console.log("Posting guess..." + guessWord);
    postGuess(userId, guessWord).then((response) => handleWordleApiResponse(response));
  }
  setEnterEventFunction(makeGuess);
  return (
    <div ref={ref} tabIndex={-1} className="App" onKeyDown={onKeyPress}>
        {modalOpen && <GameOutcomeOverlay isOpen={modalOpen} data={data} />}
	<WordleGrid data={data} />
	<button onClick={() => setModalOpen(true)}>TEST MODAL</button>
    </div>
  );
}

export default App;
