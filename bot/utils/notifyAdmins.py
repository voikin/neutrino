import logging

from aiogram import Dispatcher
from data.config import admins

async def notify_admins(dp: Dispatcher):
    for admin in admins:
        try:
            await dp.bot.send_message(chat_id=admin, text='Бот запущен')
        except Exception as e:
            logging.exception(e)