module github.com/cutlery47/gopong/client

go 1.23.0

require github.com/hajimehoshi/ebiten/v2 v2.7.8

require github.com/cutlery47/gopong/common v0.0.0

replace github.com/cutlery47/gopong/common => ../common

require (
	github.com/ebitengine/gomobile v0.0.0-20240518074828-e86332849895 // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/ebitengine/purego v0.7.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
)
