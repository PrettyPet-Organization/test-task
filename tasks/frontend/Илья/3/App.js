import React, { useState } from "react";
import { Select } from "./select";

const users = [
  { id: 1, name: "Alice" },
  { id: 2, name: "Bob" },
  { id: 3, name: "Charlie" },
];

const App = () => {
  const [selectedUser, setSelectedUser] = useState(users[0]);

  return (
    <div>
      <Select
        options={users}
        selectedOption={selectedUser}
        getOptionValue={(user) => user.id}
        getOptionLabel={(user) => user.name}
        onSelect={(user) => setSelectedUser(user)}
      />
    </div>
  );
};

export default App;
