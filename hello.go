package main

import "github.com/Shopify/go-lua"
import "math"

func my_sin(l *lua.State) int {
  l.PushNumber(math.Sin(lua.CheckNumber(l, 1)))

  return 1
}

func main() {
  l := lua.NewState()
  lua.OpenLibraries(l)

  var myLibrary = []lua.RegistryFunction{
    {"my_sin", my_sin},
  }

  l.Global("_G")
  lua.SetFunctions(l, myLibrary, 0)

  if err := lua.DoFile(l, "hello.lua"); err != nil {
    panic(err)
  }
}
