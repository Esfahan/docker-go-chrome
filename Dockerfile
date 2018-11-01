From centos:7

ARG GO_VER=1.11
ARG CHROME_DRIVER_VER=2.43

RUN yum update -y

########
# go
########
WORKDIR /usr/local/src
RUN yum install -y git
RUN curl -O https://dl.google.com/go/go${GO_VER}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go${GO_VER}.linux-amd64.tar.gz

ENV GOPATH=/code
ENV PATH=$PATH:$GOPATH/bin
ENV PATH=$PATH:/usr/local/go/bin


########################
# chromedriver, chrome
########################
# chromedriver
RUN yum install -y unzip
RUN curl -O https://chromedriver.storage.googleapis.com/${CHROME_DRIVER_VER}/chromedriver_linux64.zip
RUN unzip chromedriver_linux64.zip
RUN mv chromedriver /usr/local/bin/

# Dependencies
RUN yum install -y libX11 \
                   GConf2 \
                   fontconfig

# chrome
ADD files/etc/yum.repos.d/google-chrome.repo /etc/yum.repos.d/google-chrome.repo
RUN yum install -y google-chrome-stable \
                   libOSMesa

# fonts
RUN yum install -y google-noto-cjk-fonts \
                   ipa-gothic-fonts
