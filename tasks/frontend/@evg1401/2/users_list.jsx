import React, { useState, useEffect } from 'react';

const UserList = ({ users }) => {
  const [search, setSearch] = useState('');
  const [filteredUsers, setFilteredUsers] = useState(users);

  useEffect(() => {
    setFilteredUsers(
      users.filter(user => 
        user.name.toLowerCase().includes(search.toLowerCase())
      )
    );
  }, [search, users]);

  const handleSearchChange = (e) => {
    setSearch(e.target.value);
  };

  return (
    <div>
      <input
        type="text"
        value={search}
        onChange={handleSearchChange}
        placeholder="Найти пользователя..."
      />
      <ul>
        {filteredUsers.map(user => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
};

const SearchUsersApp = () => {
  const users = [
    { id: 1, name: 'Alice' },
    { id: 2, name: 'Bob' },
    { id: 3, name: 'Charlie' }
  ];

  return (
    <div>
      <h1>Список пользователей</h1>
      <UserList users={users} />
    </div>
  );
};

export default SearchUsersApp;