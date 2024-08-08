side = int(input("Введите длину катета: "))
c = int(input("Введите длину гипотенузы: "))
other_side = (c ** 2 - side ** 2) ** 0.5

print(f"Длина другого катета примерно равна {round(other_side, 2)}")
