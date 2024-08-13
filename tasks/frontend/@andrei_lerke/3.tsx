import React, { useState, useEffect, useCallback } from 'react';
import './3.css';

interface IItem {
  id: number;
  name: string;
}

const items: IItem[] = [
  { id: 1, name: 'BMV' },
  { id: 2, name: 'Mercedes' },
  { id: 3, name: 'Volvo' },
  { id: 4, name: 'VolksWagen' },
];

export const DropList: React.FC = () => {
  const [showMenu, setShowMenu] = useState<boolean>(false);
  const [selectedItem, setSelectedItem] = useState<IItem | null>(null);

  const handleKeyDown = useCallback(
    (event: { code: string }) => {
      if (event.code === 'Escape') {
        setShowMenu(false);
      }
    },
    [setShowMenu],
  );

  useEffect(() => {
    if (showMenu) window.addEventListener('keydown', handleKeyDown);

    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };
  }, [showMenu, handleKeyDown]);

  const handleBackdropClick: React.MouseEventHandler<
    HTMLDivElement
  > = event => {
    if (event.currentTarget === event.target) {
      event.preventDefault();
      setShowMenu(false);
    }
  };

  const handleButtonClick = () => {
    setShowMenu(!showMenu);
  };

  const handleItemClick = (item: IItem) => {
    setSelectedItem(item);
    setShowMenu(!showMenu);
  };

  return (
    <div className="dropdown" onClick={handleBackdropClick}>
      <button type="button" onClick={handleButtonClick} className="dropbtn">
        Click me
      </button>

      {showMenu && (
        <ul id="myDropdown" className="dropdown-content">
          {items.map(i => (
            <li key={i.id} value={i.name} onClick={() => handleItemClick(i)}>
              {i.name}
            </li>
          ))}
        </ul>
      )}

      <p>{selectedItem ? `Auto: ${selectedItem.name}` : 'Select a car...'}</p>
    </div>
  );
};
