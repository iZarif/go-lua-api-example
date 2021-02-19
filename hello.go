package main

import "github.com/Shopify/go-lua"
import "math"

func l_sin(l *lua.State) int {
  l.PushNumber(math.Sin(lua.CheckNumber(l, 1)))

  return 1
}

func l_sqlSflow(l *lua.State) int {
  statement := lua.CheckString(l, 1)
  items := sql_sflow(statement)
  l.NewTable()

  for idx, item := range items {
    l.PushInteger(idx+1)
    l.PushUserData(item)
    l.RawSet(-3)
  }

  return 1
}

func main() {
  l := lua.NewState()
  lua.OpenLibraries(l)

  var myLibrary = []lua.RegistryFunction{
    {"sin", l_sin},
    {"sqlSflow", l_sqlSflow},
  }


  lua.NewLibrary(l, myLibrary)
  l.SetGlobal("my")

  if err := lua.DoFile(l, "hello.lua"); err != nil {
    panic(err)
  }
}
