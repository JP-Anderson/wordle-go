export default function GameOutcomeOverlay(props) {
	console.log(props);
	if (props.isOpen) return (
	  <div id="outcomeModal" className="modal">
	    <h3>Game finished</h3>
	    <button>New game</button>
	  </div>
	)
}
