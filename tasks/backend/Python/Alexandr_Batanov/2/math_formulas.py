def difference_squares(a, b):
    """
    Разность квадратов
    a^2 - b^2 = (a - b)(a + b)
    """

    difference = (a - b) * (a + b)
    return difference

def square_sum_square_diff(a, b):
    """
    Квадрат суммы и квадрат разности
    (a + b)^2 = a^2 + 2ab + b^2
    (a - b)^2 = a^2 - 2ab + b^2
    """

    sqr_sum = a**2 + 2 * a * b + b**2
    sqr_diff = a**2 - 2 * a * b + b**2
    return sqr_sum, sqr_diff

def cudes_sum_cubes_diff(a, b):
    """
    Сумма и разность кубов
    a^3 + b^3 = (a + b)(a^2 - ab + b^2)
    a^3 - b^3 = (a - b)(a^2 + ab + b^2)
    """

    cubes_sum = (a + b) * (a**2 - a * b + b**2)
    cubes_diff = (a - b) * (a**2 + a * b + b**2)
    return cubes_sum, cubes_diff

def cub_sum_cub_diff(a, b):
    """
    Куб суммы и разности
    (a + b)^3 = a^3 + 3a^2b + 3ab^2 + b^3
    (a - b)^3 = a^3 - 3a^2b + 3ab^2 - b^3
    """

    cube_sum = a**3 + 3 * a**2 * b + 3 * a * b**2 + b**3
    cube_diff = a**3 - 3 * a**2 * b + 3 * a * b**2 - b**3
    return cube_sum, cube_diff
