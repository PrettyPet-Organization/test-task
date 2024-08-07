

class Product:
    """Класс Product для отправки данных

        self.name = name
        :param self.name: name
        :type self.name: str
        :return: self.name
        :rtype: str

        self.price = price
        :param self.price: price
        :type self.price: float
        :return: self.price
        :rtype: float

        self.count = count
        :param self.count: count
        :type self.count: int
        :return: self.count
        :rtype: int
    """

    def __init__(self, name, price, count):
        self.name = name
        self.price = price
        self.count = count

    def to_dict(self):
        return {
            'name': self.name,
            'price': self.price,
            'count': self.count
        }
