# QqWwEeRrTtYyUuIiOoPpAaSsDdFfGgHhJjKkLlZzXxCcVvBbNnMm

import logging
import os

from dotenv import load_dotenv
from aiogram import Bot, Dispatcher, executor, types
import requests
from wetherApi import ApiClient

logging.basicConfig(level=logging.INFO)

load_dotenv()

API_KEY = os.getenv("BOT_API_KEY")

bot = Bot(token=API_KEY)
dp = Dispatcher(bot)

@dp.message_handler(commands=["start"])
async def send_welcome(message: types.Message):
    await message.answer(text="hello")
    
@dp.message_handler(commands=["weather"])
async def get_weather(message: types.Message):
    # if len(message.get_args().split) != 1:
        # await message.answer(text="
    # body = {
    #     "city": message.get_args
    # }
    # print(requests.post(url="http://37.140.199.169:8080/api/weather-by-city", json=body) )
    print(message.get_args())
    ac = await ApiClient.connect('37.140.199.169:8080')
    await ac.getWeather('moscow')
    await message.answer(text="введи место где ты хочешь узнать погоду")
    
    
if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)
