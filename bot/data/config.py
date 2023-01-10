import os
from dotenv import load_dotenv

load_dotenv()
TOKEN = os.getenv("BOT_API_KEY")
URL =  os.getenv("WEATHER_IP")

admins = [430625699, ]
# 1391459225