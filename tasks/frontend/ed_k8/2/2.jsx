// В представленном в задание коде есть ошибка в логике работы компонента `UserList`. Проблема заключается в том, что при каждом изменении значения в поле поиска `search`, фильтрация выполняется по исходному списку пользователей `users`, который не изменяется при вводе текста. Поэтому фильтрация работает некорректно, особенно если нужно, чтобы список пользователей изменялся динамически.
// Вот исправленный код:

import React, { useState, useEffect } from "react";

const UserList = ({ users }) => {
	const [search, setSearch] = useState("");
	const [filteredUsers, setFilteredUsers] = useState(users);

	useEffect(() => {
		// Фильтрация пользователей при изменении search или users
		setFilteredUsers(
			users.filter((user) =>
				user.name.toLowerCase().includes(search.toLowerCase())
			)
		);
	}, [search, users]); // Эффект срабатывает при изменении search и users

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
