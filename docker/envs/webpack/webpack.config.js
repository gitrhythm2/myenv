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
        test: /\.(jpe?g|gif|png|svg)$/,
        type: 'asset/resource',
        generator: {
          filename: 'images/[name][ext]',
          publicPath: './',
        }
      },
      {
        test: /\.scss$/,
        include: path.resolve(__dirname, 'src/scss'),
        use: [
          'style-loader',
          'css-loader',
          'sass-loader'
        ]
      },
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

