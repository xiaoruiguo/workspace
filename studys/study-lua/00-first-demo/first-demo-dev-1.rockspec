package = "first-demo"
version = "dev-1"
source = {
   url = "git+https://github.com/lijiaocn/workspace.git"
}
description = {
   homepage = "https://github.com/lijiaocn/workspace",
   license = "*** please specify a license ***"
}
build = {
   type = "builtin",
   modules = {
      ["01-hello-world"] = "01-hello-world.lua"
   }
}
