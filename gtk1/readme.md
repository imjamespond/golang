```
pacman -S mingw-w64-x86_64-toolchain base-devel
pacman -S mingw-w64-x86_64-gdb
```

vscode 使用gdb调试  

查看header和libs
```
pkg-config.exe --cflags gtk4
pkg-config.exe --libs gtk4
```
### **c_cpp_properties.json** 中加入includePath, IDE提示用  
*打开<gtk/gtk.h>通过vscode分别加入报错的header,有些需要打开更深层的header再加入基所需的header*  
注意有些header在mingw64/lib, 例如glib-2.0, graphene-1.0  
### **tasks.json** 加入includepath and ldpath, 编译器用
tasks.json中的args不能有空格，可以将pkg-config的参数copy后，用ctrl+f search and replace, then format 文档  
args 中的${files} 改为 "*.cpp",

