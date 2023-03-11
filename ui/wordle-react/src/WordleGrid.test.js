import React from "react";
import { createRoot } from "react-dom/client"
import { render, unmountComponentAtNode } from "react-dom";
import { act } from "react-dom/test-utils";
import WordleGrid from "./WordleGrid.js";

let container = null;
beforeEach(() => {
  // setup a DOM element as a render target
  container = document.createElement("div");
  document.body.appendChild(container);
});

afterEach(() => {
  // cleanup on exiting
  container.remove();
  container = null;
});

it("has tr for each array element", () => {
  act(() => {
     const root = createRoot(container);
     root.render(<WordleGrid targetLength='5' guesses={[null,null,null,null,null]} />);
  });
  expect(container.querySelectorAll('tr')).toHaveLength(5);
});


it("has tds for length of word", () => {
  act(() => {
     const root = createRoot(container);
     root.render(<WordleGrid targetLength='5' guesses={[null,null,null,null,null]} />);
  });
  const rows = container.querySelectorAll('tr');
  rows.forEach((row) => {
    expect(row.querySelectorAll('td')).toHaveLength(5);
  });
});

it("places guess letter in each td", () => {
  act(() => {
    const root = createRoot(container);
    root.render(<WordleGrid targetLength='5' guesses={[["snack", "22222"] ,null,null,null,null]} />);
  });
  const rows = container.querySelectorAll('tr');
  const firstGuess = rows[0];
  const expectedChars = [ "s", "n", "a", "c", "k"];
  const tableDatas = firstGuess.querySelectorAll('td');
  [...tableDatas].map( (item, index) => {
    expect(item.innerHTML).toBe(expectedChars[index]);
  });
});
