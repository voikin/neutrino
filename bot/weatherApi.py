# import requests
from aiohttp import ClientSession, client_exceptions


class ServiceIsUnavailableException(Exception):
    def __str__(self):
        return "Service is unavailable"


class ApiAnswerForOneDay:
    def __init__(self, res):
        if isinstance(res["weather"], list):
            tres = res["weather"][0]
        else:
            tres = res["weather"]
        self.city = res["nearest_area"]["areaName"]["value"]
        self.T = res["current_condition"]["temp_C"]
        self.T_feels_like = res["current_condition"]["feelsLikeC"]
        self.min_t = tres["mintempC"]
        self.max_t = tres["maxtempC"]

    def __str__(self):
        return f'{self.city}:\n'\
               f'\tТемпература: {self.T}\n'\
               f'\tОщущается как: {self.T_feels_like}\n'\
               f'\tМинимальная температура: {self.min_t}\n'\
               f'\tМаксимальная температура: {self.max_t}'


class ApiAnswerShort:
    def __init__(self, res):
        self.min = res["mintempC"]
        self.avg = res["avgtempC"]
        self.max = res["maxtempC"]
        self.date = res["date"]

    def __str__(self):
        return f'\t\tМинимальная температура: {self.min}\n' \
               f'\t\tСредняя температура: {self.avg}\n' \
               f'\t\tМаксимальная температура: {self.max}\n'


class ApiAnswer:
    def __init__(self, res):
        self.days = len(res["weather"])
        self.today = ApiAnswerForOneDay(res)
        self.other = [ApiAnswerShort(i) for i in res["weather"][1::]]

    def __str__(self):
        day_names = ('Завтра', 'Послезавтра')
        answer = str(self.today)
        for i in range(self.days - 1):
            answer += f'\n\n\t{day_names[i]}:\n' \
                      f'{self.other[i]}'
        return answer


class ApiClient:
    def __init__(self, url):
        self.url = url

    async def getWeather(self, city):
        async with ClientSession() as session:
            try:
                async with session.get(
                    f"http://{self.url}/api/weather-by-city?city={city}"
                ) as response:
                    if response.status != 200:
                        return {}
                    return ApiAnswerForOneDay(await response.json())
            except client_exceptions.ClientConnectorError:
                return {}

    async def getWeatherForecast(self, city, days):
        async with ClientSession() as session:
            async with session.get(f"http://{self.url}/api/forecast-by-city?city={city}&days={days}") as response:
                if response.status != 200:
                    return {}
                return ApiAnswer(await response.json())
