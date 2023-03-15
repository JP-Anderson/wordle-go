import { KeyPressHandler } from './KeyPressHandler';

export default function WordleRow(props) {
	console.log(props);
	return(
	  <tr>
	    {Array.from({ length: props.targetLength }, (_, i) => {
	      if ( !props.guess ) {
	        if (props.isGuessingRow) {
		  return <td key={i}><div className="letter empty active-guessing-row"><KeyPressHandler />{props.buffer[i]}</div></td>;
		}
	        return <td key={i}><div className="letter empty"></div></td>;
	      }
	      else {
	         const guessRowClasses = statusIntToClasses[props.guess.letter_statuses[i]] + " letter"
	         return <td key={i}><div className={guessRowClasses}>{props.guess.guess_word[i]}</div></td>
	      }
	    })
	    }
	  </tr>
	);
}

const statusIntToClasses = { "2" : "correct", "1" : "misplaced", "0" : "incorrect" };
