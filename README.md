# go_rss

https://qiita.com/tmknom/items/08b69594e32a92bccee5

## パッケージ分ける方法がスーパー特殊でマジ卍

1. GOPATHにプロジェクトのルートを設定
1. src/github/自分のアカウント名/アプリ名　のディレクトリを作成
1. 使いたいモジュール書く
1. 使いたいモジュールを
  
  1. コミットします

  1. プッシュします

1. プッシュが成功したらimportにgithubから始まるパッケージのパスを書く

これでモジュールを参照して利用できるようになる

ディレクトリにドメイン名がちゃんと含まれるようになりそうで、そこはいいと思った。

でもプロジェクトの文脈にgithubががっつり入ってくるの好きじゃない。（受け付けないわけではない

あとgithubと一蓮托生みたいな雰囲気があんまり好きじゃない。(受け付けないわけではない

参考
https://qiita.com/shopetan/items/eddcacec21cc7ea274f9

## 命名規則わっかんねぇ・・・

funcがlower camelだったりupper camelだったりするのなんなの・・・


  →Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポート（公開）された名前( exported name )です。 例えば、 Pi は math パッケージでエクスポートされています。


  　ってtour of goが言ってた


funcは動詞で始まるイメージがあったけど、他人のサンプル"NoteUpdate"とか"NewFile"とかあってわっかんねぇ・・・
