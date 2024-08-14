class ComplexNumber:
    """Класс для работы с комплексными числами."""

    def __init__(self, real: float, imag: float) -> None:
        """
        Инициализация комплексного числа.

        :param real: Действительная часть.
        :param imag: Мнимая часть.
        """
        self.real = real
        self.imag = imag

    def __add__(self, other: 'ComplexNumber') -> 'ComplexNumber':
        """
        Складывает два комплексных числа.

        :param other: Второе комплексное число.
        :return: Результат сложения в виде нового комплексного числа.
        """
        return ComplexNumber(self.real + other.real, self.imag + other.imag)

    def __sub__(self, other: 'ComplexNumber') -> 'ComplexNumber':
        """
        Вычитает одно комплексное число из другого.

        :param other: Второе комплексное число.
        :return: Результат вычитания в виде нового комплексного числа.
        """
        return ComplexNumber(self.real - other.real, self.imag - other.imag)

    def __mul__(self, other: 'ComplexNumber') -> 'ComplexNumber':
        """
        Умножает два комплексных числа.

        :param other: Второе комплексное число.
        :return: Результат умножения в виде нового комплексного числа.
        """
        real_part = self.real * other.real - self.imag * other.imag
        imag_part = self.real * other.imag + self.imag * other.real
        return ComplexNumber(real_part, imag_part)

    def __truediv__(self, other: 'ComplexNumber') -> 'ComplexNumber':
        """
        Делит одно комплексное число на другое.

        :param other: Второе комплексное число.
        :return: Результат деления в виде нового комплексного числа.
        :raises ZeroDivisionError: Если деление на ноль.
        """
        if other.real == 0 and other.imag == 0:
            raise ZeroDivisionError('Деление на ноль невозможно')

        denominator = other.real ** 2 + other.imag ** 2
        real_part = (self.real * other.real + self.imag * other.imag) / denominator
        imag_part = (self.imag * other.real - self.real * other.imag) / denominator
        return ComplexNumber(real_part, imag_part)

    def __str__(self) -> str:
        """
        Возвращает строковое представление комплексного числа.

        :return: Комплексное число в строковом формате.
        """
        return f'{self.real} + {self.imag}i'
