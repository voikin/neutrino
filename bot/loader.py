from aiogram import Bot, Dispatcher

from aiogram.contrib.fsm_storage.memory import MemoryStorage

from data import config
from data.config import POSTGRES_URI

from weatherApi import ApiClient

bot = Bot(token=config.TOKEN)
dp = Dispatcher(bot, storage=MemoryStorage())

api = ApiClient(config.URL)

