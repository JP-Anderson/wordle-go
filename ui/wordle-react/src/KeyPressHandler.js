import { useEffect } from 'react';
import { onKeyPress } from './keyPress';

export function KeyPressHandler(props) {
  useEffect(() => {
    window.addEventListener("keydown", onKeyPress);
    return () => {
      window.removeEventListener("keydown", onKeyPress);
    }
  }, []);
}
