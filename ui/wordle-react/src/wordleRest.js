
module.exports = {
  postGame,
  postGuess,
}

const localApiUrl = 'localhost:8080';
const gameEndpoint = '/game';
const guessEndpoint = '/guess';

async function postGame (id, setStateFunc) {
  await fetch('http://' + localApiUrl + gameEndpoint, {
    method: 'POST',
    body: JSON.stringify({
      user_id: id,
    }),
    headers: {
     'Content-type': 'application/json; charset=UTF-8',
    },
  })
  .then((response) => response.json())
  .then((js) => setStateFunc(js))
  .catch((error) => console.log(error))
}

async function postGuess(userID, guessWord) {
  return await fetch('http://' + localApiUrl + guessEndpoint, {
    method: 'POST',
    body: JSON.stringify({
      user_id: userID,
      guess: guessWord,
    }),
    headers: {
      'Content-type': 'application/json; charset=UTF-8',
    },
  })
  .then((response) => response.json())
  .catch((error) => console.log(error))
}
