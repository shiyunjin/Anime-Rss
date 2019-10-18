#动漫花园RSS 分类过滤
动漫花园的RSS对自动下载动漫产生了极大的方便,但是由于其父分类(id:2)RSS会包含子分类 季度全集(id:31) 导致经常会下载季度全集的动漫,而且会额外占用很多空间.
本项目通过对动漫花园的RSS进行过滤剔除子分类,方便接入各种自动种子下载工具.

##  编译
请参阅go程序的编译方法方便自行修改,如使用linux x86_64系统可直接前往[Releases](https://github.com/shiyunjin/Anime-Rss/releases)进行下载.

tips: 需要go 13以上版本

## 使用方法

``` bash
/opt/rss/anime_server -h
可选参数:
 -port 默认2333
 -rssCookie 动漫花园RSS COOKIE,可前往个人中心获取如: ":COOKIE:uid=******;rsspass=**7b47a43******86***ba*"

```

## systemd file

``` systemd
[Unit]
Description=Anime Rss

[Service]
ExecStart=/opt/rss/anime_server -rssCookie ":COOKIE:uid=******;rsspass=**7b47a43******86***ba*"
Restart=always
RestartSec=20

[Install]
WantedBy=network.target
```
