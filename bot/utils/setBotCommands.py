from aiogram import types, Dispatcher

async def set_commands(dp: Dispatcher):
    await dp.bot.set_my_commands([
        types.BotCommand('start', 'Запустить бота'),
        types.BotCommand('help', 'Помощь'),
        types.BotCommand('menu', 'Главное меню'),
        types.BotCommand('cancel', 'Отмена'),
    ])
