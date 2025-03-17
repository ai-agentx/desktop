# desktop

**desktop** provides a windows desktop application using Wails that connects to the
rest api of server.

## Installation

### Clone repository

```bash
git clone https://github.com/ai-agentx/desktop.git
cd desktop
```

### Install dependencies

```bash
apt update
apt install -y libgtk-3-dev libwebkit2gtk-4.0-dev pkg-config build-essential

pushd frontend
npm install
popd
```

### Build application

```bash
wails build -platform windows/amd64
wails build -platform darwin/universal
wails build -platform linux/amd64
```

## Reference

- [wails](https://github.com/wailsapp/wails)
