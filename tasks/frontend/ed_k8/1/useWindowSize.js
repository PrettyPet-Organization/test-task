import { useState, useEffect } from "react";

function useWindowSize() {
	const [windowSize, setWindowSize] = useState({
		width: window.innerWidth,
		height: window.innerHeight,
	});

	useEffect(() => {
		// Обработчик изменения размера окна
		const handleResize = () => {
			setWindowSize({
				width: window.innerWidth,
				height: window.innerHeight,
			});
		};

		// Добавляем слушатель события resize
		window.addEventListener("resize", handleResize);

		// Удаляем слушатель события при размонтировании компонента
		return () => window.removeEventListener("resize", handleResize);
	}, []); // Пустой массив зависимостей для запуска эффекта только при монтировании

	return windowSize;
}

export default useWindowSize;
