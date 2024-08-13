import math


class Point:
    def __init__(self, x: int, y: int):
        if not (isinstance(x, int)) or not (isinstance(y, int)):
            raise ValueError('x and y must be integers')
        self.x = x
        self.y = y

    def distance(self, other):
        if not isinstance(other, Point):
            raise ValueError('other must be a Point')
        return math.sqrt((self.x - other.x) ** 2 + (self.y - other.y) ** 2)

    def manhattan_distance(self, other):
        if not isinstance(other, Point):
            raise ValueError('other must be a Point')
        return abs(self.x - other.x) + abs(self.y - other.y)


def get_coordinate(name):
    while True:
        try:
            coord = int(input(f"{name} = "))
            return coord
        except ValueError:
            print('only integers are accepted, please try again')


if __name__ == '__main__':
    print("Hello! Let's calculate the distance between two points.")
    print("Please, enter the coordinates of the first point:")
    p1 = Point(get_coordinate('x1'), get_coordinate('y1'))
    print("Please, enter the coordinates of the second point:")
    p2 = Point(get_coordinate('x2'), get_coordinate('y2'))

    print(f"The Euclidean distance between the points is {p1.distance(p2)}")
    print(f"The Manhattan distance between the points is {p1.manhattan_distance(p2)}")
