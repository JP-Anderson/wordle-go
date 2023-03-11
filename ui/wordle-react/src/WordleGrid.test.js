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
     const root = createRoot(container)
     root.render(<WordleGrid guesses={[null,null,null,null,null]} />);
  });
  expect(container.querySelectorAll('tr')).toHaveLength(5);
});
