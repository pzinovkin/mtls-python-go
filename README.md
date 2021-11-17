# Mutual TLS with gRPC between Python and Go services

Shell commands use [fish](https://fishshell.com/) syntax.

mTLS with Nginx terminating TLS
![nginx](./nginx.svg)

Install required tools:
```
brew install python3 go protobuf nginx cfssl tmux
```

Create and provision virtualenv
```
make initvenv
```

Generate certificates and keys
```
make certs
```

Activate virtualenv and start client's web app server:
```
. /venv/bin/activate.fish
./manage.py runserver
```

Start server in new terminal window:
```
go run server/server.go
```

Run single-process nginx in new terminal window:
```
nginx -p (pwd) -c nginx.conf
```

Or you can run server and nginx in on terminal with tmux:
```
tmux \
  new-session "go run server/server.go" \; \
  split-window -h "nginx -p (pwd) -c nginx.conf"
```