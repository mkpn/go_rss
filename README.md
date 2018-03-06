# go_rss

https://qiita.com/tmknom/items/08b69594e32a92bccee5

## パッケージ分ける方法がスーパー特殊だった

1. GOPATHにプロジェクトのルートを設定
2. src/github/自分のアカウント名/アプリ名　のディレクトリを作成
3. 使いたいモジュール書く
4. 使いたいモジュールを
    1. コミットします
    2. プッシュします
5. プッシュが成功したらimportにgithubから始まるパッケージのパスを書く

これでモジュールを参照して利用できるようになる

https://qiita.com/shopetan/items/eddcacec21cc7ea274f9

## 命名規則わっかんねぇ・・・

funcがlower camelだったりupper camelだったりするのなんなの・・・


  →Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照できるエクスポート（公開）された名前( exported name )です。 例えば、 Pi は math パッケージでエクスポートされています。


  　ってtour of goが言ってた


funcは動詞で始まるイメージがあったけど、他人のサンプル"NoteUpdate"とか"NewFile"とかあってわっかんねぇ・・・
