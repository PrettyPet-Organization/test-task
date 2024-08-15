import { useState } from "react";

type Option = {
  value: string;
  label: string;
};

const options: Option[] = [
  {
    value: "test 1",
    label: "test 1",
  },
  {
    value: "test 2",
    label: "test 2",
  },
  {
    value: "test 3",
    label: "test 3",
  },
];

const App = () => {
  const [isCollapsed, setIsCollapsed] = useState(true);
  const [dropdownValue, setDropdownValue] = useState<string | null>(null);

  return (
    <div>
      <div>
        <h1
          style={{ cursor: "pointer" }}
          onClick={() => setIsCollapsed((prev) => !prev)}
        >
          dropdown
        </h1>

        {dropdownValue && <h2>Значение выпадашки: {dropdownValue}</h2>}
        {!isCollapsed && (
          <ul>
            {options.map((option) => (
              <li
                style={{ cursor: "pointer" }}
                onClick={() => {
                  setDropdownValue(option.value);
                  setIsCollapsed(true);
                }}
                key={option.value}
              >
                {option.label}
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
};

export { App };
