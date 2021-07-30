# $ python --version
# Python 3.7.4
# $ pip install geopy

import pandas as pd
from geopy.distance import geodesic

df = pd.read_csv("retty_gourmet_open_q1.csv")
df.head()

station = (35.6546783, 139.7369874)
 
answer = []
distance = []
category = []
hour_from = []
hour_to = []

for row in df.itertuples():
    answer.append(row[1])
    category.append(row[4])
    hour_from.append(int(row[5].replace(":","")))
    hour_to.append(int(row[6].replace(":","")))

    latlng = (row[2], row[3])
    t_dis = geodesic(station, latlng).km
    distance.append(t_dis)

df_dis = pd.DataFrame(index=[], columns=['answer', 'distance', 'category', 'hour_from', 'hour_to'])
df_dis['answer'] = answer
df_dis['distance'] = distance
df_dis['category'] = category
df_dis['hour_from'] = hour_from
df_dis['hour_to'] = hour_to

df_dis.dtypes
df_dis.head()

print(df_dis.query('category=="焼肉" and distance < 1.5 and hour_from <= 170000 and 170000 < hour_to'))