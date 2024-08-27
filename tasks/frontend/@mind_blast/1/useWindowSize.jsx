import { useState, useEffect } from 'react';
import { throttle } from 'lodash';

function useWindowSize() {
    const [windowSize, setWindowSize] = useState({
        width: window.innerWidth,
        height: window.innerHeight
    });

    useEffect(() => {
        const handleResize = throttle(() => {
            setWindowSize({
                width: window.innerWidth,
                height: window.innerHeight
            });
        }, 200);
        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    return windowSize;
}

export default useWindowSize;
