# $ python --version
# Python 3.7.4
# $ pip install geopy

import pandas as pd
from geopy.distance import geodesic

df = pd.read_csv("retty_gourmet_open_q1.csv")
df.head()

station = (35.6697, 139.76714)

ans = []
dis = []

for row in df.itertuples():
    ans.append(row[1])
    restaurant = (row[2], row[3])
    t_dis = geodesic(station, restaurant).km 
    dis.append(t_dis)

df_dis = pd.DataFrame(index=[], columns=['answer', 'distance'])
df_dis['answer'] = ans
df_dis['distance'] = dis

print(df_dis.sort_values('distance').iloc[[0]])