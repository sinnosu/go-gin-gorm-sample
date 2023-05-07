FROM golang:1.19

ENV GO111MODULE on
WORKDIR /workspace

# コンテナ内の非ルートユーザの設定
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# 非ルートユーザー作成
# sudoコマンドの追加
# Go用の開発ツール
# go mod用のパーミッション変更 など
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt update \
    && apt install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME \
    && go install golang.org/x/tools/gopls@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/josharian/impl@latest \
    && sudo chmod -R a+rwX /go/pkg/ \
    && sudo curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64 \
    && sudo chmod +x /usr/local/bin/dbmate
