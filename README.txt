项目需要的东西：
参考https://blog.csdn.net/love906897406/article/details/125353510

#目前该项目实现了2种功能：1、有界面的功能：绘制界面、建立socket、心跳包、重连、账密登录功能  2、无界面的功能：读取excel，实现批量登录（好像批量登录偶尔会卡主，没有收到回包）
#目录结构
#basefunc目录：实现基础功能，目前是日志打印和结构体序列化和反序列化
#business目录：用于用户业务信息
#fileopera目录: 目前实现了Excel文档读取
#image目录：qt用到的图片及图片路径
#protocol目录：消息协议定义，对接C++Agent代理
#qtui目录：QT Designer中搭建好ui界面后，保存生成的.ui格式文件
#tcpconn目录：socket实现和登录基础逻辑实现
#uitogo目录：带_ui后缀的都是通过golang的qt命令"goqtuic"通过.ui格式自动生成的代码文件，实现了界面的绘制,命令见createuitogo.bat，没有_ui后缀的文件是参考了_ui的界面逻辑，改进实现的逻辑

#其他目录
#qtbox目录：golang实现的qt动态库，目的是在debug模式下加快编译和调试，没有qtbox每次调试还要编译C++的库，非常慢
#deploy目录：qtDeploy命令生成的指定平台可以运行的基础文件和.exe文件，正式版用该方式编译，参考https://www.cnblogs.com/apocelipes/p/9300335.html
#github.com\therecipe目录下有两个qt，一个qt中有C++文件，一个没有。   qtDeploy编译时目录选有C++文件的， 平时调试时用没有C++文件的