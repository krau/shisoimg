# ShiSoImg

轻松构建一个图片 API 服务, 具有一定的 [ManyACG](https://github.com/krau/ManyACG) API 兼容性

## 安装

从 [release](https://github.com/krau/shisoimg/releases) 页面下载最新的二进制文件, 解压后赋予权限.

## 使用

### 将图片添加到数据库

```shell
shisoimg add -d /path/to/image
```

这会遍历 `/path/to/image` 目录下的所有图片, 计算其 hash 值, 并将其添加到数据库中.

重复的图片文件只会添加一次.

### 添加直链规则

shisoimg 本身提供 `/images/:md5` 路由, 用于获取图片文件. 但同时可以自定义直链前缀规则, 在获取图片时会跳转到对应的直链.

例, 将 /path/to/image 目录下的所有图片使用 `https://cdn.example.com/` 作为前缀:

```shell
shisoimg rule add https://cdn.example.com /path/to/image
```

列出所有规则:

```shell
shisoimg rule list
```

删除规则:

```shell
shisoimg rule del /path/to/image
```

### 启动服务

```shell
shisoimg serve -a ":34180"
```

## API

- GET `/ping`
- GET `/random` 随机获取一张图片
- GET `/images/:md5` 获取图片文件

### ManyACG Compatibility API

- `/v1/artwork/random`
- `/v1/artwork/random/preview`
- `/v1/artwork/list`
- `/v1/artwork/:id` (use md5 as id)
