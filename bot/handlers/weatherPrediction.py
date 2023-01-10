from aiogram import types
from aiogram.dispatcher import FSMContext


from loader import dp, api
from states import WeatherPrediction
from utils.text import CHOOSE_CITY, CITY_NOT_FOUND, OR_CANCEL_TEXT, CHOOSE_DAYS
from keyBoards import kb_cancel, kb_menu


day_names = ('Завтра', 'Послезавтра')

@dp.message_handler(text='Прогноз погоды', state='*')
async def weather_prediction_start(message: types.Message):
    await WeatherPrediction.start.set()
    await message.answer(CHOOSE_CITY, reply_markup=kb_cancel)


@dp.message_handler(state=WeatherPrediction.start)
async def weather_prediction_choose_city(message: types.Message, state: FSMContext):
    res = await api.getWeather(message.text)
    if res:
        await message.answer(CHOOSE_DAYS, reply_markup=kb_cancel)
        await state.update_data(city=message.text)
        await WeatherPrediction.city.set()
    else:
        await message.reply(f'{CITY_NOT_FOUND}\n{OR_CANCEL_TEXT}', reply_markup=kb_cancel)


# @dp.message_handler(state=WeatherPrediction.city)
# async def weather_prediction_choose_days(message: types.Message, state: FSMContext):
#     if int(message.text) in [1, 2, 3]:
#         user_data = await state.get_data()
#         res = await api.getWeatherForecast(user_data['city'], message.text)
#         answer = f'{res["nearest_area"]["areaName"]["value"]}:\n'\
#                  f'\tСегодня:\n'\
#                  f'\t\tТемпература: {res["current_condition"]["temp_C"]}\n'\
#                  f'\t\tОщущается как: {res["current_condition"]["feelsLikeC"]}\n'\
#                  f'\t\tМинимальная температура: {res["weather"][0]["mintempC"]}\n'\
#                  f'\t\tМаксимальная температура: {res["weather"][0]["maxtempC"]}'
#         for i in range(1, int(message.text)):
#             answer += f'\n\n\t{day_names[i-1]}:\n' \
#                       f'\t\tМинимальная температура: {res["weather"][i]["mintempC"]}\n' \
#                       f'\t\tСредняя температура: {res["weather"][i]["avgtempC"]}\n' \
#                       f'\t\tМаксимальная температура: {res["weather"][0]["maxtempC"]}'
#
#         await message.answer(answer, reply_markup=kb_menu)
#         await state.finish()
#     else:
#         await message.reply(f'{CITY_NOT_FOUND}\n{OR_CANCEL_TEXT}', reply_markup=kb_cancel)
@dp.message_handler(state=WeatherPrediction.city)
async def weather_prediction_choose_days(message: types.Message, state: FSMContext):
    if int(message.text) in [1, 2, 3]:
        user_data = await state.get_data()
        res = await api.getWeatherForecast(user_data['city'], message.text)
        await message.answer(str(res), reply_markup=kb_menu)
        await state.finish()
    else:
        await message.reply(f'{CITY_NOT_FOUND}\n{OR_CANCEL_TEXT}', reply_markup=kb_cancel)
