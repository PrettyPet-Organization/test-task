import React, { useState } from 'react';

interface IUser {
  id: number;
  name: string;
}

interface IUserList {
  users: IUser[];
}

const useSearchUsers = (initUserList: IUser[]) => {
  const [search, setSearch] = useState('');
  const [filteredUsers, setFilteredUsers] = useState(initUserList);

  const handleSearchChange = e => {
    setSearch(e.target.value);
    setFilteredUsers(
      initUserList.filter(user =>
        user.name.toLowerCase().includes(e.target.value.toLowerCase()),
      ),
    );
  };

  return { search, filteredUsers, handleSearchChange };
};

const UserList: React.FC<IUserList> = ({ users }) => {
  const { search, filteredUsers, handleSearchChange } = useSearchUsers(users);

  return (
    <div>
      <input
        type="text"
        value={search}
        onChange={handleSearchChange}
        placeholder="Search users..."
      />
      <ul>
        {filteredUsers.map(user => (
          <li key={user.id}>{user.name}</li>
        ))}
      </ul>
    </div>
  );
};

const App = () => {
  const users: IUser[] = [
    { id: 1, name: 'Alice' },
    { id: 2, name: 'Bob' },
    { id: 3, name: 'Charlie' },
  ];

  return (
    <div>
      <h1>User List</h1>
      <UserList users={users} />
    </div>
  );
};

export default App;
