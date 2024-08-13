import { useState, useEffect } from "react";

const useWindowSize = () => {
  const [value, setWindowSize] = useState({
    width: window ? window.innerWidth : 0,
    height: window ? window.innerHeight : 0,
  });

  useEffect(() => {
    const handleResize = () => {
      setWindowSize({
        width: window.innerWidth,
        height: window.innerHeight,
      });
    };

    window.addEventListener("resize", handleResize);

    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  return value;
};

export { useWindowSize };
