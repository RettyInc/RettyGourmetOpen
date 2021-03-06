# 問題2

## 問題

Rettyでは, ユーザー情報や店舗情報, 口コミ情報, 画像に関する情報など様々な情報をデータベースで管理していて, テーブルの数やレコードの数は膨大です. そんな複雑なデータベースを効率よく操作できないとサービスを利用するユーザーさんに価値を提供することができません. この問題では効率よくデータベースを操作して回答を導けるように心がけてみましょう. 

次の数字をデータベースから取得し, 順番につなぎ合わせた数字が回答です. 

1. 店舗についたMYBESTの口コミID
2. 店舗についたEXCELLENTの数
3. 店舗についたGOODの数
4. 店舗についたAVERAGEの数

※ 店舗は「中華麺舗 虎」です. 
※ 有効なユーザーの口コミ数を数えてください.  
※ Discordのテキストで解いた方法（SQL文など）を教えて下さい.   

例）
店舗についたMYBESTの口コミID: 10  
店舗についたEXCELLENTの数: 20  
店舗についたGOODの数: 30  
店舗についたAVERAGEの数: 40  
→ 10203040  

## テーブル情報

### Userテーブル

| カラム名 | 説明 |
|:---:|:------|
| id | ユーザーID |
| name | ユーザー名 |
| status | ユーザーの状態 |

user.status = 1: 有効なユーザー  
user.status = 2: 退会済みのユーザー

### Restaurantテーブル

| カラム名 | 説明 |
|:---:|:------|
| id | 店舗ID |
| name | 店舗名 |

### Reportテーブル

| カラム名 | 説明 |
|:---:|:------|
| id | 口コミID |
| rate | 評価 |
| user_id | ユーザーID |
| restaurantID | 店舗ID |

rateの種類

- MYBEST
- EXCELLENT
- GOOD
- AVERAGE
- 評価なし

[回答](https://github.com/RettyInc/RettyGourmetOpen/blob/main/20210730/question2/answer.md)
