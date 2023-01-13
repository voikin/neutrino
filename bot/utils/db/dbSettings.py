import logging
from time import sleep

from gino import Gino

from data.config import POSTGRES_URI

db = Gino()


class User_city(db.Model):
    __tablename__ = 'user_city'

    id = db.Column(db.Integer, primary_key=True)
    user_id = db.Column(db.Integer, autoincrement=False)
    city = db.Column(db.String)
    _uix = db.UniqueConstraint("user_id", "city", name="uix_1")


async def create_db():
    while True:
        try:
            await db.set_bind(POSTGRES_URI)
            break
        except OSError as err:
            logging.fatal(f'Не удалось подключиться к БД\n{err}')
            sleep(5)
    await db.gino.create_all()
