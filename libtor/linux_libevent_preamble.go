// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build linux android

package libtor

/*
#cgo CFLAGS: -I${SRCDIR}/../libevent_config
#cgo CFLAGS: -I${SRCDIR}/../linux/libevent
#cgo CFLAGS: -I${SRCDIR}/../linux/libevent/compat
#cgo CFLAGS: -I${SRCDIR}/../linux/libevent/include
#cgo CFLAGS: -I${SRCDIR}/../linux/libevent/include/event2
#cgo windows LDFLAGS: -lbcrypt
*/
import "C"
