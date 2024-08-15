def geometric_series_sum(a1, q, n):
    if q == 1:
        return a1 * n
    else:
        return a1 * (q**n - 1) / (q - 1)


# Пример: вычисление суммы первых 5 членов прогрессии с a1 = 2, q = 3, n = 5
a1 = 2
q = 3
n = 5

sum_n = geometric_series_sum(a1, q, n)

print(f'Сумма первых {n} членов геометрической прогрессии: {sum_n}')
