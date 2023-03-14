import WordleRow from './WordleRow';
import { buffer } from './keyPress';
import { useEffect } from 'react';

export default function WordleGrid(props) {
	console.log(props);
	if ( props.data.guesses === undefined ) return <p>Loading...</p>;
	const firstEmptyIndex = props.data.guesses.findIndex((item) => item === null);
	return (
	  <div className="wordle-container"><table>
	    <tbody>
	      {props.data.guesses.map((item, index) => {
	         return (
	           <WordleRow key={index} buffer={buffer} guess={item} targetLength={props.data.target_length} isGuessingRow={firstEmptyIndex === index}/>
	         );
	      })}
	    </tbody>	
	  </table></div>
	);
}
