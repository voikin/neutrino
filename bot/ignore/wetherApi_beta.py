from aiohttp import ClientSession

class ServiceIsUnavailableException(Exception):
    def __str__(self):
        return "Service is unavailable"


class ApiClient:
    @classmethod
    async def connect(cls, url):
        async with ClientSession() as session:
            async with session.get(f'http://{url}/ping') as response:
                if response.status != 200:
                    raise ServiceIsUnavailableException
        
        return ApiClient(url)


    
    def __init__(self, url):
        self.url = f'http://{url}'
        print(url)

    async def getWeather(self, city):
        async with ClientSession() as session:
            async with session.post(f'{self.url}/api/weather-by-city', json={"city": city}) as response:
                print(await response.text())