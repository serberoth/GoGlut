package main

// //cgo CFLAGS: -Wnodeprecated-declarations
// #cgo CFLAGS: -g
// #cgo LDFLAGS: -F/System/Library/Frameworks
// #cgo LDFLAGS: -framework OpenGL
// #cgo LDFLAGS: -framework GLUT
// #include <stdlib.h>
// #include <GLUT/glut.h>
//
// typedef void(*func_void)();
// typedef void(*func_vis)(int);
// typedef void(*func_key)(unsigned char, int, int);
// typedef void(*func_mouse)(int, int, int, int);
// typedef void(*func_motion)(int, int);
// typedef void(*func_special)(int, int, int);
import "C"
import "unsafe"

import "fmt"
import "os"

const (
  DO_ACCUM = C.GLUT_ACCUM
  DO_ALPHA = C.GLUT_ALPHA
  DO_DEPTH = C.GLUT_DEPTH
  DO_STENCIL = C.GLUT_STENCIL
  DO_MULTISAMPLE = C.GLUT_MULTISAMPLE
)

type GlutWindow struct {
  Name string
  Options uint
  DisplayString string
  Width int
  Height int

  Render func()
  Idle func()
  // Key func(byte, int, int)
  // Mouse(int, int, int, int)
}

var display func()
var visible func(C.int)
var idle func()
var keyboard func(C.uint, C.int, C.int)
var mouse func(C.int, C.int, C.int, C.int)
var motion func(C.int, C.int)
var special func(C.int, C.int, C.int)

func (win GlutWindow) Init() bool {
  argc := C.int(0)
  C.glutInit((*C.int)(&argc), nil)
  fmt.Printf("GLUT Successfully Initialized\n")
  // C.glutInitDisplayMode(C.GLUT_RGBA | C.GLUT_DOUBLE | C.GLUT_ALPHA | C.GLUT_DEPTH | C.GLUT_STENCIL | C.GLUT_MULTISAMPLE)
  C.glutInitDisplayMode(C.GLUT_RGBA | C.GLUT_DOUBLE | C.uint(win.Options))
  fmt.Printf("GLUT Display Mode initialized\n");
  // var dispStr := C.CString("samples stencil>=2 rgba double depth");
  dispStr := C.CString(win.DisplayString)
  // defer C.free(unsafe.Pointer(dispStr))
  C.glutInitDisplayString(dispStr)
  fmt.Printf("GLUT Display String initialized\n")
  // C.glutInitWindowSize(1024, 768)
  C.glutInitWindowSize(C.int(win.Width), C.int(win.Height))
  fmt.Printf("GLUT Window Size set %d x %d\n", win.Width, win.Height)

  // var nameStr := C.CString(name)
  nameStr := C.CString(win.Name)
  // defer C.free(unsafe.Pointer(nameStr))
  C.glutCreateWindow(nameStr)

  fmt.Printf("GLUT Window Initialized: %d x %d (%d-r, %d-g, %d-b, %d-a) %d-d %d-s\n",
    C.glutGet(C.GLUT_WINDOW_WIDTH),
    C.glutGet(C.GLUT_WINDOW_HEIGHT),
    C.glutGet(C.GLUT_WINDOW_RED_SIZE),
    C.glutGet(C.GLUT_WINDOW_GREEN_SIZE),
    C.glutGet(C.GLUT_WINDOW_BLUE_SIZE),
    C.glutGet(C.GLUT_WINDOW_ALPHA_SIZE),
    C.glutGet(C.GLUT_WINDOW_DEPTH_SIZE),
    C.glutGet(C.GLUT_WINDOW_STENCIL_SIZE))

  display = func() { win.Render() }
  C.glutDisplayFunc(C.func_void(unsafe.Pointer(&display)))
  visible = func(flag C.int) {
    if flag == C.GLUT_VISIBLE {
      idle = win.Idle
      C.glutIdleFunc(C.func_void(unsafe.Pointer(&idle)))
    } else {
      C.glutIdleFunc(nil)
    }
  }
  C.glutVisibilityFunc(C.func_vis(unsafe.Pointer(&visible)))
  keyboard = keyboardImpl
  C.glutKeyboardFunc(C.func_key(unsafe.Pointer(&keyboard)))
  // mouse = mouseImpl
  // C.glutMouseFunc(C.func_mouse(unsafe.Pointer(&mouse)))
  // motion = motionImpl
  // C.glutMotionFunc(C.func_motion(unsafe.Pointer(&motion)))
  // special = specialImpl
  // C.glutSpecialFunc(C.func_special(unsafe.Pointer(&special)))

  return true
}

func (win GlutWindow) MainLoop() {
  C.glutMainLoop()
}

func keyboardImpl(c C.uint, x, y C.int) {
  if c == 27 {
    os.Exit(0)
  }

  C.glutPostRedisplay()
}

func mouseImpl(button, state, x, y C.int) {
}

func motionImpl(x, y C.int) {
}

func specialImpl(k, x, y C.int) {
  C.glutPostRedisplay()
}

