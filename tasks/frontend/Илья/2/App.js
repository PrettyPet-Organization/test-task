import React, { useMemo, useState } from "react";

const UserList = ({ users }) => {
  const [search, setSearch] = useState("");

  const filterUsers = () =>
    users.filter((user) =>
      user.name.toLowerCase().includes(search.toLocaleLowerCase())
    );

  const filteredUsers = useMemo(filterUsers, [search, users]);

  const handleSearchChange = (changedText) => {
    setSearch(changedText);
  };

  return (
    <div>
      <input
        type="text"
        value={search}
        onChange={(e) => handleSearchChange(e.target.value)}
        placeholder="Search users..."
      />
      <ul>
        {filteredUsers.map((user) => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
};

const App = () => {
  const users = [
    { id: 1, name: "Alice" },
    { id: 2, name: "Bob" },
    { id: 3, name: "Charlie" },
  ];

  return (
    <div>
      <h1>User List</h1>
      <UserList users={users} />
    </div>
  );
};

export default App;
