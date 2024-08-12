import { useWindowSize } from "./windowSize";

function Task() {
  const { width, height } = useWindowSize();

  return (
    <>
      <div>
        <div>Ширина: {width}</div>
        <div>Высота: {height}</div>
      </div>
    </>
  );
}

export default Task;
