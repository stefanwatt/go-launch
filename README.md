# 🚀 go-launch - A Highly Customizable App Launcher

<img style="width: 100%; margin: 2rem 0px; border-radius: 1rem;" src="https://i.imgur.com/HnWx6vz.png" alt="go-launch.png">

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
