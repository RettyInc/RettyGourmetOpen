# 問題5: ヒントを掘り当て消えたブランチを取り戻せ!

ソースコード(リポジトリ)は[こちら](https://drive.google.com/file/d/1w-6CeIualFxUsxBu63aiyn8E7GqhABZA/view?usp=sharing)

## 解答

* `git grep "HINT" ${git rev-list --all}` でcommit内容に"HINT"があるものを検索
* すると "HINT: 消したブランチ上での最後のコミットメッセージには "complete" と記述した!" という内容にヒット
* そこで `git reflog --grep-reflog="complete"` で過去のcheckoutを検索
* 出てきたcommit ID `98b99c1` より `git log master..98b99c1` すればcommit messageが表示される
* 必要な情報を得るための整形を加えて `git log master..98b99c1 --oneline | awk '{print $2}'` でcommit messageのみを抽出

## 解説

### `git grep`

git grepは引数で指定したcommit時点のリポジトリ上のファイル(git管理下のもの)から、同じく引数指定のキーワードを含むファイルと該当箇所を表示します。

```shell
git grep "KEYWORD" 1234567
```

commit IDを指定しない場合は現在checkoutしているコードで検索します。
またcommit IDを複数指定することができます。

```shell
git grep "KEYWORD" # 現在checkoutしているコード上で検索
git grep "KEYWORD" 1234567 890abcd # 複数commitで検索
```

ところが今回はどのcommitでHINTが加えられたかが不明で、かつその後に内容は消されています。
そこでこれまでの全commitに対して検索をかけることにします。

### `git rev-list`

commit IDをリストするには `git rev-list` を使います。
これはcommit logからIDのみを抽出し日時の降順に表示するコマンドです。
これに `--all` オプションをつけることでマージされていないブランチなどすべてのcommit logを対象にIDをリストすることができます。

```shell
git rev-list # HEADに対するcommit logのIDのみを表示
git rev-list --all # 全commitに対するcommit logのIDのみを表示
```

このコマンドの出力結果を先程の `git grep` に与えることで「全commitの中から"HINT"を内容に含むファイルと該当箇所を抽出」することができます。

```shell
g grep "HINT" $(git rev-list --all)

> 17c81dc8de139d2e0ee27401c149731c2e6e8d9d:q5.txt:HINT: 消したブランチ上での最後のコミットメッセージには "complete" と記述した!
```

これで手順1は終了です。

### `git reflog`

commit IDが分かったので次は消えたブランチを探します。
ローカルリポジトリ上でcheckoutをするとgit上ではすべての操作の履歴が記録されるようになっています。
この記録を見るには `git reflog` を使います。

```shell
git reflog
> 1234567 HEAD@{1}: commit: wine
> 890abcd HEAD@{2}: commit: beer
...
```

これで "complete" というcommit messageを含むcommitを探します。
表示をスクロールして目視で探しても良いですが `--grep-reflog` オプションを使って検索すると早いです。

```shell
git reflog --grep-reflog="complete"
> 02fe374 HEAD@{66}: commit: complete
```

## `git log --oneline` & `awk`

あとはここにcheckoutしmasterからのログを表示させれば答えが出ます。

```shell
git log master..02fe374~1

> commit 98b99c1ec2850937be97ab9d1c8096b415b04903 (HEAD -> feature3, origin/feature3)
> Author: Tomohiro Imaizumi <bonriki.life@gmail.com>
> Date:   Sat May 22 19:31:12 2021 +0900
>
>     retty.me
>
> commit be4075b94cef4807ee97407e730a98786c4d9cd5
> Author: Tomohiro Imaizumi <bonriki.life@gmail.com>
> Date:   Sat May 22 19:30:54 2021 +0900
>
>     area
> ...
```

`--oneline` オプションを使うとよりすっきりです、 `awk` を使えばさらにすっきり!

```shell
git log --oneline master..02fe374~1 | awk '{print $2}'

> retty.me
> area
> PRE13
> ARE2
> SUB201
> 100000821612
> 1019559
```

## 参考資料

* [Git \- git\-grep Documentation](https://git-scm.com/docs/git-grep)
* [Git \- git\-rev\-list Documentation](https://git-scm.com/docs/git-rev-list)
* [Git \- git\-reflog Documentation](https://git-scm.com/docs/git-reflog)
* [Git \- git\-log Documentation](https://git-scm.com/docs/git-log)
* [とほほのAWK入門 \- とほほのWWW入門](https://www.tohoho-web.com/ex/awk.html)
* [(再演) detached HEADを理解して脱Git初心者を目指す方のためのGit入門勉強会 @サポーターズCoLab - Speaker Deck](https://speakerdeck.com/imaizume/zai-yan-detached-headwoli-jie-sitetuo-gitchu-xin-zhe-womu-zhi-sufang-falsetamefalsegitru-men-mian-qiang-hui-at-sapotazucolab?slide=41)
* [きれいなcommit, pull requestを知りたい/作りたい方のためのgit勉強会 - Speaker Deck](https://speakerdeck.com/imaizume/zuo-ritaifang-falsetamefalsegitmian-qiang-hui)

