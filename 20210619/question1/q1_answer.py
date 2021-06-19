# $ python --version
# Python 3.7.4
# $ pip install geopy

import pandas as pd
from geopy.distance import geodesic

df = pd.read_csv("retty_gourmet_open_q1.csv")
df.head()

station = (35.6721903, 139.7363287)

ans = []
dis = []
cat = []

for row in df.itertuples():
    ans.append(row[1])
    restaurant = (row[2], row[3])
    cat.append(row[4])
    t_dis = geodesic(station, restaurant).km
    dis.append(t_dis)

df_dis = pd.DataFrame(index=[], columns=['answer', 'distance'])
df_dis['answer'] = ans
df_dis['distance_km'] = dis
df_dis['category'] = cat

# print(df_dis.sort_values('distance').iloc[[0]])
print(df_dis.sort_values('distance_km'))