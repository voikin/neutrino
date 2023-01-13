from aiogram import types
from aiogram.dispatcher import FSMContext


from loader import dp
from keyBoards import kb_menu
from utils.text import MENU_TEXT


@dp.callback_query_handler(text='/cancel', state='*')
async def cancel_inline(call: types.CallbackQuery, state: FSMContext):
    await state.finish()
    await call.message.answer(text=MENU_TEXT, reply_markup=kb_menu)
    await call.answer()


@dp.message_handler(commands=['cancel'], state='*')
async def cancel(message: types.Message, state: FSMContext):
    await state.finish()
    await message.answer(text=MENU_TEXT, reply_markup=kb_menu)