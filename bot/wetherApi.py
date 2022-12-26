from aiohttp import ClientSession

class ServiceIsUnavailableException(Exception):
    def __str__(self):
        return "Service is unavailable"


class ApiClient:
    @classmethod
    async def connect(cls, url):
        return ApiClient(url)
    
    def __init__(self, url):
        self.url = url
        
    async def getWeather(self, city):
        return f'Погода в городе {city}'
    
    async def getWeatherForecast(self, city, days):
        return f'Погода в городе {city} на {days} дней'