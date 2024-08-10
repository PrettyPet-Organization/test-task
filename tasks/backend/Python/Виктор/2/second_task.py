# Терпеть не могу математику, поэтому вот так
def check_chase(distance: float, speed_of_voenkom: float, speed_of_recruit: float):
    if speed_of_voenkom <= speed_of_recruit:
        print("Военкомат не догоняет вас.")
    else:
        time_to_catch = distance / (speed_of_voenkom - speed_of_recruit)
        print(f"Военкомат догоняет вас. Время до повестки: {time_to_catch:.2f} часов.")


# Примеры использования
distance = float(input("Введите начальную дистанцию (в километрах): "))
speed_of_voenkom = float(input("Введите скорость военкомата (в километрах в час): "))
speed_of_recruit = float(input("Введите скорость бегуна (в километрах в час): "))

check_chase(distance, speed_of_voenkom, speed_of_recruit)

# Все персонажи вымышлены, а совпадения случайны
# Да ладно, просто ради шутки сделал, чтобы было весело мой недокод читать
