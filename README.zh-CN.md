# 功能描述

Based on edge-tts encapsulation of speech to text, speech is generated according to text, the same copy is generated only once, real-time generation
基于edge-tts封装语音转文字，根据文本生成语音，相同文案只生成一次，实时生成服务部署，推荐nginx本地代理。

### 测试体验入口

Web UI：
[https://tts-api.wdft.com](https://tts-api.wdft.com/)

### 部署服务环境

```
pip3 install edge_tts
```

### 基于edge-tts封装语音转文字，根据文本生成语音，相同文案只生成一次，实时生成

编译Linux包

config.json:

```

{
"serverURI": "localhost",
"localPort": 59443,
"audioPath": "/media-files/",
"mediaFilesRoot": "/home/wwwroot/tts-service",
"validTokenStr": "github.com/ljq"
}

```

**热更新调试、构建推荐使用air：**

```
# .air.toml

# build
[build]
cmd = "env GOOS=linux GOARCH=amd64 go build -o ./build/bin/tts-service tts_service.go"

# 启动命令
[run]
port = "8081"
```

### 接口请求方式

音频文件部署web root目录：

```

mkdir -p /home/wwwroot/tts-service/media-files

### nginx反代

/home/wwwroot/tts-service

```

web请求音频文件路径：https://tts-api.wdft.com/media-files/

```

### GET或POST(使用Nginx反代)

http://localhost:52843/tts?text=txt&role=zh-CN-YunxiNeural&access_token=ljq@GitHub

```

**test access_token: 02bd6f28f95dc84cfbf22e0ce1083f77**

### 守护进程脚本(service-manager.sh)

###### 常用命令参数(start|stop|restart),默认进程守护

# 定义服务名称和启动命令

```
SERVICE_NAME="tts-service"
START_COMMAND="/data/edge-tts/tts-service"
```

```

nohup /data/edge-tts/tts-service > /dev/null 2>&1 &

```

###### 角色

- 云曦: 【zh-CN-YunxiNeural】
- 晓伊: 【zh-CN-XiaoyiNeural】
- 筱筱: 【zh-CN-XiaoxiaoNeural】
- 云间: 【zh-CN-YunjianNeural】
- 云霞: 【zh-CN-YunxiaNeural】
- 云扬: 【zh-CN-YunyangNeural】
- 辽宁-晓贝: 【zh-CN-liaoning-XiaobeiNeural】
- 陕西-晓妮: 【zh-CN-shaanxi-XiaoniNeural】
- 香港-晓佳: 【zh-HK-HiuGaaiNeural】
- 香港-晓曼: 【zh-HK-HiuMaanNeural】
- 香港-云龙: 【zh-HK-WanLungNeural】
- 台湾-晓晨: 【zh-TW-HsiaoChenNeural】
- 台湾-晓瑜: 【zh-TW-HsiaoYuNeural】
- 台湾-云哲: 【zh-TW-YunJheNeural】

```

```
