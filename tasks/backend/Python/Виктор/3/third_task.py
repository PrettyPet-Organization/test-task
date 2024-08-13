class Member:
    """
    Класс для представления участника нашего проекта

    Атрибуты:
    name (str): Имя участника
    age (int): Возраст участника
    position (str): Позиция участника
    """

    def __init__(self, name, age, position):
        """
        Инициализируем атрибуты участника
        """
        self.name = name
        self.age = age
        self.position = position

    def __str__(self):
        """
        Вывод строкой
        """
        return f"Имя: {self.name}, Возраст: {self.age}, Должность: {self.position}"


class PettyPet:
    """
    Класс нашего авантюрного проекта

    Атрибуты:
    members (list): Список участников
    """

    def __init__(self):
        self.members = []

    def add_member(self, member):
        """
        Добавляет участника в список

        Параметры:
         member (Member): Участник для добавления
        """
        if not isinstance(member, Member):
            raise ValueError("Можно добавить только объекты класса Member.")
        self.members.append(member)

    def remove_member(self, name):
        """
        Удаляем участника по имени.

        Параметры:
        name (str): Имя участника для удаления
        """
        self.members = [p for p in self.members if p.name != name]

    def list_members(self):
        """
        Возвращает список всех участников

        list: Список участников
        """
        return [str(p) for p in self.members]


def main():
    petty_pet = PettyPet()

    # Создаем участников
    viktor = Member("Витёк", 25, "Backend")
    alex = Member("Алекс", 47, "Frontend")

    # Добавляем участников в проект
    petty_pet.add_member(viktor)
    petty_pet.add_member(alex)
    print("Участники после добавления:")
    for member in petty_pet.list_members():
        print(member)

    # Удаляем участника из проекта
    petty_pet.remove_member("Витёк")

    print("\nУчастники после удаления:")
    for member in petty_pet.list_members():
        print(member)


if __name__ == "__main__":
    main()
