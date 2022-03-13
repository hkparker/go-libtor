// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
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
#cgo CFLAGS: -I${SRCDIR}/../windows/libevent/include/event2
#cgo CFLAGS: -I${SRCDIR}/../windows/openssl/include/openssl
#cgo CFLAGS: -I${SRCDIR}/../windows/zlib

#cgo CFLAGS: -DED25519_CUSTOMRANDOM -DED25519_CUSTOMHASH -DED25519_SUFFIX=_donna

#cgo LDFLAGS: -lm
*/
import "C"
