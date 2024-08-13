import { useWindowSize } from "./useWindowSize";

export const WindowSizeViewer = () => {
  const { windowSize } = useWindowSize();

  return (
    <div>
      <h1>Размеры окна:</h1>
      <p>Ширина: {windowSize.width}</p>
      <p>Высота: {windowSize.height}</p>
    </div>
  );
};
