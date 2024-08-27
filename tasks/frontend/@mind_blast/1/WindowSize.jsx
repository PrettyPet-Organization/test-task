import useWindowSize from '../hooks/useWindowSize';

function WindowSize() {
    const { width, height } = useWindowSize();

    return (
        <div>
            <h1>Размеры окна:</h1>
            <p>Ширина: {width}px</p>
            <p>Высота: {height}px</p>
        </div>
    );
}

export default WindowSize;
