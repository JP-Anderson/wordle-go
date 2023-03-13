
module.exports = {
  postGame,
  postGuess,
}

const localApiUrl = 'localhost:8080';
const gameEndpoint = '/game';
const guessEndpoint = '/guess';

function handleResponse(response) {
  if (!response.ok) {
    console.log(response);
  } else {
    return response.json();
  }
}

async function postGame (id) {
  return await fetch('http://' + localApiUrl + gameEndpoint, {
    method: 'POST',
    body: JSON.stringify({
      user_id: id,
    }),
    headers: {
     'Content-type': 'application/json; charset=UTF-8',
    },
  })
  .then((response) => {
    if (!response.ok) {
      // 400 response for POST Game means a Game already exists for user ID.
      // Therefore in this case we can just hit the GET endpoint for user ID.
      // In the future we could just update the POST endpoint to return the model,
      // but I like this solution for now as it makes the API response clearer.
      if (response.status === 400) {
        return getGame(id);
      }
    }
    else handleResponse(response);
  });
}

async function getGame(id) {
  return await fetch('http://' + localApiUrl + gameEndpoint + "/" + id, {
    method: 'GET',
    headers: {
      'Content-type': 'application/json; charset=UTF-8',
    },
  })
  .then((response) => handleResponse(response));
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
  .then((response) => handleResponse(response));
}
