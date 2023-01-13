from loader import dp

from aiogram import types

from keyBoards import kb_start

from utils.text import START_TEXT

@dp.message_handler(commands=['start'])
async def start(message: types.Message):
    await message.answer(text=START_TEXT, reply_markup=kb_start)