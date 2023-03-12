import WordleRow from './WordleRow';
import { useEffect } from 'react';
export default function WordleGrid(props) {
	console.log(props);
	if ( props.data.guesses === undefined ) return <p>Loading...</p>;
	return (
	  <div className="wordle-container"><table>
	    <tbody>
	      {props.data.guesses.map((item, index) => {
	         return (
	           <WordleRow key={index} guess={item} targetLength={props.data.target_length} />
	         );
	      })}
	    </tbody>	
	  </table></div>
	);
}
