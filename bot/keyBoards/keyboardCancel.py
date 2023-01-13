from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton


kb_cancel = InlineKeyboardMarkup(row_width=1, inline_keyboard=[
    [
        InlineKeyboardButton(text='Отмена', callback_data='/cancel'),
    ]
], resize_keyboard=True)