// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build windows

package libtor

/*
#cgo CFLAGS: -I${SRCDIR}/../libevent_config
#cgo CFLAGS: -I${SRCDIR}/../windows/libevent
#cgo CFLAGS: -I${SRCDIR}/../windows/libevent/compat
#cgo CFLAGS: -I${SRCDIR}/../windows/libevent/include
*/
import "C"
