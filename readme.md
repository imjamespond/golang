## Start your module using the ``go mod init`` command to create a go.mod file
``go mod init example.com/greetings``, 生产模式应该module可下载地址  
此命令创建一个go.mod 文件, 让 你的代码作为一个module 可区分. 此文件只命名你的module和所需go 版本,  
当你添加依赖时,go.mod file会列出指定的module版本  

## Edit the hello module to use the unpublished greetings module.
生产模式下,你可以发布你的modules到server,go会下载. 现在暂时来说,你需要将就caller's module,这样它才能找到在你本地的greetings code, 你可以将hello module’s go.mod file改成:  
1, ``replace example.com/greetings => ../greetings``  
2, ``go build``, 让go定位module并添加到依赖到go.mod, replace指令告诉go 到哪找greetings module因为其还没发布. 要引用发布的module, go.mod会忽略replace指令并使用 末尾有tag版本的require指令  

## 测试
Go支持单元测试, 使用命名约定,testing包,``go test``命令.  
1, 在greetings中增加一个以 _test.go结尾的文件.  
2, in the greetings directory, run the ``go test -v`` command to execute the test