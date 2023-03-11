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

it("td gets correct css classes for each letter correctness", () => {
  act(() => {
    const root = createRoot(container);
    root.render(<WordleGrid targetLength='5' guesses={[["crane", "10210"], ["snack", "22222"] ,null,null,null,null]} />);
  });
  const rows = container.querySelectorAll('tr');

  const guess1 = rows[0];
  const expectedChars1 = [ "c", "r", "a", "n", "e"];
  const expectedClass1 = [ "misplaced", "incorrect", "correct", "misplaced", "incorrect" ];
  const tableDatas1 = guess1.querySelectorAll('td');
  [...tableDatas1].map( (item, index) => {
    expect(item.innerHTML).toBe(expectedChars1[index]);
    expect(item).toHaveClass("letter");
    expect(item).toHaveClass(expectedClass1[index]);
  });
  
  const guess2 = rows[1];
  const expectedChars2 = [ "s", "n", "a", "c", "k"];
  const expectedClass2 = ["correct", "correct", "correct", "correct", "correct"];
  const tableDatas2 = guess2.querySelectorAll('td');
  [...tableDatas2].map( (item, index) => {
    expect(item.innerHTML).toBe(expectedChars2[index]);
    expect(item).toHaveClass("letter");
    expect(item).toHaveClass(expectedClass2[index]);
  });
});
