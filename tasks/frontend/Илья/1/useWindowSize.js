import { useEffect, useState } from "react";

const RESIZE_EVENT = "resize";

export const useWindowSize = () => {
  const [windowSize, setWindowSize] = useState({
    width: window.screen.width,
    height: window.screen.height,
  });

  function handleChangeWindowSize() {
    setWindowSize({
      width: window.screen.width,
      height: window.screen.height,
    });
  }

  useEffect(() => {
    window.addEventListener(RESIZE_EVENT, handleChangeWindowSize);

    return () =>
      window.removeEventListener(RESIZE_EVENT, handleChangeWindowSize);
  }, []);

  return { windowSize };
};
