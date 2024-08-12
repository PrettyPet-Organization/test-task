def factorial(x):
    if x == 1:
        return 1
    return x * factorial(x - 1)

def get_binomial_coeff(n, k):
    if n == k or k == 0:
        return 1
    elif n == k - 1 or k == 1:
        return n
    return factorial(n) / (factorial(k) * factorial(n - k))

def get_power_of_sum(a, b, n):
    res = 0
    for k in range(n + 1):
        res += get_binomial_coeff(n, k) * a**(n - k) * b**k
    return res


if __name__ == '__main__':
    print(
        'Это программа для расчета n-степени',
        'суммы двух чисел.'
    )
    a, b = map(int, input('Введите два целых числа через пробел: ').split())
    n = int(input('Введите степень, в которую хотите возвести (a + b): '))
    print(f'({a} + {b})^{n} == {get_power_of_sum(a, b, n)}')
