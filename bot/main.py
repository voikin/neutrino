# QqWwEeRrTtYyUuIiOoPpAaSsDdFfGgHhJjKkLlZzXxCcVvBbNnMm

import logging
import os

from dotenv import load_dotenv
from aiogram import Bot, Dispatcher, executor, types

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
    await message.answer(text="введи место где ты хочешь узнать погоду")
    
    
if __name__ == "__main__":
    executor.start_polling(dp, skip_updates=True)
