# 問題
public.ecr.aws/a8o1y5d1/saku2saku:retty_open_20210619 というdockerイメージを取得し、イメージに記述されたステップに従い暗号化されたファイルとパスワードを見つけ、最終的にopensslにより正解を見つけ出そう！

opensslのエンコードはAESのCipher Block Chaining動作モードで暗号化しました。暗号化に用いた鍵長は256bitです。

openssl enc -d {エンコード方法} -salt -k {パスワード} -in {暗号化されたファイル} -out {出力先}

# ポイント説明
1. docker pullによるdockerイメージの取得
2. docker history or docker inspectによるステップ確認
3. docker cp によるファイルコピー
4. base64デコード
5. nkf等による文字エンコードと改行コードの修正
6. md5値によるパスワードの確認
7. opensslによる復号化

# 解法
まずはdockerイメージをpullで取得します。

```shell
$ docker pull public.ecr.aws/a8o1y5d1/saku2saku:retty_open_20210619
```

イメージを取得したらhistoryまたはinspectでdockerイメージをみます。

```shell
$ docker history --no-trunc public.ecr.aws/a8o1y5d1/saku2saku:retty_open_20210619
IMAGE                                                                     CREATED              CREATED BY                                                                                                                    SIZE      COMMENT
sha256:743496a254a18fc1f631dd54d3e349627feef53e2c1c1abcaa625e229df31f8e   About a minute ago   CMD ["/bin/sh"]                                                                                                               0B        buildkit.dockerfile.v0
<missing>                                                                 About a minute ago   COPY retty /welcome/to/retty # buildkit                                                                                       121B      buildkit.dockerfile.v0
<missing>                                                                 About a minute ago   COPY encrypted.txt /path/to/encrypted.txt # buildkit                                                                          96B       buildkit.dockerfile.v0
<missing>                                                                 About a minute ago   LABEL Step3=Change file(/welcome/to/retty) encode and line feed and see file contents. UTF-8(without BOM) and LF line feed.   0B        buildkit.dockerfile.v0
<missing>                                                                 About a minute ago   LABEL Step2=Decode file(/welcome/to/retty) by this way https://gist.github.com/saku/58679ac9162bf17170912d832e95aa34.         0B        buildkit.dockerfile.v0
<missing>                                                                 About a minute ago   LABEL Step1=Encrypted file is /path/to/encrypted.txt. Let's get this file from container to local.                            0B        buildkit.dockerfile.v0
<missing>                                                                 3 days ago           /bin/sh -c #(nop)  CMD ["/bin/sh"]                                                                                            0B
<missing>                                                                 3 days ago           /bin/sh -c #(nop) ADD file:6797caacbfe41bfe44000b39ed017016c6fcc492b3d6557cdaba88536df6c876 in /                              5.33MB
```

下記のステップを確認します。
```
Step1=Encrypted file is /path/to/encrypted.txt. Let's get this file from container to local.
Step2=Decode file(/welcome/to/retty) by this way https://gist.github.com/saku/58679ac9162bf17170912d832e95aa34.
Step3=Change file(/welcome/to/retty) encode and line feed and see file contents. UTF-8(without BOM) and LF line feed.
```

Step1, Step2の指示に従ってコンテナからファイルを取得します。

```shell
$ docker run -it --name retty_open public.ecr.aws/a8o1y5d1/saku2saku:retty_open_20210619
$ docker cp retty_open:/path/to/encrypted.txt ~/
$ docker cp retty_open:/welcome/to/retty ~/
```

Step2の説明を確認します。  
https://gist.github.com/saku/58679ac9162bf17170912d832e95aa34 を読み、 `/welcome/to/retty` を読んでbase64デコードすることを確認します。

```
データを64種類の印字可能な英数字のみを用いて、それ以外の文字を扱うことの出来ない通信環境にてマルチバイト文字やバイナリデータを扱うためのエンコード方式でエンコードされている。
上記方式に沿った方法でファイルをデコードすること。
```

```shell
$ cat ~/retty | base64 -d > ~/retty_decoreded 
```

Step3の説明を確認します。  
Step2でbase64 decodeしたファイルを文字コードUTF-8(without BOM)にして、改行コードをLFにします。  
nkfを使うと簡単です。

```shell
$ nkf -w -Lu --overwrite ~/retty_decoreded
```

文字コードと改行コードをなおしたらファイルの中身を確認します。
```shell
$ cat ~/retty_decoreded
あなたはだんだんRettyに入社したくなる
このファイルが読めるようになったらmd5値を取ろう！
```

上記に従って文字コードUTF-8(without BOM)にして、改行コードをLFにしたファイルのmd5値を確認します。

```shell
$ md5 ~/retty_decorded
af9aa132755d95816bf0dad3a74f8d63
```

これにて、パスワードと暗号化ファイルが揃ったのでopensslで復号化を行います。  
`opensslのエンコードはAESのCipher Block Chaining動作モードで暗号化しました。暗号化に用いた鍵長は256bitです。` とあるのでエンコード方法は `aes-256-cbc` であることがわかります。

これまでの内容を組み合わせると下記のようになります。

```shell
$ openssl enc -d -aes-256-cbc -salt -k af9aa132755d95816bf0dad3a74f8d63 -in ~/encrypted.txt
https://retty.me/area/PRE13/ARE14/SUB1401/100000000111/43632415/
```

ちなみにStep2とStep3をワンライナーでやるとこうなります。
```shell
$ openssl enc -d -aes-256-cbc -salt -k `cat ~/retty | base64 -d | nkf -w -Lu | md5` -in ~/encrypted.txt
https://retty.me/area/PRE13/ARE14/SUB1401/100000000111/43632415/
```
