// go-libtor - Self-contained Tor from Go
// Copyright (c) 2018 Péter Szilágyi. All rights reserved.
// +build windows

package libtor

/*
#include <compat/sys/queue.h>
#include <../win32select.c>
*/
import "C"
