import refreshPage from './refresh.js';

export default function GameOutcomeOverlay(props) {
	console.log(props);
	if (props.isOpen) {
		const target = props.data.answer.answer;
		const message = props.data.game_state == 2 ? "You won!" : "You lost! The answer was "+target;
 		return (
		  <div id="outcomeModal" className="modal">
		    <div className="modal-content">
		      <h3>Game finished</h3>
		        <span>{message}</span>
		      <button onClick={refreshPage}>New game</button>
		    </div>
		  </div>
		)
	}
}
