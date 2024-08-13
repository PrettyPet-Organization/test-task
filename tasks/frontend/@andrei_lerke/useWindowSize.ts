import { useState, useEffect } from 'react';

export const useWindowSize = () => {
  const [width, setWidth] = useState<number>(window.innerWidth);
  const [height, setHeight] = useState<number>(window.innerHeight);

  useEffect(() => {
    const handleResize = event => {
      setWidth(event.target.innerWidth);
      setHeight(event.target.innerHeight);
    };
    window.addEventListener('resize', handleResize);
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  });

  return { width, height };
};
