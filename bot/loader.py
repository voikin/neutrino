from aiogram import Bot, Dispatcher, types

from aiogram.contrib.fsm_storage.memory import MemoryStorage

from data import config

from weatherApi import ApiClient

bot = Bot(token=config.TOKEN)
dp = Dispatcher(bot, storage=MemoryStorage())

api = ApiClient(config.URL)
