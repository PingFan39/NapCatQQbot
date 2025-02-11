# NapCatQQbot

注：由中括号包围的描述需要替换为相应的内容

# 可用功能

每个功能冒号后面是发送文本的格式，格式中的空格可有可无

GPT对话：`@[你的QQbot] [任何不符合下面所有格式的文本]` 

接着上一次GPT对话：`@[你的QQbot] ， [任何不符合下面所有格式的文本]` （这个逗号是全角的）

帮助：`@[你的QQbot]` 或 `@[你的QQbot] 帮助` 或 `@[你的QQbot] help`

发送随机图片：`[指定内容]`

其他功能正在锐意制作中。。

# 使用教程

## 安装步骤

0. 前期准备

    - API-Key (重要！这可能是所有步骤中最难的一步)

    - bot的QQ

    - 能运行go的环境

1. 下载本项目

2. 下载[NapCatQQ](https://github.com/NapNeko/NapCatQQ)（不用点进去）

    - 打开终端，工作目录为本项目文件夹，执行以下命令（执行后可能要等一会）：

```
curl -o install.ps1 https://nclatest.znin.net/NapNeko/NapCat-Installer/main/script/install.ps1
powershell -ExecutionPolicy ByPass -File ./install.ps1 -verb runas
```

3. 打开终端，工作目录为NapCatQQ的文件夹，执行命令：`./launcher.bat [bot的QQ号]`

    - 按提示操作，直到没得操作

    - 做完第4步前别关

4. 进行网络配置

    1. 在第3步终端的输出中找到以`[info] [NapCat] [WebUi] WebUi Local Panel Url: `开头的一行（哪个都行），打开网址（记网址的`token=`后面的一串为token）

    2. 点击左侧的网络配置

    3. 点击右上角的添加配置

        - 名称：[任意]

        - 类型：HTTP服务器

        - 启用：开启

        - 端口：[1024~49151中任意值（记为server_port）]

        - 主机：0.0.0.0

        - （下面还有的，鼠标滚几下就行）

        - 消息格式：String

        - Token：[记下的token]

        - 其他随意，最后点确认

    4. 再点击添加配置

        - 名称：[任意]

        - 类型：HTTP客户端

        - 启用：开启

        - URL：`http://localhost:`[1024~49151中任意值，不能和server_port一样（记为client_port）]

        - 消息格式：String

        - Token：[不用填]

        - 其他随意，最后点确认

    5. 网址就可以关了

5. 打开项目的main.go，按注释在双引号中填写内容

6. 没了

## 日常使用

- 打开终端，工作目录为NapCatQQ的文件夹，执行命令：`./launcher.bat [bot的QQ号]`

- 打开终端，工作目录为本项目文件夹，执行命令：`go run main.go`

# 文件来源

src/GPT/openai/openai.go：[gobot_gpt_api](https://github.com/OIerNekoPass/gobot_gpt_api)/gpt_chat.go

src/qq_reply/qq_reply.go：[NapCatQQreply](https://github.com/OIerNekoPass/NapCatQQreply)/qq_reply.go

src/QQbot.go：[OIerNekoPass](https://github.com/OIerNekoPass)/?