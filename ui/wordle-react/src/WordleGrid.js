export default function WordleGrid(props) {
	return (
	  <table>
	    <tbody>
	      {props.guesses.map((item, index) => {
	         return (
	           <tr key={index}>
	           </tr>
	         );
	      })}
	    </tbody>	
	  </table>
	);
}
