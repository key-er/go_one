1- First get protoc

```
cd ~/go/src
go get -u github.com/golang/protobuf/protoc-gen-go

```

2- Then get latest release from https://developers.google.com/protocol-buffers/docs/downloads
for  osx from and read the readme.txt you get after unnzipping
```
cd protoc-3.6.1-osx-x86_64

cp -r include/google /usr/local/include/
cp -r bin/protoc /usr/local/bin/
```

3- make sure you have GOPATH set
` echo $$GOPATH`



4- Generate code now
mkdir generated_code (to save output)



```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

// my SRC_DIR is the pwd so used "." here
protoc -I=. --go_out=generated_code addressbook.proto

```