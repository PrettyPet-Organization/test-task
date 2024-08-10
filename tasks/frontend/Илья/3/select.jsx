import { useState } from "react";
import "./select.css";

export const Select = ({
  options,
  selectedOption,
  getOptionLabel,
  getOptionValue,
  onSelect,
}) => {
  const [isOpen, setIsOpen] = useState(false);

  const handleSelect = (option) => {
    onSelect(option);
    setIsOpen(false);
  };

  return (
    <div className="custom-select">
      <div
        className="custom-select__field"
        onClick={() => setIsOpen((prevValue) => !prevValue)}
      >
        <span className="custom-select_field__option-label option-label">
          {getOptionLabel(selectedOption)}
        </span>
        <span
          className={`custom-select__field__arrow arrow ${isOpen && "open"}`}
        />
      </div>
      <ul
        className={`custom-select__dropdown-list dropdown-list 
          ${isOpen && "open"}`}
      >
        {options.map((option) => (
          <li
            key={getOptionValue(option)}
            className="dropdown-list__list-item list-item"
            onClick={() => handleSelect(option)}
          >
            {getOptionLabel(option)}
          </li>
        ))}
      </ul>
    </div>
  );
};
