from aiogram.types import ReplyKeyboardMarkup, KeyboardButton

kb_menu = ReplyKeyboardMarkup([
    [
        KeyboardButton(text='Погода сейчас'),
        KeyboardButton(text='Прогноз погоды'),
    ],
    [
        KeyboardButton(text='Подписаться на город'),
        KeyboardButton(text='Добавить город в избранное')
    ]
], resize_keyboard=True, one_time_keyboard=True)