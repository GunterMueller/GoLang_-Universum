export GOVERSION=1.7.4
export GOOS=linux
arch=$(uname -m)
if [ "$arch" = "x86_64" ]; then
  export GOARCH=amd64
else
  echo "works only for x86_64"; exit 1
fi
export GOSHA256=47fda42e46b4c3ec93fa5d4d4cc6a748aa3f9411a2a2b7e08e3a6d80d753ec8b
export GOTGZ=go$GOVERSION.$GOOS-$GOARCH.tar.gz
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
if [ $UID = 0 ]; then
  export GOSRC=$GOROOT/src
else
  export GO=$HOME/GoLangWork
  export GOPATH=$GO
  export GOSRC=$GO/src
  export GOBIN=$GO/bin
  mkdir -p $GOSRC $GO/pkg $GOBIN
  export PATH=$PATH:$GOBIN
fi
