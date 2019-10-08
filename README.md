# Get Started

[https://dev.classmethod.jp/etc/serverless-framework-golang-lambda-create/](https://dev.classmethod.jp/etc/serverless-framework-golang-lambda-create/)

下記コマンドを実行すれば、awsへデプロイできる。

```
$ cd goserverless
$ GOOS=linux go build -o bin/main
$ ../node_modules/.bin/serverless deploy
```

lambda用のgoモジュールをダウンロードしていないので、あれば下記コマンド実行。

```
$ go get github.com/aws/aws-lambda-go/lambda
```
