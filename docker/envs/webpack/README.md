
# node webpack環境
webpack5＋ホットリロード機能を使って、VueやReactを使わずに気軽にHTML/CSS/Javascriptのサンプルを作りたいと思って作成した環境です。

単独のリポジトリとはなっておらず、加えてDocker環境もfishシェル前提だったりと、汎用的(一般的)な環境にはなっていません。（あくまで自分自身用の環境です)

## インストール
```
$ npm install
$ npm run build   # ビルドする場合
$ npm run start   # サーバーを起動する場合
```

## 環境構築メモ
ディレクトリ構成
```
├── public
│   ├── index.html
│   ├── sample.html
│   ├── images
│   │   └── sample.png
│   └── js
│       └── bundle.js
├── src
│   ├── html
│   │   ├── index.html
│   │   └── sample.html
│   ├── images
│   │   └── sample.png
│   └── js
│       └── index.js
├── package-lock.json
├── package.json
└── webpack.config.js
```

### 初期設定・ホットリロードのための開発用サーバー
```
$ npm init -y

# webpackとcliから操作するためのパッケージ
$ npm i -D webpack webpack-cli
# 開発用サーバー
$ npm i -D webpack-dev-server
```

package.jsonのscriptsフィールドの編集
```
  "scripts": {
    "build": "webpack",
    "start": "webpack serve"
  },
```

webpack.config.js
```
const path = require('path')

module.exports = {
  context: path.resolve(__dirname, 'src'),
  mode: 'development',
  entry: './js/index.js',
  output: {
    path: path.resolve(__dirname, 'public'),
    filename: 'js/bundle.js'
  },
  devServer: {
    open: true,
    port: 9000,
    contentBase: './public'
  }
}
```

### HTMLもバンドル可能とする
参考: [Webpackを使用してhtmlの共通部分をテンプレート化](https://chocolat5.com/tips/webpack-html-template/)

```
$ npm i -D html-webpack-plugin html-loader
```

webpack.config.jsの編集
(pluginsでindex.htmlとsample.htmlを指定している)
```
const HtmlWebpackPlugin = require('html-webpack-plugin')
  module: {
    rules: [
      {
        test: /\.html$/,
        loader: "html-loader"
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: './html/index.html',
      filename: 'index.html',
      inject: true,
      hash: true
    }),
    new HtmlWebpackPlugin({
      template: './html/sample.html',
      filename: 'sample.html',
      inject: true,
      hash: true
    })
  ],
```

### 画像を指定したディレクトリに出力する
<code>src/images</code>配下に格納されている画像ファイルを<code>public/images</code>に出力する。

webpack5から追加されたAsset Modulesを使えば良いので別途インストールするモジュールはなし。
webpack.config.jsにルールを追加する。
- [Asset Modules](https://webpack.js.org/guides/asset-modules/)

```
  module: {
    rules: [
      {
        // Asset Modulesの処理対照ファイル
        test: /\.(jpe?g|gif|png|svg)$/,
        // Asset Modulesのタイプ
        type: 'asset/resource',
        generator: {
          // 画像の出力先と出力ファイル名
          // [name]にはファイル名、[ext]には拡張子が入る
          filename: 'images/[name][ext]',
          publicPath: './',
        }
      },
    ]
  },
```
