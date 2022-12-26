from aiogram import Bot, Dispatcher, types

from data import config

bot = Bot(token=config.TOKEN)
dp = Dispatcher(bot)