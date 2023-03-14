import WordleRow from './WordleRow';

export default function WordleGrid(props) {
	console.log(props);
	if ( props.data.guesses === undefined ) return <p>Loading...</p>;
	const firstEmptyIndex = props.data.guesses.findIndex((item) => item === null);
	return (
	  <div className="wordle-container"><table>
	    <tbody>
	      {props.data.guesses.map((item, index) => {
	        const isGuessingRow = firstEmptyIndex === index;   
	        return (
	           <WordleRow key={index} guess={item} targetLength={props.data.target_length} isGuessingRow={isGuessingRow} buffer={props.buffer} />
	         );
	      })}
	    </tbody>	
	  </table></div>
	);
}
