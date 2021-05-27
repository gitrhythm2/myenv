const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
  context: path.resolve(__dirname, 'src'),
  mode: 'development',
  entry: './js/index.js',
  output: {
    path: path.resolve(__dirname, 'public'),
    filename: 'js/bundle.js'
  },
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
  devServer: {
    open: true,
    port: 9000,
    contentBase: './public'
  }
}

