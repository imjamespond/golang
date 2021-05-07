## Main points
ts, webpack dev server, script inject, react  
```
cnpm i --save-dev webpack webpack-cli webpack-dev-server \
html-webpack-plugin@next clean-webpack-plugin \
style-loader css-loader file-loader
```

## 配置 tsconfig.json
启动webpack时plugin报ts错误, 原因是将ts限定在src路径内
```
...
  "include": [
    "src"
  ]
```
[支持jsx,es5](https://webpack.js.org/guides/typescript/)
```
  "compilerOptions": {
    "outDir": "dist",
    "module": "es6",
    "target": "es5",
    "jsx": "react",
```
outDir不指定可能会报**Cannot write file ... because it would overwrite input file**之错  
支持动态加载  
``"module": "esnext",``  
支持文件导入, important for import @...  
``"moduleResolution": "node",``  
关于alias  
```
# tsconfig加入 
    "baseUrl": ".",
    "paths": {
      "@denver/*": ["./src/denver/*"]
    }, 
# webpack加入
    resolve: {
      alias: { 
        '@denver': path.resolve(__dirname, '../src/denver'),
      },
```

## React
[typescript](https://www.typescriptlang.org/docs/handbook/react.html)  
npm install --save react react-dom @types/react @types/react-dom

## [Loaders](https://webpack.js.org/loaders/) and plugins  


[ts-loader](https://github.com/TypeStrong/ts-loader)  The tsconfig.json file controls TypeScript-related options so that your IDE, the tsc command, and this loader all share the same options. **tsconfig.json控制所有相关配置, 因此IDE, tsc命令, 和ts-loader全部共享一套选项**  


``npm i --save-dev html-webpack-plugin@next``  
HtmlWebpackPlugin 将js注入html  


SplitChunksPlugin 允许将通用的依赖放到已存在的entry chunk中,或放到全新的chunk中  


``npm install --save-dev clean-webpack-plugin``
清除所有output的内容

``npm install --save-dev style-loader``
将css注入dom
``npm install --save-dev css-loader``
解释@import 和 url()
``npm install file-loader --save-dev``
解析import/require()导入的文件成url,并将文件生成到output目录

## 支持 IE 11
polyfill  
[``import "core-js/stable";``](https://www.npmjs.com/package/core-js)  
webpack 配置[multiple targets](https://webpack.js.org/configuration/target/#string-1)  
```
module.exports = {
  // ...
  target: ['web', 'es5']
};
```
或 package.json 的 [browserslist](https://webpack.js.org/configuration/target/#browserslist)
