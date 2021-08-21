# タイトル: URLのパターンを調べてお店を探し当てよう!

URLが複数記載されたリスト urls.txt から、抽出要件を満たすURLを抽出してください。
抽出した各URLの中で最後の小軸パラメータの下1桁を抜き出し、先頭のURLから順につなげた値が答えです。

例: 抽出したURLが次の順序だった場合、答えは「1234」になります。

- https://retty.me/category/LCAT1/
- retty://retty.me/area/PRE09/ARE542/SUB54105/100001264572/
- https://retty.me/area/PRE09/PUR3/
- retty.me/area/PRE13/ARE13/SUB704/

## 抽出要件

次に示すscheme/authority/pathの要件をすべて満たすURLを抽出してください。
(各要素の意味はRFC3986の「3. 構文の構成要素」などを参照)

### scheme: 次のいずれかに完全一致するか存在しない(= authorityから始まる)

- https
- retty

### authority : 次のいずれかに完全一致する

- retty.me
- retty.co.th
- retty.world
- reserve.retty.me

### path : 「/(大軸)/(小軸1)/(小軸2)/.../(小軸N)/(お店ID)/」 のパターンに一致する

pathの構成要素は大軸、小軸、お店IDから成ります。

#### 大軸 (必須)

- area
- purpose
- category

#### 小軸 (必須)

小軸は識別子とパラメータを組み合わせることで1個を構成します、両者必須です。

- 識別子: 上にあるものほど先に現れます (例: AREのあとにPREは現れない)
    - PRE
    - ARE
    - SUB
    - STAN
    - LND
    - LCAT
- パラメータ
    - PRE: 2桁以下の整数、1桁の場合は0埋めがある (例: PRE01)
    - PRE以外: 0以上の整数
    
(有効な小軸の例)

- PRE01
- ARE2
- STAN34567

(無効な小軸の例)

- ABC1 (識別子に定義されていない)
- CAT (パラメータがない)
- PRE100 (パラメータの要件を満たしていない)

### お店ID (任意)

最後の構成要素にお店ID(12桁の整数)が付与されていることもあります。
付与されていないURLも有効です。
#pathの例

有効

- /category/LCAT6/ 
- /area/PRE13/ARE14/SUB1405/STAN6033/
- /purpose/PRE13/ARE13/SUB1303/100000700357/

無効

- /abc/PRE13/ (大軸が要件を満たしていない)
- /area/PRE999/ (小軸(PRE)のパラメータ要件を満たしていない)
- /area/PRE01/ARE-1/SUB/ (小軸(ARE)のパラメータ要件を満たしていないかつSUBのパラメータがない)
- /PRE13/ARE13/SUB1303/100/ (お店IDが要件を満たしていない)

### その他要件
大文字小文字は区別してください
末尾スラッシュの有無は抽出条件に関係ありません(どちらも有効です)


# 解答
## 抽出条件を満たす正規表現のパターン

```
^((https|retty):\/\/)?(reserve\.)?retty\.(me|co\.th|world)\/(purpose|area|category)(\/PRE\d{1}\d{1})?(\/ARE\d{1,})?(\/SUB\d{1,})?(\/STAN\d{1,})?(\/LND\d{1,})?(\/LCAT\d{1,})?(\/\d{12})?(\/)?$
```

## 答え

43632415

