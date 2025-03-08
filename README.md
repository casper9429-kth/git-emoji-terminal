# Git Emoji Terminal Selector

A lightweight and user-friendly terminal-based Git emoji selector for your commit messages.

## Features

- Fast and responsive terminal UI
- Fuzzy search for emojis
- Keyboard navigation (arrow keys)
- Instant clipboard copy of selected emoji
- Filter emojis by name or symbol

## System Requirements

- Linux with X11 (X Window System)
- `xdotool` for keyboard simulation
- `xterm` for displaying the emoji selector popup
- `fzf` for the fuzzy finder interface

Install system dependencies:

```bash
# For Debian/Ubuntu:
sudo apt-get install xdotool xterm fzf

# For Fedora/RHEL:
sudo dnf install xdotool xterm fzf

# For Arch Linux:
sudo pacman -S xdotool xterm fzf
```

## Installation

```bash
go install github.com/casper/gitemoji/cmd/daemon@latest
```

## Usage

There are several ways to integrate this tool into your workflow:

### 1. Direct Usage

Simply run:
```bash
gitemoji-daemon
```

This will open the emoji picker. Use arrow keys to navigate, type to filter, and press Enter to select.
The selected emoji will be copied to your clipboard.

### 2. Shell Integration

Add this to your `.bashrc` or `.zshrc`:

```bash
# Bind Ctrl+E to open the emoji picker
bind '"\C-e": "gitemoji-daemon\n"'  # For Bash
# bindkey '^E' 'gitemoji-daemon\n'   # For Zsh
```

Now you can press Ctrl+E anywhere in your terminal to open the emoji picker.

### 3. Git Alias

Add this to your `.gitconfig`:

```ini
[alias]
    emoji = !gitemoji-daemon
```

Now you can use:
```bash
git emoji
```

## How It Works

1. Run the command or use your configured shortcut
2. A TUI (Terminal User Interface) appears with a list of Git emojis
3. Type to filter the list or use arrow keys to navigate
4. Press Enter to select an emoji (it's automatically copied to clipboard)
5. Press Esc or Ctrl+C to cancel

## Available Emojis

- ✨ (feat) - New feature
- 🐛 (fix) - Bug fix
- 📚 (docs) - Documentation
- 💎 (style) - Formatting/styling
- 📦 (refactor) - Code refactoring
- 🚀 (perf) - Performance improvement
- 🔍 (test) - Testing
- 🛠️ (build) - Build system
- ⚙️ (ci) - CI/CD
- ♻️ (chore) - Chores/maintenance
- ⏪ (revert) - Revert changes
- 🚧 (wip) - Work in progress
- 🎉 (init) - Initial commit
- 🔒 (security) - Security fixes
- 📦 (deps) - Dependencies
- 🌐 (i18n) - Internationalization
- ✏️ (typo) - Typo fix
- 🔀 (merge) - Merge
- ⚙️ (config) - Configuration changes

## Tips

1. For Git commits, select your emoji first, then type your commit message:
   ```bash
   git commit -m "✨ Add new feature"
   ```

2. Use the search feature by typing part of the emoji name:
   - Type "fix" to quickly find the 🐛 bug fix emoji
   - Type "feat" to find the ✨ feature emoji

3. The selected emoji is automatically copied to your clipboard, so you can immediately paste it anywhere

## Security Note

Since this app uses X11 keyboard hooks to monitor your typing, it has access to all keyboard input across applications. The code is open source and doesn't collect any data, but you should always review such applications carefully before running them.

## Customization

You can modify the list of emojis by editing the `GitEmojis` slice in `gitemoji.go`.

## Troubleshooting

- **No emoji picker appears**: Make sure you have xterm and fzf installed
- **Cannot grab keyboard**: You may need to run the daemon with elevated privileges
- **Wrong characters detected**: The keycode mapping may need adjustment for your keyboard layout

## License

MIT