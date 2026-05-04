
# Go Filter

Go Filter is a TUI (Terminal User Interface) tool to transform images with filters like grey-scale, inverted colors, and more.

## How to start

```bash
# Clone the repository
git clone https://github.com/OrlandoRomo/go-filter.git

# Change directory to go-filter
cd go-filter

# Install dependencies
go mod tidy

# Build project
cd cmd/gofilter && go build -o gofilter .
```

Move the binary to your local environment:
- Linux/Mac: `sudo mv gofilter /usr/local/bin`

## Usage

To use `gofilter`, open a terminal and run:

```bash
gofilter
```

The application will launch a TUI (Terminal User Interface) with the following flow:

### 1. File Selection
- A file picker will appear showing directories and image files
- **Supported formats**: `.png`, `.jpg`, `.jpeg`
- `.gif` and `.webp` are NOT supported
- Use arrow keys to navigate, `enter` to select a file or directory

### 2. Filter Selection
After selecting an image, you'll see a list of available filters:
- **gray** - Gray scale filter
- **sepia** - Sepia filter
- **negative** - Negative/inverted colors filter
- **sketch** - Sketch filter
- **red** - Red scale filter
- **green** - Green scale filter
- **blue** - Blue scale filter
- **mirror** - Mirror/flip filter
- **sharp** - Sharp filter
- **blur** - Blur filter

Use arrow keys (`up`/`down` or `k`/`j`) to navigate, `enter` to select.

### 3. Output Path Configuration
- Default output path is your home directory
- You can edit the path to change where the filtered image will be saved
- Press `enter` to confirm and start processing, `esc` to go back

### 4. Processing
- A progress bar will show the processing status
- The image is being transformed in the background

### 5. Success
- A success message will display the full output path and filename
- Press any key to exit

## Examples

| Filter       | Result |
| ------------- |:-------------:|
| `gray` | <img src="https://user-images.githubusercontent.com/34588445/133297171-c9b00477-4a1e-49ad-8d6d-0b730ba0285f.jpg" width="150" height="150"> |
| `negative` | <img src="https://user-images.githubusercontent.com/34588445/133297151-c4494112-7856-4c27-aae1-b07a2bd6b384.jpg" width="150" height="150"> |
| `red` | <img src="https://user-images.githubusercontent.com/34588445/133297177-714859ee-301c-429e-851a-dce40378e25c.jpg" width="150" height="150"> |
| `blue` | <img src="https://user-images.githubusercontent.com/34588445/133297188-843f51fe-f9d7-473d-9c54-ae7a1faf25f2.jpg" width="150" height="150"> |
| `green` | <img src="https://user-images.githubusercontent.com/34588445/133297128-16c7ad56-f2f6-4a8d-8684-d1c795177e5c.jpg" width="150" height="150"> |
| `mirror` | <img src="https://user-images.githubusercontent.com/34588445/133297175-8d2aa032-902d-4f6a-8459-2c52ad12148b.jpg" width="150" height="150"> |
| `sepia` | <img src="https://user-images.githubusercontent.com/34588445/133297141-c022155d-05ef-4162-a61e-4920509cad8e.jpg" width="150" height="150"> |
| `sketch` | <img src="https://user-images.githubusercontent.com/34588445/133297150-646feaaa-4126-46d7-aecc-2c3df93e28aa.jpg" width="150" height="150"> |
| `sharp` | <img src="https://user-images.githubusercontent.com/34588445/133297162-3abbd4b1-1d35-4997-bebc-05d61ab95cfa.jpg" width="150" height="150"> |

## Project Structure

```
gofilter/
├── cmd/
│   └── gofilter/
│       └── main.go              # Entry point
├── internal/
│   ├── filter/                  # Filter implementations
│   │   ├── filter.go           # Filter interface & Effect struct
│   │   ├── gray.go             # Gray filter
│   │   ├── sepia.go            # Sepia filter
│   │   ├── negative.go         # Negative filter
│   │   ├── sketch.go           # Sketch filter
│   │   ├── red.go              # Red filter
│   │   ├── green.go            # Green filter
│   │   ├── blue.go             # Blue filter
│   │   ├── mirror.go           # Mirror filter
│   │   ├── sharp.go            # Sharp filter
│   │   ├── blur.go             # Blur filter
│   │   └── list.go             # Filter list utility
│   ├── image/                  # Image reading/writing
│   │   ├── reader.go          # Image reader
│   │   └── writer.go          # Image writer
│   └── tui/                    # Bubble Tea TUI
│       ├── model.go            # Main TUI model
│       ├── filepicker.go       # File picker view
│       ├── filterlist.go       # Filter list view
│       ├── outputpath.go       # Output path view
│       ├── progress.go         # Progress bar view
│       └── success.go          # Success/error views
├── go.mod
└── README.md
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components (progress bar, text input)
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styles

## TODO
1. Add more filters
2. Add option to overwrite original file
3. Add batch processing support
4. Add image preview before/after filter application
