# RettyGourmetOpen 2021/07/30

## データベース操作

### データベースをインポート

```
mysql -h localhost -u root < question2.sql
```

### データベースに接続

```
mysql -h localhost -u root
```

### データベースを確認

```
show databases;
```

### 操作するデータベースに移動

```
use rgo20210730
```

## テーブル操作

### テーブルを確認

```
show tables;
```

### テーブルの中身を確認

```
SELECT * FROM user LIMIT 20;
SELECT * FROM restaurant LIMIT 20;
SELECT * FROM report LIMIT 20;
```

## 回答例

### RestaurantテーブルとReportテーブルを結合

```
# Restaurantテーブルから「中華麺舗 虎」を取得
SELECT * FROM restaurant WHERE name = "中華麺舗 虎"

# RestaurantテーブルとReportテーブルを結合
SELECT *
FROM report LEFT JOIN restaurant ON report.restaurant_id = restaurant.id
LIMIT 20;

# 結合したテーブルから店舗名が「中華麺舗 虎」の口コミを取得
SELECT *
FROM report LEFT JOIN restaurant ON report.restaurant_id = restaurant.id
WHERE restaurant.name = "中華麺舗 虎"
LIMIT 20;
```

### UserテーブルとReportテーブルを結合

```
# Userテーブルから有効なユーザーを取得
SELECT * FROM user WHERE status = 1;

# UserテーブルとReportテーブルを結合
SELECT *
FROM report LEFT JOIN user ON report.user_id = user.id
LIMIT 20;

# 結合したテーブルから有効なユーザーの口コミを取得
SELECT *
FROM report LEFT JOIN user ON report.user_id = user.id
WHERE user.status = 1
LIMIT 20;
```

### UserテーブルとRestaurantテーブル, Reportテーブルを結合

```
# UserテーブルとRestaurantテーブル, Reportテーブルを結合
SELECT *
FROM
	report
	LEFT JOIN restaurant ON report.restaurant_id = restaurant.id
	LEFT JOIN user ON report.user_id = user.id
LIMIT 20;

# 結合したテーブルから「中華麺舗 虎」の有効なユーザーの口コミを取得
SELECT *
FROM
	report
	LEFT JOIN restaurant ON report.restaurant_id = restaurant.id
	LEFT JOIN user ON report.user_id = user.id
WHERE
	restaurant.name = "中華麺舗 虎"
	AND	user.status = 1
LIMIT 20;

# 結合したテーブルから「中華麺舗 虎」の有効なユーザーのMYBESTの評価がついた口コミを取得
SELECT *
FROM
	report
	LEFT JOIN restaurant ON report.restaurant_id = restaurant.id
	LEFT JOIN user ON report.user_id = user.id
WHERE
	restaurant.name = "中華麺舗 虎"
	AND	user.status = 1
	AND	report.rate = "MYBEST"
LIMIT 20;

# 結合したテーブルから「中華麺舗 虎」の有効なユーザーのEXCELLENT, GOOD, AVERAGEの評価がついた口コミの数を取得
SELECT
	COUNT(report.rate = "EXCELLENT" OR NULL) AS EXCELLENT_COUNT,
	COUNT(report.rate = "GOOD" OR NULL) AS GOOD_COUNT,
	COUNT(report.rate = "AVERAGE" OR NULL) AS AVERAGE_COUNT
FROM
	report
	LEFT JOIN restaurant ON report.restaurant_id = restaurant.id
	LEFT JOIN user ON report.user_id = user.id
WHERE
	restaurant.name = "中華麺舗 虎"
	AND	user.status = 1;
```
