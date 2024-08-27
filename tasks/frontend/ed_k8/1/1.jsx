// import useWindowSize from "../hooks/useWindowSize";
import useWindowSize from "./useWindowSize.js";
function WindowSizeDisplay() {
	const { width, height } = useWindowSize();
	return (
		<div>
			<h1>Current Window Size</h1>
			<p>Width: {width}px</p>
			<p>Height: {height}px</p>
		</div>
	);
}

export default WindowSizeDisplay;

