from math import sqrt


def solving_quadratic_equation(a: int, b: int, c: int) -> None:
    if a == 0:
        print(f"a = 0, not a quadratic equation, but answer {-c / b}")
        return
    
    D = b ** 2 - 4 * a * c

    if D < 0:
        print("Discriminant < 0,\nno roots")
    
    elif D == 0:
        print(f"Discriminant = 0,\n1 roots {-b / 2 * a}")
    
    else:
        print(f"Discriminant > 0,\n2 roots {(-b + sqrt(D))/ 2 * a}, {(-b - sqrt(D))/ 2 * a}")


if __name__ == "__main__":
    print("enter integers")
    a = int(input("a - "))
    b = int(input("b - "))
    c = int(input("c - "))
    solving_quadratic_equation(a, b, c)