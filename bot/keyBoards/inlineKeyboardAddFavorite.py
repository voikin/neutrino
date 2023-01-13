from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton

kb_add_favorite = InlineKeyboardMarkup(1,[
    [
        InlineKeyboardButton('Добавить место в Избранные', callback_data='add_to_favorite')
    ]
])