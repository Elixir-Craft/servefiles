VERSION=v$(cat VERSION | tr -d '[:space:]')

GOOS=linux GOARCH=amd64 go build -o ./build/servefiles-linux-amd64-$VERSION ./servefiles
GOOS=linux GOARCH=arm64 go build -o ./build/servefiles-linux-arm64-$VERSION ./servefiles
# build for windows
GOOS=windows GOARCH=amd64 go build -o ./build/servefiles-windows-amd64-$VERSION.exe ./servefiles

# build for mac
GOOS=darwin GOARCH=amd64 go build -o ./build/servefiles-darwin-amd64-$VERSION ./servefiles
GOOS=darwin GOARCH=arm64 go build -o ./build/servefiles-darwin-arm64-$VERSION ./servefiles


