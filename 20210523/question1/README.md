# 問題1
## ある駅 (緯度経度座標) から最も近いお店を探す

- 1000店舗の緯度経度情報を持っています。(CSVファイル)
- 1000店舗のうち、ある駅 (緯度 : 35.6697 / 経度 : 139.76714) から最も近いお店を探してください。
- そのお店の【answer (回答項目)】が、問題1の回答です。

※ 緯度経度の座標はすべて世界測地系の座標

CSVファイル (1000店舗の緯度経度情報) は retty_gourmet_open_q1.csv

| answer (回答候補) | latitude (緯度) | longitude (経度) | 備考 |
| --- | --- | --- | --- |
| PRE13/ARE8/SUB802 | 35.6625840 | 139.6973305 | 店舗A の緯度経度情報 |
| PRE13/ARE21/SUB2104 | 35.7799264 | 139.7524000 | 店舗B の緯度経度情報 |
| ... | ... | ... | ... |
| PRE14/ARE31/SUB3103 | 35.5614942 | 139.5052934 | 店舗C の緯度経度情報 | 

## 回答・解説
緯度経度とは、場所を数値的に表現する方法の一つ。

日本測地系 : 日本独自のもの。2002年に世界測地系へ移行している
世界測地系 : 国際的に定められた基準となる測地系が世界測地系

ご存知の通り、地球は平面ではなく球体 (正確には楕円体) になるので、球面を気にして計算する必要がある。

今回の回答例では、Python の GeoPy とよばれるジオコーディングライブラリを利用して実装。
GeoPy では、カーニー法とよばれる、測地線距離を求める新しいアルゴリズムを採用。
https://geopy.readthedocs.io/en/stable/#module-geopy.distance

## その他・備考
今回の問題に限っていくと、平面と捉えても正しい回答が導けます。
三平方の定理を用いて、2点の距離を求めるで回答しても問題ありません。
