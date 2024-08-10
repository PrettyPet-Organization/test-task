from enum import Enum
import random


class Move(Enum):
    ROCK = 1
    PAPER = 2
    SCISSORS = 3

    @property
    def translation(self):
        return {
            Move.ROCK: 'камень',
            Move.PAPER: 'бумага',
            Move.SCISSORS: 'ножницы',
        }[self]


def get_user_move():
    while True:
        try:
            move = int(input("Выбери движение (1 - камень, 2 - бумага, 3 - ножницы): "))
            if move in range(1, 4):
                return Move(move)
            else:
                raise ValueError
        except ValueError:
            print('Неверный ввод, попробуй снова')


def get_computer_move():
    return Move(random.randint(1, 3))


def get_winner(user_move, computer_move):
    match user_move, computer_move:
        case Move.ROCK, Move.SCISSORS | Move.PAPER, Move.ROCK | Move.SCISSORS, Move.PAPER:
            return 'Ты победил'
        case _ if user_move == computer_move:
            return 'Ничья'
        case _:
            return 'Ты проиграл'


if __name__ == '__main__':
    print('Давай сыграем!')
    user_move = get_user_move()
    computer_move = get_computer_move()
    print(f'Ход ИИ: {computer_move.translation}')
    print(get_winner(user_move, computer_move))
