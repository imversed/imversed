# Installation
## Install Go

> Imversed is built using Go (opens new window)version 1.19+

```text
go version
```

> If the imversed: command not found error message is returned, confirm that your [GOPATH](https://golang.org/doc/gopath_code#GOPATH) is correctly configured by running the following command:
> ```go
> export PATH=$PATH:$(go env GOPATH)/bin
> ```

## Install Binaries

> The latest Imversed [version](https://github.com/imversed/imversed/releases) is imversed v10.0.0-rc1

### GitHub
Clone and build Imversed using git:

```text
git clone https://github.com/imversed/imversed.git
cd imversed
git fetch
git checkout <tag>
make install
```

`<tag>` refers to a released tag from the tags [page](https://github.com/imversed/imversed/tags).

After installation is done, check that the imversed binaries have been successfully installed:

```text
imversed version
```

### Docker
You can build Imversed using Docker by running:

```text
make build-docker
```

The command above will create a docker container: `tharsishq/imversed:latest`. Now you can run `imversed` in the container.

```text
docker run -it -p 26657:26657 -p 26656:26656 -v ~/.imversed/:/root/.imversed tharsishq/imversed:latest imversed version

# To initialize
# docker run -it -p 26657:26657 -p 26656:26656 -v ~/.imversed/:/root/.imversed tharsishq/imversed:latest imversed init test-chain --chain-id test_9000-2

# To run
# docker run -it -p 26657:26657 -p 26656:26656 -v ~/.imversed/:/root/.imversed tharsishq/imversed:latest imversed start
```
