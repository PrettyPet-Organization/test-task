import random
from string import ascii_lowercase, ascii_uppercase

chars = ascii_lowercase + ascii_uppercase
random.seed(1)
indx = random.randint(0, len(chars) - 1)
dot = '@gmail.com'


def creat_email(type_len):
    global indx
    domen_name = ''
    while len(domen_name) != type_len:
        domen_name += chars[indx]
        indx = random.randint(0, len(chars) - 1)
    yield domen_name + dot


type_len = int(input('ВВедите количество символов'))

for i in range(5):
    print(next(creat_email(type_len)))
