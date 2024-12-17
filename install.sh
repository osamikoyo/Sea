#bin/bash

go build -o bin/app cmd/sea/main.go &&
sudo mv bin/app /usr/bin