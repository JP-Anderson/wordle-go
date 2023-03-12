export default function WordleRow(props) {
	console.log(props);
	return(
	  <tr>
	    {Array.from({ length: props.targetLength }, (_, i) => {
	      if ( !props.guess ) return <td key={i}><div className="letter"></div></td>;
	      else return <td key={i}><div className={statusIntToClasses[props.guess.letter_statuses[i]]}>{props.guess.guess_word[i]}</div></td>
	    })
	    }
	  </tr>
	);
}

const statusIntToClasses = { "2" : "correct letter", "1" : "misplaced letter", "0" : "incorrect letter" };
