export default function GameOutcomeOverlay(props) {
	console.log(props);
	if (props.isOpen) {
		const message = props.data.game_state == 2 ? "You won!" : "You lost!";
 		return (
		  <div id="outcomeModal" className="modal">
		    <div className="modal-content">
		      <h3>Game finished</h3>
		        <span>{message}</span>
		      <button>New game</button>
		    </div>
		  </div>
		)
	}
}
