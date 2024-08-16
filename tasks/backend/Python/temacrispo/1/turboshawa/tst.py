# KISS - OK

def calculate_sum(inter):
    return sum(inter)


# KISS - НЕ ОК

def calculate_sum(inter):
    total = 0
    indx = 0
    while indx < len(inter):
        total += inter[indx]
        indx += 1
    return total
