from aiogram import types
from aiogram.dispatcher import FSMContext
from keyBoards import kb_cancel, kb_menu
from loader import api, dp
from states import WeatherToday
from utils.text import CHOOSE_CITY, CITY_NOT_FOUND, OR_CANCEL_TEXT


@dp.message_handler(text="Погода сейчас", state="*")
async def weather_now_start(message: types.Message):
    await message.answer(text=CHOOSE_CITY, reply_markup=kb_cancel)
    await WeatherToday.start.set()


@dp.message_handler(state=WeatherToday.start)
async def weather_now_city(message: types.Message, state: FSMContext):
    res = await api.getWeather(message.text)
    if res:
        await message.answer(str(res), reply_markup=kb_menu)
        await state.finish()
    else:
        await message.reply(
            f"{CITY_NOT_FOUND}\n{OR_CANCEL_TEXT}", reply_markup=kb_cancel
        )
