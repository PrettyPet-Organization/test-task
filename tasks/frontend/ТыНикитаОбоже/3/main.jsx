import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from '../1/App.jsx'
import Users from "../2/Users.jsx";
import {Container} from "./Dropdown.jsx";

createRoot(document.getElementById('root')).render(
  <StrictMode>
    {/*<App />*/}
    {/*<Users />*/}
      <Container />
  </StrictMode>,
)
