package main

import "github.com/Shopify/go-lua"
import "math"

type item_t struct {
  x int
  y int
  address string
}

func l_sin(l *lua.State) int {
  l.PushNumber(math.Sin(lua.CheckNumber(l, 1)))

  return 1
}

func l_isItem(l *lua.State) int {
  if (lua.TestUserData(l, 1, "my.item") != nil) {
    l.PushBoolean(true)
  } else {
    l.PushBoolean(false)
  }

  return 1
}

func l_pushItem(l *lua.State, item item_t) {
  l.PushUserData(item)
  lua.SetMetaTableNamed(l, "my.item")
}

func l_sqlSflow(l *lua.State) int {
  //statement := lua.CheckString(l, 1)
  // items := sql_sflow(statement)
  items := []item_t{{0, 0, "add1"}, {0, 1, "add2"}, {0, 3, "add3"}}
  l.NewTable()

  for idx, item := range items {
    l.PushInteger(idx+1)
    l_pushItem(l, item)
    l.RawSet(-3)
  }

  return 1
}

func l_getItemAddress(l *lua.State) int {
  item := lua.CheckUserData(l, 1, "my.item").(item_t)

  l.PushString(item.address)

  return 1
}


func main() {
  l := lua.NewState()
  lua.OpenLibraries(l)
  lua.NewMetaTable(l, "my.item")

  myLibrary := []lua.RegistryFunction{
    {"sin", l_sin},
    {"sqlSflow", l_sqlSflow},
    {"isItem", l_isItem},
    {"getItemAddress", l_getItemAddress},
  }


  lua.NewLibrary(l, myLibrary)
  l.SetGlobal("my")

  if err := lua.DoFile(l, "hello.lua"); err != nil {
    panic(err)
  }
}
