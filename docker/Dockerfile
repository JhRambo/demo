# 该指令表示使用 alpine 镜像作为基础镜像。alpine 是一个轻量级的 Linux 发行版。
FROM alpine

# 该指令设置当前工作目录为 /app/demo。后续的指令将在这个目录下执行。
WORKDIR /app/demo

# 该指令将当前目录下的所有文件复制到容器中的 /app/demo 目录。这里的 . 表示当前目录。
COPY . /app/demo

# 指定运行时环境变量
ENV GIN_MODE=release \
  PORT=8088

EXPOSE 8088

ENTRYPOINT ["./docker"]