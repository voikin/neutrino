import os
from dotenv import load_dotenv

load_dotenv()
TOKEN = os.getenv("BOT_API_KEY")
URL = os.getenv("WEATHER_IP")

admins = [430625699, ]
# 1391459225

DB_IP = os.getenv("DB_IP")
DB_USER = os.getenv("DB_USER")
DB_PASS = os.getenv("DB_PASS")
DB_NAME = os.getenv("DB_NAME")

POSTGRES_URI = f'postgresql://{DB_USER}:{DB_PASS}@{DB_IP}/{DB_NAME}'
