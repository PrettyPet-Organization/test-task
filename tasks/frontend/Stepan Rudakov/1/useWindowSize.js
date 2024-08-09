import { useState, useEffect } from 'react';

// Кастомный хук useWindowSize
const useWindowSize = () => {
  // Состояние для хранения размеров окна
  const [windowSize, setWindowSize] = useState({
    width: window.innerWidth,
    height: window.innerHeight,
  });

  useEffect(() => {
    // Функция для обновления размеров окна
    const handleResize = () => {
      setWindowSize({
        width: window.innerWidth,
        height: window.innerHeight,
      });
    };

    // Добавление обработчика события resize
    window.addEventListener('resize', handleResize);

    // Удаление обработчика при размонтировании компонента
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []); // Пустой массив зависимостей означает, что эффект выполнится только один раз

  // Возвращаем текущее значение ширины и высоты окна
  return windowSize;
};

export default useWindowSize;