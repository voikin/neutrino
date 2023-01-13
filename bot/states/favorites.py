from aiogram.dispatcher.filters.state import State, StatesGroup


class FavoritesStates(StatesGroup):
    start = State()
    check = State()