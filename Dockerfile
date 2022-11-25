FROM golang:1.18.3-bullseye AS build-env

WORKDIR /go/src/github.com/imversed/imversed

RUN apt-get update -y
RUN apt-get install git -y

COPY . .
# install ignite
RUN curl https://get.ignite.com/cli! | bash

# checkout current version tag
# stash changes to checkout
RUN git stash
RUN git checkout v3.11

# build current version
RUN yes y | ignite chain build --release
RUN tar -zxvf release/imversed_linux_amd64.tar.gz
RUN mv imversed imversed_current

RUN cp imversed_current /usr/bin/imversed

# checkout next version tag or branch
# stash changes to checkout
#RUN git stash
#RUN git checkout pulling_updates_from_ethermint
#
## build next version
#RUN yes y | ignite chain build --release
#RUN tar -zxvf release/imversed_linux_amd64.tar.gz
#RUN mv imversed imversed_next

FROM golang:1.18.3-bullseye
WORKDIR /root

COPY --from=build-env /go/src/github.com/imversed/imversed/imversed_current /usr/bin/imversed
COPY --from=build-env /go/src/github.com/imversed/imversed/imversed_current /root/imversed_current
#COPY --from=build-env /go/src/github.com/imversed/imversed/imversed_next /root/imversed_next
COPY --from=build-env /go/src/github.com/imversed/imversed/imversed_current /root/imversed_next

RUN go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v0.1.0

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["cosmovisor", "start"]
