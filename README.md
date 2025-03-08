# Git Emoji Terminal Selector

A terminal-based Git emoji selector that appears with Ctrl+E, making it easy to add emojis to your Git commits.

## Features

- Quick access with Ctrl+E keyboard shortcut
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

1. Install the package:
```bash
go install github.com/casper9429-kth/git-emoji-terminal/cmd/daemon@latest
```

2. Add the keyboard shortcut to your `.bashrc`:
```bash
# Git emoji selector keybinding
bind -x '"\C-e": "/home/casper/go/bin/gitemoji-daemon"'
```

3. Reload your shell configuration:
```bash
source ~/.bashrc
```

## Usage

1. Press Ctrl+E anywhere in your terminal to open the emoji picker
2. Use arrow keys to navigate or type to filter emojis
3. Press Enter to select an emoji (it's automatically copied to your clipboard)
4. Press Esc or Ctrl+C to cancel

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

## Tips for Git Users

1. Use the emoji at the start of your commit messages:
```bash
git commit -m "✨ Add new feature"
```

2. Quick search examples:
   - Type "fix" to quickly find the 🐛 bug fix emoji
   - Type "feat" to find the ✨ feature emoji
   - Type the emoji itself to find it directly

## Contributing

Feel free to contribute by opening issues or submitting pull requests.

## License

MIT