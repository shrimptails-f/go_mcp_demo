FROM golang:1.24.0-bookworm

ENV TZ=Asia/Tokyo

# 環境変数を設定して、日本語UTF-8ロケールを使用
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
# 必要なパッケージのインストールとロケールの設定
# commitのコメントで日本語を扱いたいので追加。
RUN apt-get update && \
    apt-get install -y locales && \
    localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

# ブランチ名がbashで表示されるようにカスタマイズ
COPY ../.bashrc /root/.bashrc

ENV APP_ROOT /data
WORKDIR $APP_ROOT

RUN apt-get update && \
    apt-get install -y vim && \
    rm -rf /var/lib/apt/lists/*

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/rinchsan/gosimports/cmd/gosimports@latest