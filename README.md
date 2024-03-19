# 🚀 go!launch - a highly customizable app launcher

![Banner](assets/banner.svg)

## ✨ Features

### 📜 Configurable/Scriptable with lua

- Set options via the config object as [described here](#configuration)
- Hook into the api to add custom functionalty (e.g. [Calculator](#calculator))

### 🖥 Fully customzable UI (yes FULLY)

Run your own frontend on the specified port and it will be embedded via iframe.

```lua
  local config = require("go-launch.config").generate()
  config.frontend.port = 6969
  return config
```

This means you can:

- bring your own frame work
- get hot reload for free (with `vite dev` or similar)

Get access to the backend via `app.iframe` module.
Minimal example:

```svelte
<script>
  import { LaunchApp, FuzzyFindDesktopEntry } from '$lib/go-launch/app.iframe';
  let apps = [], searchTerm = ""
  onMount(async ()=>{
    apps = await FuzzyFindDesktopEntry("")
  })

  $:{
    apps = await FuzzyFindDesktopEntry(searchTerm)
  }

</script>

<input bind:value={searchTerm} type="text" placeholder="Search for an app" />
<div>
  {#each apps as app}
    <div>
      <button on:click={()=>{LaunchApp(app.id)}}>
        {app.name}
      </button>
    </div>
  {/each}
</div>
```

This is great for developing your frontend and immediately seeing the changes, but once you're stable
you probably want to switch to a static build.

```lua
local config = require("go-launch.config").generate()
config.frontend.dist = "~/.config/go-launch/frontend/dist"
return config
```

A server running on config.frontend.port has higher priority than the static build.
If no server is running on the specified port or if `config.frontend.static` is set to true, the static build will be used.

## 🔧 Configuration

You can configure all kinds of things. Here are some examples:

```lua
local my_config = require("go-launch.config").generate()
local keymap = require("go-launch.keymap")
local api = require("go-launch.api")
config.window.max_height = 600
config.window.max_width = 800
config.keymaps = {
  keymap("<C-Space>", function() api.search("foo") end),
  keymap("<C-S-Space>", function() api.search("bar") end),
}
return my_config
```

## 🧩 Applets

You can add applets via the `config.applets` array.
To trigger them you can specify a prefix, a pattern and/or a keymap.

```lua
local config = require("go-launch.config").generate()
local applet = require("go-launch.applet")
local calculator = require("go-launch.builtin.applets").calculator
-- these are the default value
calculator.prefix = "+"
calculator.pattern = [[^(-?\d+(\.\d+)?)(\+|-|\*|\/)(\d+(\.\d+)?|\(-\d+(\.\d+)?\))$]]
calculator.keymap = "<C-c>"
config.applets = {calculator}
return config
```

### 📝 Examples

#### 🧮 Calculator

```lua
local config = require("go-launch.config").generate()
local applet = require("go-launch.applet")
local calculator = applet("Calculator",{})
config.applets = {calculator}
return config
```

## 🛣 Roadmap

### 🏁 Goals

- define and implement lua API for configuration
- add some applet examples (Calculator, List of arbitrary commands, etc.)
- finish the example UI

### 🚫 Non-Goals

- implement any non-core-functionality in Go
- implement any UI once the example UI is finshed
- Windows support
