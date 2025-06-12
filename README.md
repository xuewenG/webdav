# WebDAV

用于快速启动 WebDAV 服务，用法：

准备配置文件：

```yaml
# config.yaml
# 端口
port: 3002
# 请求前缀
prefix: /
# 文件目录
root_dir: /data
# 只读模式
read_only: true
# 用户列表
users:
  - username: test
    password: test-123
```

启动服务，使用 docker-compose：

```yaml
# docker-compose.yaml
services:
  webdav:
    image: ixuewen/webdav
    container_name: webdav
    restart: always
    ports:
      # 和配置文件中的端口号保持一致
      - 3002:3002
    volumes:
      - ./webdav/config.yaml:/app/config.yaml
      # 和配置文件中的文件目录保持一致
      - ./webdav/data:/data
```
