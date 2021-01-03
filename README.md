# go-checkout

## Usage
gitのブランチを移動します(表示されるブランチは最新コミットの日時順)
```
[0] refs/heads/master
[1] refs/heads/b
>> Select branch number you move on: 0
You moved on: refs/heads/master
```

## setup 
```
// goの環境を整備した上でコンパイル 
> go build -o hoge main.go  
// shellの設定ファイルに実行ファイルへのエイリアスを追加など
> echo alias ggc='/path/to/this/repository/hoge'

// 実行
> ggc 
```

## TODO
- ソートの基準追加
- 最新コミットのAuthor自分(開発用)/自分以外(レビュー用)でフィルタリング
- そもそもCLIとしての配布方法がたぶん不適切
