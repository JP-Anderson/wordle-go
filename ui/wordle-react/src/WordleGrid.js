import WordleRow from './WordleRow';

export default function WordleGrid(props) {
	return (
	  <table>
	    <tbody>
	      {props.guesses.map((item, index) => {
	         return (
	           <WordleRow key={index} guess={item} targetLength={props.targetLength} />
	         );
	      })}
	    </tbody>	
	  </table>
	);
}

