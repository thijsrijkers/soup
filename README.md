## Soup
Soup is a lightweight 3D engine written in Go using Ebiten. It allows you to create simple first-person 3D environments in a classic “Doom” style.

## Map generation
In the program, the map is represented as a 2D array of integers. In this array, a 1 represents a wall, while a 0 represents a walkable area. For example, a simple map could be:

```
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
```
## Usage
Run the engine:
```bash
go run ./main.go
```

Then control the player using:
- `W` – Move forward  
- `S` – Move backward  
- `A` – Rotate left  
- `D` – Rotate right

## Requirements
- **Go 1.21+**
- **Ebiten v2** (`github.com/hajimehoshi/ebiten/v2`)
- **Operating System:**
  - Linux / WSL / Windows / macOS
- **Linux / WSL additional dependencies** (for X11 / OpenGL support):

```bash
sudo apt update
sudo apt install -y \
  libx11-dev \
  libxrandr-dev \
  libxinerama-dev \
  libxcursor-dev \
  libxi-dev \
  libxxf86vm-dev \
  libgl1-mesa-dev \
  libglu1-mesa-dev
```

## Installation
Clone the repository:
```bash
git clone https://github.com/thijsrijkers/soup.git
cd soup
go mod tidy
```
