# 运行方法
##### 先clone仓库
* 项目根目录下 go run .  [ 或直接运行 main.go ]
***
# 打包成linux 程序软件
## 项目根目录下
###### 安装fyne命令工具:
    go install fyne.io/fyne/v2/cmd/fyne@latest  
###### 打包
    fyne package -os linux -icon 1.png   
***
###### windows打包命令: 
    fyne package -os windows -icon 1.png
***
###### 打包成功后会得到一个压缩包(上面仓库中提供了linux版本的程序压缩包[可直接使用])，
###### 打开解压后的文件夹，进入根目录右击运行终端
###### `sudo make install`即可安装
###### `sudo make uninstall`卸载
***
###### 注意:Linux环境下C++默认已经安装，fyne环境几乎无需配置，
###### 但是Windows环境下C++需要安装，等环境，具体见fyne官网：
###### [fyne官网](https://docs.fyne.io/)
***
###### 运行示例
![运行示例1](/eg/eg1.png "运行示例1")
***
![运行示例2](/eg/eg2.png "运行示例2")
