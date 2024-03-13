<style>
  .tech-logo {
    max-height: 7rem;
    max-width: 7rem;
    padding: 0px 1rem;
  }
  .plus-icon {
    height: 3rem;
  }
</style>

# 🚀 go-launch - A Highly Customizable App Launcher

<div style="display:flex; justify-content:center; width:100%; align-items:center;'>
  <img style="max-height:7rem; max-width:7rem; padding: 0px 1rem;" src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" alt="golang.svg">
  <svg fill="#fff" stroke="#fff" class="plus-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>plus</title><path d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z" /></svg>
<svg style="height:4rem; max-width:7rem; padding: 0px 1rem;" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="100" height="100" viewBox="0 0 48 48">
<path fill="#E65100" d="M41,5H7l3,34l14,4l14-4L41,5L41,5z"></path><path fill="#FF6D00" d="M24 8L24 39.9 35.2 36.7 37.7 8z"></path><path fill="#FFF" d="M24,25v-4h8.6l-0.7,11.5L24,35.1v-4.2l4.1-1.4l0.3-4.5H24z M32.9,17l0.3-4H24v4H32.9z"></path><path fill="#EEE" d="M24,30.9v4.2l-7.9-2.6L15.7,27h4l0.2,2.5L24,30.9z M19.1,17H24v-4h-9.1l0.7,12H24v-4h-4.6L19.1,17z"></path>
</svg>
<svg fill="#fff" stroke="#fff" class="plus-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>plus</title><path d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z" /></svg>
<svg style="height:4rem; max-width:7rem; padding: 0px 1rem;" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="100" height="100" viewBox="0 0 48 48">
<path fill="#0277BD" d="M41,5H7l3,34l14,4l14-4L41,5L41,5z"></path><path fill="#039BE5" d="M24 8L24 39.9 35.2 36.7 37.7 8z"></path><path fill="#FFF" d="M33.1 13L24 13 24 17 28.9 17 28.6 21 24 21 24 25 28.4 25 28.1 29.5 24 30.9 24 35.1 31.9 32.5 32.6 21 32.6 21z"></path><path fill="#EEE" d="M24,13v4h-8.9l-0.3-4H24z M19.4,21l0.2,4H24v-4H19.4z M19.8,27h-4l0.3,5.5l7.9,2.6v-4.2l-4.1-1.4L19.8,27z"></path>
</svg>
<svg fill="#fff" stroke="#fff" class="plus-icon" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><title>plus</title><path d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z" /></svg>
<svg style="height:4rem; max-width:7rem; padding: 0px 1rem;" xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="100" height="100" viewBox="0 0 48 48">
<path fill="#f7df1e" d="M6,42V6h36v36H6z"></path><path fill="#000001" d="M29.538,32.947c0.692,1.124,1.444,2.201,3.037,2.201c1.338,0,2.04-0.665,2.04-1.585 c0-1.101-0.726-1.492-2.198-2.133l-0.807-0.344c-2.329-0.988-3.878-2.226-3.878-4.841c0-2.41,1.845-4.244,4.728-4.244 c2.053,0,3.528,0.711,4.592,2.573l-2.514,1.607c-0.553-0.988-1.151-1.377-2.078-1.377c-0.946,0-1.545,0.597-1.545,1.377 c0,0.964,0.6,1.354,1.985,1.951l0.807,0.344C36.452,29.645,38,30.839,38,33.523C38,36.415,35.716,38,32.65,38 c-2.999,0-4.702-1.505-5.65-3.368L29.538,32.947z M17.952,33.029c0.506,0.906,1.275,1.603,2.381,1.603 c1.058,0,1.667-0.418,1.667-2.043V22h3.333v11.101c0,3.367-1.953,4.899-4.805,4.899c-2.577,0-4.437-1.746-5.195-3.368 L17.952,33.029z"></path>
</svg>
</div>
<img style="width: 100%; margin: 2rem 0px; border-radius: 1rem;" src="https://i.imgur.com/WnwDfWz.png" alt="go-launch.png">

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Customization](#customization)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Fuzzy-Finding:** Quickly find and launch your applications with fuzzy-finding
- **Keyboard Navigation:** Navigate the search results with the arrow keys like in rofi
- **Highly Customizable UI:** Thanks to web technologies, tweak the UI to your heart's content

<img style="width: 100%; border-radius: 1rem; " src="https://i.imgur.com/2ezvb22.gif" alt="go-launch.gif">

## Getting Started

### Prerequisites

- [wails-cli](https://wails.io/docs/gettingstarted/installation)
- Ensure `gtk-launch` is installed on your system, typically available through the `gtk3` package.

### Installation

1. Clone the repository:
   ```bash
   git clone [your-repo-link]
   ```
2. Navigate to the project directory and build the project:
   ```bash
   cd your-app-launcher
   go build
   ```

## Customization

- The UI, powered by Svelte, is designed for easy customization. Dive into the UI code and make it truly yours.
- Encouraged to compile the project yourself to replace the existing UI with your creative touch.

## Roadmap

- [ ] Lua-based configuration for ultimate flexibility.
- [ ] Plugin API to extend functionality with Lua scripts.
- [ ] Introduction of basic applets, like a calculator and a customizable command list.

## Contributing

We welcome contributions of all forms. Check out our contributing guidelines for more information.

## License

Distribrted under the MIT License. See `LICENSE` for more information.
