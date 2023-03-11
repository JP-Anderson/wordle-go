export default function WordleRow(props) {
	return(
	  <tr>
	    {Array.from({ length: props.targetLength }, (_, i) => <td key={i}>A</td>)}
	  </tr>
	);
}
