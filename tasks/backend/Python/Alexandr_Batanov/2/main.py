from math_formulas import (
    difference_squares, cub_sum_cub_diff,
    cudes_sum_cubes_diff, square_sum_square_diff
)


choice_functions = {
    '1': difference_squares,
    '2': square_sum_square_diff,
    '3': cudes_sum_cubes_diff,
    '4': cub_sum_cub_diff,
    '5': 'close'
}

def display_menu():
    menu = """
    Выберите формулу для решения:
    1. Разность квадратов
    2. Квадрат суммы и квадрат разности
    3. Сумма и разность кубов
    4. Куб суммы и разности
    5. Выход
    """
    print(menu)

def input_data():
    a = float(input("Введите a: "))
    b = float(input("Введите b: "))
    return a, b

def call_function(func, args):
    print(func.__doc__)
    result = func(*args)
    print(f"Результат вычисления: {result}")

def main():
    display_menu()

    while True:
        choice = input("Введите номер формулы: ")
        func = choice_functions.get(choice)

        if not func:
            print('Не верный выбор')
            continue
        if func == 'close':
            break

        try:
            data = input_data()
            call_function(func, data)
        except:
            print('Не корректные данные')


if __name__ == "__main__":
    main()
