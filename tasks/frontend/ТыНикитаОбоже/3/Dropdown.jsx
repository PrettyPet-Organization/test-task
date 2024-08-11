import {createContext, useContext, useEffect, useState} from "react";
import "./style.css";

const ValueContext = createContext(null);

export function DropdownMenu({children}){
    return (
        <div className="dropdown_menu">
            <ul>{children}</ul>
        </div>
    )
}

export function DropdownItem({children}){
    const {select} = useContext(ValueContext);
    return <li onClick={() => select(children)}>{children}</li>
}

export function DropdownToggle({children, onClick}){
    return <button onClick={onClick}>{children}</button>
}


export function Dropdown({children}){
    const [isOpen, setIsOpen] = useState(false);
    const [selected, setSelected] = useState(null);

    const toggleDropdown = () => {
        setIsOpen(!isOpen);
    };

    const select = (value) => {
        setSelected(value)
        setIsOpen(false);
    }


    return (
        <ValueContext.Provider value={{select}}>
            <div className="dropdown">
                <DropdownToggle onClick={toggleDropdown}>
                    {selected || "Select an option"}
                </DropdownToggle>
                {isOpen && <DropdownMenu>{children}</DropdownMenu>}
            </div>
        </ValueContext.Provider>

    );
}

export function Container() {
    return (
        <div>
            <Dropdown>
                <DropdownItem>a</DropdownItem>
                <DropdownItem>b</DropdownItem>
                <DropdownItem>c</DropdownItem>
                <DropdownItem>c2222</DropdownItem>
            </Dropdown>
        </div>
    )
}



