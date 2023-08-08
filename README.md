# tpb-tui

![alt](screenshot.png)

`tpb-tui` is a terminal user interface for The Pirate Bay. It lets you search for torrents and display them in a table. Selecting a torrent will open the magnet link in your default torrent client.

## Features

- Search for torrents directly from your terminal.
- Display torrents in a neat table format.
- Open magnet links in the default torrent client.
- User-friendly controls: Enter to select, Esc to go back or exit.

## Installation

Make sure you have Go installed (v1.20 or newer is recommended).

To install `tpb-tui`, run:

```bash
go get -u github.com/jackbillstrom/tpb-tui
```

## Usage

After installation, run the application using:

```bash
tpb-tui
```

Follow the on-screen prompts to search for torrents and navigate through the results.

## Contributing

If you'd like to contribute, please fork the repository and make changes as you'd like. Pull requests are warmly welcome.

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Add some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Dependencies

tpb-tui relies on several third-party Go packages. Here are some of the major dependencies:

[bubbletea/examples/table/main.go](https://github.com/charmbracelet/bubbletea/blob/master/examples/table/main.go) for the table UI.

[github.com/charmbracelet/bubbles](github.com/charmbracelet/bubbles) for UI components.

[github.com/charmbracelet/bubbletea](github.com/charmbracelet/bubbletea) for the app's model-view-update architecture.

[github.com/charmbracelet/lipgloss](github.com/charmbracelet/lipgloss) for styling.

## Acknowledgements

A special thanks to the developers of all the Go packages used in this project.
