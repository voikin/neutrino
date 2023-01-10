async def on_startup(dp):

    from utils.setBotCommands import set_commands

    await set_commands(dp)

    from utils.notifyAdmins import notify_admins

    await notify_admins(dp)

    from utils.db import dbSettings

    await dbSettings.create_db()

    print("Бот запущен")


if __name__ == "__main__":
    from aiogram import executor
    from handlers import dp

    executor.start_polling(dp, skip_updates=True, on_startup=on_startup)
