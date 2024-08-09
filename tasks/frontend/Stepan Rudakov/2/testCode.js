// import React, { useState } from 'react';

// const UserList = ({ users }) => {
//   const [search, setSearch] = useState('');
//   const [filteredUsers, setFilteredUsers] = useState(users);

//   const handleSearchChange = (e) => {
//     setSearch(e.target.value);
//     setFilteredUsers(users.filter(user => 
//       user.name.toLowerCase().includes(e.target.value.toLowerCase())
//     ));
//   };

//   return (
//     <div>
//       <input
//         type="text"
//         value={search}
//         onChange={handleSearchChange}
//         placeholder="Search users..."
//       />
//       <ul>
//         {filteredUsers.map(user => (
//           <li key={user.id}>{user.name}</li>
//         ))}
//       </ul>
//     </div>
//   );
// };

// const App = () => {
//   const users = [
//     { id: 1, name: 'Alice' },
//     { id: 2, name: 'Bob' },
//     { id: 3, name: 'Charlie' }
//   ];

//   return (
//     <div>
//       <h1>User List</h1>
//       <UserList users={users} />
//     </div>
//   );
// };

// export default App;


import React, { useState, useMemo } from 'react';

const UserList = ({ users }) => {
  const [search, setSearch] = useState('');

  const filteredUsers = useMemo(() => {
    return users.filter(user => 
      user.name.toLowerCase().includes(search.toLowerCase())
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
  const users = [
    { id: 1, name: 'Alice' },
    { id: 2, name: 'Bob' },
    { id: 3, name: 'Charlie' }
  ];

  return (
    <div>
      <h1>User List</h1>
      <UserList users={users} />
    </div>
  );
};

export default App;
