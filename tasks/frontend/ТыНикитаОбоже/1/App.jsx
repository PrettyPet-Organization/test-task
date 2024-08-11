import {useWindowSize} from "./useWindowSize.js";

function App() {
    let {width, height} = useWindowSize();

    return (
        <div>
            {width} : {height}
        </div>
    )
}

export default App
