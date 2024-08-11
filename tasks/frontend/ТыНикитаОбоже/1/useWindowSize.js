import {useEffect, useState} from "react";

export function useWindowSize() {
    const [width, setWidth] = useState(window.outerWidth)
    const [height, setHeight] = useState(window.outerHeight)

    const resize = () => {
        setWidth(window.outerWidth)
        setHeight(window.outerHeight)
    }

    useEffect(() => {
        window.addEventListener('resize', resize);

        return () => window.removeEventListener('resize', resize)
    }, []);

    return {width, height}
}