import React from 'react';
import { useWindowSize } from './useWindowSize.ts';

export const Component: React.FC = () => {
  const { width, height } = useWindowSize();

  return (
    <>
      <p>width: {width}px</p>
      <p>height: {height}px</p>
    </>
  );
};
