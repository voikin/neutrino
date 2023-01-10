from aiogram import Bot, Dispatcher, types

from aiogram.contrib.fsm_storage.memory import MemoryStorage
from gino import Gino

from data import config
from data.config import POSTGRES_URI

from weatherApi import ApiClient

bot = Bot(token=config.TOKEN)
dp = Dispatcher(bot, storage=MemoryStorage())

api = ApiClient(config.URL)

