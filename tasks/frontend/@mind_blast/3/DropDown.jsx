import { useState } from 'react';
import './DropDown.css';

const DropDown = ({ options }) => {
    const [selectedOption, setSelectedOption] = useState(null);
    const [isOpen, setIsOpen] = useState(false);

    const handleOptionClick = option => {
        setSelectedOption(option);
        setIsOpen(false);
    };

    return (
        <div className="dropdown">
            <div className="dropdown-choosen" onClick={() => setIsOpen(!isOpen)}>
                {selectedOption ? selectedOption.label : 'Выбрать'}
            </div>
            {isOpen && (
                <ul className="dropdown-menu">
                    {options.map(option => (
                        <li key={option.value} onClick={() => handleOptionClick(option)}>
                            {option.label}
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default DropDown;
