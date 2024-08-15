import { useEffect, useRef, useState } from "react";

type UseDebounceOptions = {
  debounceTime: number;
};

const useDebounce = <T,>(
  value: T,
  options: UseDebounceOptions = { debounceTime: 500 }
): T => {
  const { debounceTime } = options;
  const [debouncedValue, setDebouncedValue] = useState<T>(value);

  const timeoutRef = useRef<number | null>(null);

  useEffect(() => {
    if (timeoutRef.current) {
      clearTimeout(timeoutRef.current);
    }

    timeoutRef.current = setTimeout(() => {
      setDebouncedValue(value);
    }, debounceTime);

    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
    };
  }, [value, debounceTime]);

  return debouncedValue;
};

const useWindowSize = () => {
  const [windowSize, setWindowSize] = useState({
    width: window.innerWidth,
    height: window.innerHeight,
  });

  useEffect(() => {
    const handleResize = () => {
      setWindowSize({
        width: window.innerWidth,
        height: window.innerHeight,
      });
    };

    window.addEventListener("resize", handleResize);

    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  return windowSize;
};

const MainPage = () => {
  const windowSize = useWindowSize();

  const test = useDebounce(windowSize, { debounceTime: 1000 });

  return (
    <div>
      {test.height}

      {test.width}
    </div>
  );
};

export { MainPage };
