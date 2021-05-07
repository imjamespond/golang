const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const { publicPath } = require('./variables')

module.exports = (env, options) => {

  console.log(options)

  return {
    entry: './src/index.tsx',
    output: {
      path: path.resolve(process.cwd(), 'dist'), //path.resolve(__dirname, '../dist'),
      filename: '[name]-[chunkhash]-bundle.js',
      chunkFilename: '[id]-[chunkhash]-chunk.js',
      publicPath
    },

    resolve: {
      alias: {
        '@assets': path.resolve(__dirname, '../src/assets'),
      },
      extensions: ['.tsx', '.ts', '.js'],
    },
    module: {
      rules: [
        {
          test: /\.tsx?$/,
          // use: 'ts-loader',
          use: [{
            loader: 'ts-loader',
          }],
          exclude: /node_modules/,
        },
        {
          test: /\.css$/i,
          use: ['style-loader', 'css-loader'],
        },
        {
          test: /\.(png|jpe?g|gif)$/i,
          use: [
            {
              loader: 'file-loader',
            },
          ],
        }
      ],
    },
    plugins: [
      new HtmlWebpackPlugin({
        template: 'public/index.html'
      }),
    ],
    optimization: {
      splitChunks: {
        cacheGroups: {
          react: {
            test: /[\\/]node_modules[\\/](react|react-dom|scheduler)/,
            name: 'react',
            chunks: 'initial',
            enforce: true,
            priority: 1 //优先处理,否则会合并到commons
          },
          commons: {
            test: /[\\/]node_modules[\\/]/,
            name: "vendors",
            chunks: 'initial',
            enforce: true,
          }
        }
      }
    }
  }
}