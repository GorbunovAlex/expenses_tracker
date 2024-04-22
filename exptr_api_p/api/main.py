from fastapi import FastAPI
import psycopg2
import json

from api.entities.data import CrimeData

app = FastAPI()

conn_string = "host='localhost' dbname='crime_data' user='postgres' password='postgres'"
conn = psycopg2.connect(conn_string)
cursor = conn.cursor()

@app.get("/data")
def root():
  query = cursor.execute("SELECT * FROM crimes;")
  data = cursor.fetchall()
  records = []
  for row in data:
    item = CrimeData(row[0], row[1], row[2], row[3], row[4], row[5], row[6])
    records.append({
      "id": item.id,
      "description": item.description,
      "weapon_description": item.weapon_description,
      "area_name": item.area_name,
      "latitude": item.latitude,
      "longitude": item.longitude,
      "date": item.date
    })
  return { "count": len(records), "result": records }