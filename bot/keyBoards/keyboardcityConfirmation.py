from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton

kb_confirm = InlineKeyboardMarkup(row_width=2, inline_keyboard=[
    [
        InlineKeyboardButton(text='Да', callback_data='yes'),
        InlineKeyboardButton(text='Нет', callback_data='no')
    ]
])
