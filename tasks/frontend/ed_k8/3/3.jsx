import React, { useState, useRef, useEffect } from 'react';
// import "./3.css"; 
const Dropdown = ({ options, onSelect }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [selectedOption, setSelectedOption] = useState(null);
  const dropdownRef = useRef(null);

  const toggleDropdown = () => setIsOpen(!isOpen);

  const handleOptionClick = (option) => {
    setSelectedOption(option);
    setIsOpen(false);
    if (onSelect) {
      onSelect(option);
    }
  };

  // Закрытие выпадающего списка при клике вне его
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        setIsOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <div className="dropdown" ref={dropdownRef}>
      <button onClick={toggleDropdown} className="dropdown-button">
        {selectedOption ? selectedOption.label : 'Select an option'}
      </button>
      {isOpen && (
        <ul className="dropdown-menu">
          {options.map(option => (
            <li
              key={option.value}
              onClick={() => handleOptionClick(option)}
              className="dropdown-item"
            >
              {option.label}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

// Пример использования компонента Dropdown
const App = () => {
  const options = [
    { value: '1', label: 'Option 1' },
    { value: '2', label: 'Option 2' },
    { value: '3', label: 'Option 3' },
  ];

  const handleSelect = (option) => {
    console.log('Selected:', option);
  };

  return (
    <div>
      <h1>Dropdown Example</h1>
      <Dropdown options={options} onSelect={handleSelect} />
    </div>
  );
};

export default App;
