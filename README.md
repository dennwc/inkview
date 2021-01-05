# Go SDK for Pocketbook

Unofficial Go SDK for Pocketbook based on libinkview.

Supports graphical user interfaces and CLI apps.

## Build a CLI app

Standard Go compiler should be able to cross-compile the binary
for the device (no need for SDK):

```
GOOS=linux GOARCH=arm GOARM=5 go build main.go
```

Note that some additional workarounds are necessary if you want to access
a network from your app. In this case you may still need SDK.

Although this binary will run on the device, you will need a third-party
application to actually see an output of you program (like
[pbterm](http://users.physik.fu-berlin.de/~jtt/PB/)).

The second option is to wrap the program into `RunCLI` - it will
emulate terminal output and write it to device display.

## Preparation - Build the Docker Image

For your convenience there is a `docker-compose.yaml`.

```bash
docker-compose build
```

will ctreate a `pb-go` container which can be used to compile a go program.

Adjust the `source` path in `docker-compose.yaml` to you needs.

With the current settings you can compile an sqlite test program:

```bash
docker-compose run --rm pb-go build go-sqlite.go
```

Alternatively, after building the image, `docker` can be used to compile without adjusting the file.

## Build an app with UI

To build your app or any example, run (requires Docker):

```bash
cd ./examples/sqlitetst/
docker run --rm -v $PWD:/app inkview_pb-go build go-sqlite.go
```

```bash
cd ./examples/devinfo/
docker run --rm -v $PWD:/app inkview_pb-go build main.go
```

You may also need to mount GOPATH to container to build your app:

```
docker run --rm -v $PWD:/app -v $GOPATH:/gopath inkview_pb-go build main.go
```

To run an binary, copy it into `applications/app-name.app` folder
on the device and it should appear in the applications list.

## Notes on networking

By default, device will try to shutdown network interface to save battery,
thus you will need to call SDK functions to keep device online (see `KeepNetwork`).

Also note that establishing TLS will require Go to read system
certificate pool that might take up to 30 sec on some devices and will
lead to TLS handshake timeouts. You will need to call `InitCerts` first
to fix the problem.

IPv6 is not enabled on some devices, thus a patch to Go DNS lib is required
to skip lookup on IPv6 address (SDK already includes the patch).
Similar problems may arise when trying to dial IPv6 directly.

## Notes on workdir

Application will have a working directory set to FS root, and not to
a parent directory.
To use relative paths properly change local dir to a binary's parent
directory: `os.Chdir(filepath.Dir(os.Args[0]))`.