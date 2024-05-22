# Description

Based on edge-tts encapsulation of speech to text, speech is generated according to text, the same copy is generated only once, real-time generation.

[简体中文](README.zh-CN.md)

### Demo page

Web UI：
[https://tts-api.wdft.com](https://tts-api.wdft.com/)

### Deploy server environment

```
pip3 install edge_tts
```

### Build a linux binary package

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

**systemd: tts-service.service**

```
#/etc/systemd/system/tts-service.service

sudo systemctl anable tts-service
sudo systemctl start tts-service
sudo systemctl stop tts-service
sudo systemctl restart tts-service
sudo systemctl status tts-service
```

**Recommended tool for hot update debugging and building: air：**

**.air.toml**

```
# .air.toml

# build
[build]
cmd = "env GOOS=linux GOARCH=amd64 go build -o ./build/bin/tts-service tts_service.go"

# port
[run]
port = "8081"
```

### API Request

Audio file deploy web root directory：

```
mkdir -p /home/wwwroot/tts-service/media-files

### nginx proxy
/home/wwwroot/tts-service
```

web request audio path：https://tts-api.wdft.com/media-files/

```
### GET or POST Method(Nginx proxy)
http://localhost:52843/tts?text=txt&role=zh-CN-YunxiNeural&access_token=ljq@GitHub

```

**test access_token: 02bd6f28f95dc84cfbf22e0ce1083f77**

### daemon process script(service-manager.sh)

###### Common command parameters (start | stop | restart), default process daemon

# custom bash parameters

```
SERVICE_NAME="tts-service"
START_COMMAND="/data/edge-tts/tts-service"
```

```
nohup /data/edge-tts/tts-service > /dev/null 2>&1 &
```

###### Roles

- yunxi: 【zh-CN-YunxiNeural】
- xiaoyi: 【zh-CN-XiaoyiNeural】
- xiaoxiao: 【zh-CN-XiaoxiaoNeural】
- yunjian: 【zh-CN-YunjianNeural】
- yunxia: 【zh-CN-YunxiaNeural】
- yunyang: 【zh-CN-YunyangNeural】
- LIAONING-xiaobei: 【zh-CN-liaoning-XiaobeiNeural】
- SHNAXI-xiaoni: 【zh-CN-shaanxi-XiaoniNeural】
- HK-xiaojia: 【zh-HK-HiuGaaiNeural】
- HK-chenman: 【zh-HK-HiuMaanNeural】
- HK-yunlong: 【zh-HK-WanLungNeural】
- TW-xiaochen: 【zh-TW-HsiaoChenNeural】
- TW-xiaoyu: 【zh-TW-HsiaoYuNeural】
- TW-yunzhe: 【zh-TW-YunJheNeural】
