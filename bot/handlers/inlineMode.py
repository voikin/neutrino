from uuid import uuid4

from aiogram import types
from loader import api, dp
from utils.db.dbCommands import get_favorites_by_id


@dp.inline_handler()
async def inline_mode(query: types.InlineQuery):
    if query.query:
        text = await api.getWeatherForecast(query.query, 3)
        if not text:
            articles = [
                types.InlineQueryResultArticle(
                    id=uuid4().hex,
                    title="Такой город не найден",
                    input_message_content=types.InputMessageContent(
                        message_text=f"{query.from_user.full_name} хотел отправить погоду в городе {query.query},"
                                     f" но я не могу найти этот город. "
                    ),
                )
            ]
        else:
            articles = [
                types.InlineQueryResultArticle(
                    id=uuid4().hex,
                    title="Погода на сегодня",
                    description=f'Погода в {text.today.city}',
                    input_message_content=types.InputMessageContent(
                        message_text=str(text.today)
                    ),
                )
            ]
            articles.extend([types.InlineQueryResultArticle(
                    id=uuid4().hex,
                    title=f"Погода на {i.date}",
                    description=f'Погода в {text.today.city}',
                    input_message_content=types.InputMessageContent(
                        message_text=f'{text.today.city}:\n\t{i.date}\n{i}'
                    ),
                ) for i in text.other])
    else:
        articles = []
        for city in await get_favorites_by_id(query.from_user):
            text = await api.getWeather(city)
            articles.append(types.InlineQueryResultArticle(
                    id=uuid4().hex,
                    title="Погода на сегодня",
                    description=f'Погода в {text.city}',
                    input_message_content=types.InputMessageContent(
                        message_text=str(text)
                    ),
                ))
    await query.answer(articles, cache_time=1, is_personal=True)