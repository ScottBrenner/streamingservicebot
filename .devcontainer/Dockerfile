# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.163.1/containers/go/.devcontainer/base.Dockerfile
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN go get -x github.com/turnage/graw
RUN go get -v github.com/cweill/gotests/gotests
RUN go get -v google.golang.org/api/youtube/v3
RUN go get -v github.com/zmb3/spotify

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
