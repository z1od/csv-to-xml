build_mac:
	fyne package -os darwin --release true

build_windows:
	GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" fyne package --release true