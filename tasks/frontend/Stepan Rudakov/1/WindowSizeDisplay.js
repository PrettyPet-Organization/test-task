import React from 'react';
import useWindowSize from './useWindowSize'; // Импортируем кастомный хук

const WindowSizeDisplay = () => {
  const { width, height } = useWindowSize(); // Используем хук для получения размеров окна

  return (
    <div>
      <h1>Текущие размеры окна:</h1>
      <p>Ширина: {width}px</p>
      <p>Высота: {height}px</p>
    </div>
  );
};

export default WindowSizeDisplay;