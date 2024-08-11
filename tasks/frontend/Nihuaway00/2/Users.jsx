import React, {useEffect, useMemo, useState} from 'react';

const UserList = ({ users }) => {
    const [search, setSearch] = useState('');
    // Можно избежать линего состояния
    // const [filteredUsers, setFilteredUsers] = useState(users);


    // useMemo выполнить функцию только при изменении search или user
    const filteredUsers = useMemo(() => users.filter(user =>
        user.name.toLowerCase().includes(search.toLowerCase())
    ), [search, users]);

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

const Users = () => {
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

export default Users;