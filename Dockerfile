FROM golang:latest
RUN apt update

# Gitの補完を設定
RUN echo "source /usr/share/bash-completion/completions/git" >> ~/.bashrc

# GoのModule ModeをONにする
ENV GO111MODULE on

# ワーキングディレクトリを変更
WORKDIR /app

# jqコマンドをインストール
# jsonの整形、検索に用いる
RUN apt -y install jq

# goの各種ツールをインストール
RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/stamblerre/gocode@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/cweill/gotests/gotests@latest