export default function WordleRow(props) {
	return(
	  <tr>
	    {Array.from({ length: props.targetLength }, (_, i) => {
	      if ( !props.guess ) return <td key={i}></td>;
	      else return <td class="correct letter" key={i}>{props.guess[0][i]}</td>
	    })
	    }
	  </tr>
	);
}
