// +build !android
// +build !sdl2
// +build sdl3

package nk

/*
#cgo CFLAGS: -Wno-implicit-function-declaration
#cgo windows LDFLAGS: -Wl,--allow-multiple-definition

#define NK_SDL3_RENDERER_IMPLEMENTATION
#define NK_IMPLEMENTATION

#include "nuklear.h"
*/
import "C"
import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

