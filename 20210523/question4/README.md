# 問題4
## コンテナをデバッグしたい！中身を確認したいけど…
Docker コンテナをデバッグしたい時には `docker exec` などを実行してコンテナ内でシェルを実行するのが一般的です。  

では distroless イメージなどのシェルが含まれていないコンテナをデバッグするにはどうすれば良いでしょうか？

distroless イメージを使用したコンテナイメージを用意しました。あの手この手でコンテナの中身を調べて答えを導き出してください。
```
docker run --rm public.ecr.aws/q0z9k4a9/retty_ctf:latest
```

## 回答・解説
問題文にあるようにコンテナを実行してみます。
```
$ docker run --rm public.ecr.aws/q0z9k4a9/retty_ctf:latest
ヒント: コンテナのどこかに隠されています
```
コンテナの中身に何かが隠されているようです。  
中に入って確認したい所ですが、問題文にもあるようにシェルが含まれていません。  
`docker history` を使用して、Docker イメージの構造を調べます。`--no-trunc` オプションを付与することで、すべての情報が表示されます。
```
$ docker history --no-trunc public.ecr.aws/q0z9k4a9/retty_ctf:latest
IMAGE                                                                     CREATED        CREATED BY                                                            SIZE      COMMENT
sha256:cea02eb0ae712ae532c8f4b0674b80c412936a576e19741987be945e9e121e36   2 days ago     CMD ["/bin/retty/linux/message"]                                      0B        buildkit.dockerfile.v0
<missing>                                                                 2 days ago     COPY ./retty_kun.png ./retty_kun.png # buildkit                       90.4kB    buildkit.dockerfile.v0
<missing>                                                                 2 days ago     COPY ./flag.txt /flag.txt # buildkit                                  175B      buildkit.dockerfile.v0
<missing>                                                                 2 days ago     COPY /bin/retty/darwin/message /bin/retty/darwin/message # buildkit   2.28MB    buildkit.dockerfile.v0
<missing>                                                                 2 days ago     COPY /bin/retty/linux/message /bin/retty/linux/message # buildkit     2.18MB    buildkit.dockerfile.v0
<missing>                                                                 51 years ago   bazel build ...                                                       17.4MB    
<missing>                                                                 51 years ago   bazel build ...                                                       1.82MB    
```

`/flag.txt` という如何にもなファイルが見えるので、これを確認してみます。  
`docker run <image> cat /flag.txt` のような確認方法が使えないため `docker cp` を使用して、コンテナのファイルをホストマシンに持ってきたいと思います。

```
$ docker cp

Usage:  docker cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-
        docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH

Copy files/folders between a container and the local filesystem

Use '-' as the source to read a tar archive from stdin
and extract it to a directory destination in a container.
Use '-' as the destination to stream a tar archive of a
container source to stdout.

Options:
  -a, --archive       Archive mode (copy all uid/gid information)
  -L, --follow-link   Always follow symbol link in SRC_PATH

```

上記のように `docker cp` の使い方を見てみるとまずはコンテナを作成する必要があるようです。`docker create` でコンテナを作成します。  
(`docker run public.ecr.aws/q0z9k4a9/retty_ctf:latest` でも可)

```
$ docker create docker run --rm public.ecr.aws/q0z9k4a9/retty_ctf:latest
b0c01fc434feeff603a43124a66349f02595998ddc7ebb8f48908d6fe45163f3

$ docker ps -a
CONTAINER ID   IMAGE                                      COMMAND                  CREATED          STATUS    PORTS     NAMES
b0c01fc434fe   public.ecr.aws/q0z9k4a9/retty_ctf:latest   "/bin/retty/linux/me…"   33 seconds ago   Created             vigorous_grothendieck
```

ファイルをコピーして中身を見てみます。
```
$ docker cp b0c01fc434fe:/flag.txt ./
$ cat flag.txt
※ これは答えではありません

麻布十番のオススメランチです！行ってみてください！
https://retty.me/area/PRE13/ARE14/SUB1405/100000001471/
```

どうやらダミーだったようです。  
`docker history` を見ると、ヒントになりそうな画像ファイルがあったので、これも確認してみます。

```
$ docker cp b0c01fc434fe:/retty_kun.png ./
$ open retty_kun.png
```

![retty_kun](https://user-images.githubusercontent.com/29038315/119303412-a2f9c800-bca0-11eb-9297-0a61cd9f3582.png)

コンテナ内にある `/bin/retty/linux/message` に `retty` という引数を渡すと何かが起きるようです。  
`docker run` で上書きして実行してみます。
```
$ docker run --rm public.ecr.aws/q0z9k4a9/retty_ctf:latest /bin/retty/linux/message retty
https://retty.me/area/PRE13/ARE2/SUB201/100000821612/1019559/
```

すると正解の口コミ URL が表示されました！！

## その他・備考
実際の回答を見てみると、下記のように busybox をコンテナイメージに追加しているチームもありました。  
「シェルがないなら入れればいい」という素直な発想で面白いですね！今回は解き方を制限していないので、この方法でも正解です。

```Dockerfile
FROM public.ecr.aws/q0z9k4a9/retty_ctf:latest
ADD busybox "/bin/"
SHELL ["/bin/busybox","sh"]
CMD ["/bin/busybox","sh"]
```

```
$ docker build -t retty_ctf:busybox .
$ docker run --rm -it retty_ctf:busybox
# /bin/retty/linux/message retty
https://retty.me/area/PRE13/ARE2/SUB201/100000821612/1019559/
```

## 参考情報
- [[docker] CMD とENTRYPOINT の違いを試してみた](https://qiita.com/hihihiroro/items/d7ceaadc9340a4dbeb8f)
