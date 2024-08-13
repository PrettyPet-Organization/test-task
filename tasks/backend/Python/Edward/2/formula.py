#Это программа для вычисления объемов разных фигур

from math import pi


def cube() -> float:
    a = float(input('enter the side of the cube'))
    return a ** 3


def parallelepiped() -> float:
    print('enter the sides of the parallelepiped')
    length = float(input('length ?'))
    width = float(input('width ?'))
    height = float(input('height ?'))

    return length * width * height


def cylinder() -> float:
    radius = float(input('radius ?'))
    height = float(input('height ?'))

    return round((pi * radius ** 2) * height, 2)

def main():
    math_formul = {
        '1': cube,
        '2': parallelepiped,
        '3': cylinder
    }

    choic = input('volume, what shape you want to find? \n1. cube\n2. parallelepiped\n3. cylinder\n')
    if choic in math_formul.keys():
        result = math_formul[choic]()
        print(f'The volume is {result}')
    else:
        print('Invalid formula number')


if __name__ == "__main__":
    main()
