from aiogram import types

from loader import dp

from keyBoards import kb_menu

from utils.text import MENU_TEXT

@dp.message_handler(commands=['menu'])
async def menu(message: types.Message):
    await message.answer(text=MENU_TEXT, reply_markup=kb_menu)