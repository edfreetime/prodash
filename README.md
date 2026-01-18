
# ProDash 🚀  

**Simple TUI Project Dashboard**

ProDash is a lightweight Terminal UI application to manage and open your projects quickly from the terminal.  
Built with **Go + Bubble Tea**, it’s fast, minimal, and works across multiple operating systems.

---

## ✨ Features

- 📁 CRUD project management from TUI  
- 🔎 Path picker to avoid wrong directory input  
- 🚀 Open project directly in Neovim  
- 💾 Data stored as JSON  
- ⚡ Single binary, no dependencies  
- 🖥️ Linux / macOS / Windows support

---

## 🧩 Requirements

- Go 1.21+ (if you want to build from source)  
- Neovim installed and available in PATH

---

## 📦 Installation (Binary Release)

Download the binary from `/dist` folder according to your OS.

### Linux

```bash
sudo mv prodash-linux /usr/local/bin/prodash
chmod +x /usr/local/bin/prodash
```

### macOS

```bash
sudo mv prodash-mac /usr/local/bin/prodash
chmod +x /usr/local/bin/prodash
```

### Windows

Add `prodash.exe` to your PATH or run it directly.

## 🛠 Build From Source

Build for current OS

```bash
go build -o dist/prodash
```

Cross Compile

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o dist/prodash-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o dist/prodash-mac

# Windows
GOOS=windows GOARCH=amd64 go build -o dist/prodash.exe
```

## 🚀 Usage

Run the app:

```bash
prodash
```

## Keybindings

- ↑ / ↓ : Navigate projects
- Enter : Open in Neovim
- a : Add project
- e : Edit project
- d : Delete project
- q : Quit

## 📁 Data Location

Project data is stores as JSON:

```
~/.config/prodash/projects.json
```

Example:

```json
[
  {
    "name": "My Project",
    "path": "/home/user/projects/my-project"
  }
]
```
