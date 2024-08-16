def is_right_triangle(a, b, c):
    sides = sorted([a, b, c])
    a, b, c = sides
    return a ** 2 + b ** 2 == c ** 2


a, b, c = (3, 4, 5)
if is_right_triangle(a, b, c):
    print(f"Треугольник со сторонами {a}, {b} и {c} является прямоугольным.")
else:
    print(f"Треугольник со сторонами {a}, {b} и {c} не является прямоугольным.")
