package main

// reference https://dev.to/patarapolw/maximize-not-fullscreen-your-desktop-application-4652

import (
	/*
		#cgo darwin LDFLAGS: -framework CoreGraphics
		#cgo linux pkg-config: x11

		#if defined(__APPLE__)
		#include <CoreGraphics/CGDisplayConfiguration.h>
		int display_width() {
		return CGDisplayPixelsWide(CGMainDisplayID());
		}
		int display_height() {
		return CGDisplayPixelsHigh(CGMainDisplayID());
		}
		#elif defined(_WIN32)
		#include <wtypes.h>
		int display_width() {
		RECT desktop;
		const HWND hDesktop = GetDesktopWindow();
		GetWindowRect(hDesktop, &desktop);
		return desktop.right;
		}
		int display_height() {
		RECT desktop;
		const HWND hDesktop = GetDesktopWindow();
		GetWindowRect(hDesktop, &desktop);
		return desktop.bottom;
		}
		#else
		#include <X11/Xlib.h>
		int display_width() {
		Display* d = XOpenDisplay(NULL);
		Screen*  s = DefaultScreenOfDisplay(d);
		return s->width;
		}
		int display_height() {
		Display* d = XOpenDisplay(NULL);
		Screen*  s = DefaultScreenOfDisplay(d);
		return s->height;
		}
		#endif
	*/
	"C"
)
import "runtime"

func GetFullscreenSize() (int, int) {
	width := int(C.display_width())
	height := int(C.display_height())

	// Current method of getting screen size in linux and windows makes it fall offscreen
	if runtime.GOOS == "linux" || runtime.GOOS == "windows" {
		width = width - 50
		height = height - 100
	}

	if width == 0 || height == 0 {
		width = 1024
		height = 768
	}

	return width, height
}
