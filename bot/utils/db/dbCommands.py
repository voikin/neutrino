import logging

from .dbSettings import User_city


async def create_favorite(user_id, city):
    try:
        await User_city.create(user_id=user_id, city=city)
        return 0
    except Exception as err:
        logging.exception(f'Не удалось добавить Город в отслеживание{err}', exc_info=True)
        return 1


async def get_favorites_by_id(user_id):
    cityes = await User_city.select('city').where(User_city.user_id==user_id).gino.all()
    return [i[0] for i in cityes]
