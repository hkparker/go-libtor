// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
//go:build windows
// +build windows

package libtor

/*
#cgo CFLAGS: -I${SRCDIR}/../tor_config
#cgo CFLAGS: -I${SRCDIR}/../windows/tor
#cgo CFLAGS: -I${SRCDIR}/../windows/tor/src
#cgo CFLAGS: -I${SRCDIR}/../windows/tor/src/core/or
#cgo CFLAGS: -I${SRCDIR}/../windows/tor/src/ext
#cgo CFLAGS: -I${SRCDIR}/../windows/tor/src/ext/trunnel
#cgo CFLAGS: -I${SRCDIR}/../windows/tor/src/feature/api

#cgo CFLAGS: -DED25519_CUSTOMRANDOM -DED25519_CUSTOMHASH -DED25519_SUFFIX=_donna

#cgo LDFLAGS: -lm
#cgo windows LDFLAGS: -lshlwapi
*/
import "C"
