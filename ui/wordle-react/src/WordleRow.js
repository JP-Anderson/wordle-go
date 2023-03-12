export default function WordleRow(props) {
	console.log(props);
	return(
	  <tr>
	    {Array.from({ length: props.targetLength }, (_, i) => {
	      if ( !props.guess ) return <td key={i}></td>;
	      else return <td class={statusIntToClasses[props.guess[1][i]]} key={i}>{props.guess[0][i]}</td>
	    })
	    }
	  </tr>
	);
}

const statusIntToClasses = { "2" : "correct letter", "1" : "misplaced letter", "0" : "incorrect letter" };
