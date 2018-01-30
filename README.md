# weather-api
天気API

OpenWeatherMapのAPIを使っています

前提条件

・Go 1.9

・$GOPATHのsrc/に配置する

・glideをインストールしていること

・OpenWeatherMapのアカウントを作成してAPIKEYを取得

```
# vendoring
glide update
```

```
# exampleからtomlを作成
cp config.toml.example config.toml

# config.tomlの編集
vi config.toml

# server start
go run main.go
```