package main

func main() {
  win := GlutWindow{}
  win.Name = "GO GLUT"
  win.Options = DO_ALPHA | DO_DEPTH | DO_STENCIL | DO_MULTISAMPLE
  win.DisplayString = "samples stencil>=2 rgba double depth"
  win.Width = 1024
  win.Height = 768

  win.Render = func() { }
  win.Idle = func() { }

  win.Init()

  win.MainLoop()
}

