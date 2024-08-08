from dataclasses import dataclass, asdict


@dataclass(frozen=True)
class Duck:
    name: str = 'Dennis'
    color: str = 'Yellow'
    weight: float = 1.0


def main():
    first_duck = Duck()
    second_duck = Duck('Dennis', 'Yellow', 1)

    # Уже можно сравнивать
    assert first_duck == second_duck
    assert first_duck is not second_duck

    try:
        first_duck.weight = 1.5
    except Exception:
        """По настоящему неизменямый объект!"""
        print(first_duck.weight == 1.0)

    # Выглядит неплохо и без лишнего кода
    print(first_duck, second_duck, sep='\n')

    # Можно легко преобразовать в словарь
    print(asdict(first_duck), asdict(second_duck), sep='\n')
    
    """
    В общем датаклассы - это круто.
    """


if __name__ == '__main__':
    main()