let buffer = "";
let bufferLim = 5;

let enterEventFunction = function () {};
let propsCallBackFunction = function () {};

export function setEventToPropsCB(f) {
	propsCallBackFunction = f;
}

export function setEnterEventFunction(f) {
	enterEventFunction = f;
}

export function onKeyPress(event) {
	console.log(event);
	event = event || window.event;
	// BACKSPACE
	if (event.keyCode === 8) {
		if (buffer.length > 0) {
			buffer = buffer.slice(0, buffer.length-1);
			propsCallBackFunction(buffer);
		}
	}
	// ENTER
	else if (event.keyCode === 13) {
		console.log("enter");
		if ( buffer.length === bufferLim) {
			let guess = buffer;
			buffer = "";
			propsCallBackFunction(buffer);
			enterEventFunction(guess);
		}
	}
	// LETTERS
	else if (event.keyCode in keys) {
		if (buffer.length < bufferLim) {
			buffer = buffer + keys[event.keyCode]
			propsCallBackFunction(buffer);
		}
	}
	console.log(buffer);
}

const keys = {
  65: "A",
  66: "B",
  67: "C",
  68: "D",
  69: "E",
  70: "F",
  71: "G",
  72: "H",
  73: "I",
  74: "J",
  75: "K",
  76: "L",
  77: "M",
  78: "N",
  79: "O",
  80: "P",
  81: "Q",
  82: "R",
  83: "S",
  84: "T",
  85: "U",
  86: "V",
  87: "W",
  88: "X",
  89: "Y",
  90: "Z",
}
