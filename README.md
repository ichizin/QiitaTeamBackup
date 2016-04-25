# QiitaTeamBackup

QiitaTeamバックアップJSONファイルからmarkdownファイルを生成します

## 出力形式
ファイル名は```タイトル + .md```になります

生成されるMarkdownファイルは上から

- タイトル
- 内容
- コメント

で出力されます。

## 使い方

**```qiita_team_backup```バイナリで実行する場合**

```
// --in でinputするJSONファイルを指定
$ ./qiita_team_backup --in ./input/sample.json
```

**goコマンドで実行する**
```
// --in でinputするJSONファイルを指定
$ go run main.go --in ./input/sample.json
```
