from aiogram.dispatcher.filters.state import State, StatesGroup


class WeatherToday(StatesGroup):
    start = State()


class WeatherPrediction(StatesGroup):
    start = State()
    city = State()
    days = State()
