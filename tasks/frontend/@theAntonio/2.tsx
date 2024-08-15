import React, { useState } from "react";

// добавить типизацию для users
const UserList = ({ users }) => {
  const [search, setSearch] = useState("");
  const [filteredUsers, setFilteredUsers] = useState(users);

  // добавить типизацию ChangeEvent<HTMLInputElement>
  // если input будет отдельным компонентом, то обернуть в useCallback, чтобы не перерисовать каждый раз функцию
  const handleSearchChange = (e) => {
    setSearch(e.target.value);
    // куча лишних ререндеров
    // убрать вообще useState и сделать обычную фильтрацию юзеров и по ней проходить через .map
    setFilteredUsers(
      users.filter((user) =>
        user.name.toLowerCase().includes(e.target.value.toLowerCase())
      )
    );
  };

  return (
    <div>
      {/* можно вынести в отдельный компонент и обернуть в memo чтобы не тригерить каждый раз ререндер */}
      <input
        type="text"
        value={search}
        onChange={handleSearchChange}
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

// сделать именованный экспорт
const App = () => {
  // вынести вне тела компонента
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

// убрать экспорт дефолт
export default App;
