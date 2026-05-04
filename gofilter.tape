# Record from the repository root:
#   make demo-gif
# or:
#   vhs gofilter.tape
#
# Requires: vhs, ttyd, ffmpeg, go (see https://github.com/charmbracelet/vhs)
#
# Set Theme drives ttyd’s palette (what VHS records). Without it, GIFs look flat.
# Full list: `vhs themes` — use quotes for names with spaces, e.g. Set Theme "Rose Pine".
# Good vivid picks for this TUI: Catppuccin Mocha, TokyoNightStorm, Dracula, Snazzy, CyberPunk2077.

Output gofilter.gif

Require go

Set Shell "bash"
Set Theme "Catppuccin Mocha"
Set FontSize 18
Set Width 1040
Set Height 680
Set Padding 16
Set Framerate 30
Set PlaybackSpeed 1

Hide
Type "mkdir -p /tmp/gofilter-vhs-demo && go run ./scripts/generate_demo_png.go && cp fixtures/demo.png /tmp/gofilter-vhs-demo/demo.png && go build -o ./gofilter ./cmd/gofilter && clear"
Sleep 1s
Enter

Sleep 13s

Show

Env GOFILTER_INITIAL_DIR "/tmp/gofilter-vhs-demo"

Type "./gofilter"
Sleep 800ms
Enter

Sleep 2s

# Below `..`; pick demo.png
Down
Sleep 800ms
Enter

Sleep 2s

# First filter is gray
Enter

Sleep 2s

# Confirm output directory (default home)
Enter

Sleep 15s

Sleep 700ms
Enter

Sleep 1500ms
