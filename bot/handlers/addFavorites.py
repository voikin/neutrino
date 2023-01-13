from aiogram import types
from aiogram.dispatcher import FSMContext

from keyBoards import kb_cancel, kb_confirm, kb_menu
from states import FavoritesStates
from loader import dp, api
from utils.text import FAVORITES
from utils.db import create_favorite


@dp.message_handler(text="Добавить город в избранное", state='*')
async def start_favorites(message: types.Message):
    await FavoritesStates.start.set()
    await message.answer(FAVORITES, reply_markup=kb_cancel)


@dp.message_handler(state=FavoritesStates.start)
async def check_favorites(message: types.Message, state: FSMContext):
    res = await api.getWeather(message.text)
    if res:
        await message.answer(f'Подтвердите место: {res.city}', reply_markup=kb_confirm)
        await state.update_data(city=message.text)
        await FavoritesStates.check.set()
    else:
        await message.answer(f'Не удалось найти такое место\n{FAVORITES}', reply_markup=kb_cancel)


@dp.callback_query_handler(text='yes', state=FavoritesStates.check)
async def add_favorite_yes(call: types.CallbackQuery, state: FSMContext):
    state_data = await state.get_data()
    await state.finish()
    match await create_favorite(call.from_user, state_data['city']):
        case 0:
            await call.message.answer(f'Место {state_data["city"].capitalize()} успешно добавлено.', reply_markup=kb_menu)
            await call.answer()
        case 1:
            await call.answer('Уже в избранном')
        case -1:
            await call.answer(f'Произошла ошибка', show_alert=True)
    await state.finish()


@dp.callback_query_handler(text='no', state=FavoritesStates.check)
async def add_favorite_no(call: types.CallbackQuery, state: FSMContext):
    await call.message.answer(FAVORITES, reply_markup=kb_cancel)
    await FavoritesStates.start.set()


@dp.callback_query_handler(text='add_to_favorite')
async def inline_add_favorite(call: types.CallbackQuery):
    city = call.message.text.split(':', 1)[0]
    match await create_favorite(call.from_user, city):
        case 0:
            await call.message.edit_reply_markup(None)
            await call.answer('Добавлено')
        case 1:
            await call.answer('Уже в избранном')
        case -1:
            await call.answer('Произошла ошибка', show_alert=True)