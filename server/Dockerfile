# 使用 golang:1.23 作为基础镜像
FROM golang:1.23.1 AS builder

# 设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 定义时区参数
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 设置工作目录
WORKDIR /build

# 拷贝 go.mod 和 go.sum 文件并下载依赖
COPY go.mod .
COPY go.sum .
RUN go mod download

# 拷贝项目代码并编译可执行文件
COPY . .
RUN go build -o main ./cmd

# 使用 scratch 作为最终镜像，减小镜像大小
FROM scratch

# 拷贝可执行文件和配置文件
COPY --from=builder /build/main /
COPY --from=builder /build/config/config.yml /config/config.yml

# 暴露端口
EXPOSE 9000

# 运行可执行文件
CMD ["/main"]
