# 問題2
## いくつかのヒントから特定のお店を見つけたい！
5000件のお店データを返す API を用意しました。
curl と jq を使用して次の条件に当てはまるお店 ID (restaurant_id) を表示させてください。

### 条件
- area.area_id が 14
- restaurant_name に「珉」が含まれている
- restaurant_name に「店」が含まれていない
- restaurant_name にアルファベットが含まれていない

### API について

#### ベース URL
```
https://p3hpemjpit.ap-northeast-1.awsapprunner.com
````

#### 認証方法
BASIC 認証をかけています。認証情報は下記のとおりです。

```
ID: retty_gourmet_open
PW: aiNeT2ca
```

#### エンドポイント
存在するエンドポイントは下記のとおりです。

- /healthz
    - 疎通確認用エンドポイント
    - このエンドポイントは BASIC 認証が不要
- /api/restaurants
    - レストラン情報取得用エンドポイント
    - このエンドポイントは BASIC 認証が必要
    - HTTP メソッドは GET のみ

#### その他の要件
- できるだけワンライナー(1行)で答えてください
    - 複数行でも OK ですが、ワンライナーの方が得点が高いです



## 回答・解説
### 接続確認
まずは接続を確認するために下記のコマンドを実行します。  
`-s` オプションを使用すると進捗を非表示にできますが、こちらは任意で付与してください。
```
$ curl -s "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/healthz"
{"status":"OK"}
```

### BASIC 認証
接続が確認できたら、次に `/api/restaurants` のエンドポイントにアクセスしてみたいと思います。仕様書にもある通りこのエンドポイントには BASIC 認証がかけられているので、認証を通す必要があります。  

curl で BASIC 認証を通すためには下記のような方法があります。どれを使用しても OK です。
##### URL に埋め込む方法
`https://{ID}:{PW}@URL` の形式でアクセスすることで認証情報を渡すことができます。
```
$ curl -s "https://retty_gourmet_open:aiNeT2ca@p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants"
```


##### `-u` オプションを使用する方法
下記のように `-u` オプションのあとに `ID:PW` 形式で認証情報を付与することで認証することができます。
```
$ curl -s -u "retty_gourmet_open:aiNeT2ca" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants"
```

##### `Authorization` ヘッダを付与する方法
HTTP のリクエストヘッダに `Authorization` を付与することでも認証可能です。詳しくは [MDN Web Docs](https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Authorization) に書かれています。

curl では `-H` オプションで任意のヘッダーを付与することができます。上記のサイトにもあるように、認証情報はユーザ名とパスワードをコロンで結合して base64 エンコードする必要があるため、次のように記述します。  
```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants"
```

### jq を使用した絞り込み
分かりやすいようにまずは1つずつ解説します。

#### area.area_id が 14
`select({Key} == {Value})` のように書くと、特定の Key:Value に一致するものを絞り込むことができます。

```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select(.area.area_id == "14")'
```

#### restaurant_name に「珉」が含まれている
`select({Key} | contains("{Value}"))` のように書くと、特定の Key に特定の Value が含まれているものを絞り込むことができます。
```
$ curl -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select(.restaurant_name | contains("珉"))'
```

#### restaurant_name に「店」が含まれていない
`select({Key} | contains("{Value}") | not)` のように書くと、特定の Key に特定の Value が含まれていないものを絞り込むことができます。2つ目の条件を NOT にしたものですね。  

```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select(.restaurant_name | contains("店") | not)'
```

#### restaurant_name にアルファベットが含まれていない
`select({Key} | match("{regular expression}"))` のように書くと、正規表現を使用して絞り込むことができます。  
「アルファベットを含まない」の正規表現は `^(?!.*[a-zA-Z]).*$` なので次のようになります。

```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select(.restaurant_name | match("^(?!.*[a-zA-Z]).*$"))'
```

`select({Key} | test("{regular expression"}))` でも正規表現を使用できます。この場合は `| not` で否定することで「正規表現に一致するもの以外 = アルファベットを含んでいない」を絞り込むことができます。
```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select(.restaurant_name | test("[a-zA-Z]") | not)'
```

### 条件の結合
jq では `and` や `or` 演算子を使用することができます。今回はすべての条件を満たす店舗なので、すべての条件を `and` で結合します。

```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select((.restaurant_name | match("^(?!.*[a-zA-Z]).*$")) and (.restaurant_name | contains("珉")) and (.area.area_id == "14") and (.restaurant_name | contains("店") | not))'
```

また今回の場合はすべてが and 条件であるため、個別で絞った出力結果を `|` で結合しても答えを出すことができます。
```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select(.area.area_id == "14")' | jq '. | select(.restaurant_name | contains("珉"))' | jq '. | select(.restaurant_name | contains("店") | not)' | jq '. | select(.restaurant_name | match("^(?!.*[a-zA-Z]).*$"))'
```

### レストラン ID の抽出
もう答えは出ていますが、下記のようにして目的のキーだけを絞ることができます。`-r` はダブルクォーテーションを表示させないオプションです。

```
$ curl -s -H "Authorization: Basic $(echo -n retty_gourmet_open:aiNeT2ca | base64)" "https://p3hpemjpit.ap-northeast-1.awsapprunner.com/api/restaurants" | jq '.[] | select(.area.area_id == "14")' | jq '. | select(.restaurant_name | contains("珉"))' | jq '. | select(.restaurant_name | contains("店") | not)' | jq '. | select(.restaurant_name | match("^(?!.*[a-zA-Z]).*$"))' | jq -r .restaurant_id
```


## 参考情報
- [jqコマンドで覚えておきたい使い方17個](https://orebibou.com/ja/home/201605/20160510_001/)
