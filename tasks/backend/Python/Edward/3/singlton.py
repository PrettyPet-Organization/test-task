#простая реализация singlton
class DataBase():
    __instans = None

    def __new__(cls, *args, **kwargs):
        print('запуск __new__')
        if cls.__instans is None:
            cls.__instans = super().__new__(cls)
        return cls.__instans

    def __del__(self):
        print('удаление')
        DataBase.__instans = None

    def __init__(self, name, psw, port):
        self.name = name
        self.psw = psw
        self.port = port

    def connect(self):
        print(f'подключился {self.name}, {self.psw}, {self.port}')

    def read(self):
        return ' Данные из бд'

    def add(self, mes):
        print(f'В базу данных добавлины данные: {mes}')

    def close(self):
        print('бд закрыта')


bd1 = DataBase('noname', '1234', '881')
bd2 = DataBase('ed', '12kk34', '999')

print(bd1.name)
print(id(bd1) == id(bd2) )
