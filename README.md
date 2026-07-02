# terminal portfolio

keyboard-driven terminal user interface (TUI) portfolio that brings a personal resume straight into the terminal. built with **go** and the **bubble tea** ecosystem, incorporating ASCII text art.


## features

*   **keyboard navigation:** mouse-free interface driven by keyboard hotkeys.
*   **multi-page view:** swap between a technical overview and my personal projects.
*   **project carousel:** interactive project selector wiith embedded links.


## interface previews

<p align="center">
  <img src="docs/images/demo-gif.gif" alt="terminal portfolio demo" width="60%">
</p>

1. main hub: landing view showcasing custom ASCII text art.
2. technical profile: core technical stack capabilities and architecture focuses.
3. pinned projects: carousel module allowing viewers to navigate through projects with immediate sub-details.

<p align="center">
  <img src="docs/images/home_menu.png" alt="main hub view" width="30%" style="border: 1px solid #333; margin: 10px;">
  <img src="docs/images/tech_bio.png" alt="technical bio view" width="30%" style="border: 1px solid #333; margin: 10px;">
  <img src="docs/images/projects.png" alt="project showcase view" width="30%" style="border: 1px solid #333; margin: 10px;">
</p>

## terminal controls

| command key | system action performance |
| :--- | :--- |
| `←` / `→` | toggle navigation options |
| `` ` `` (backtick) | drill down into content tabs / fire browser command for selected app link |
| `↑` / `↓` | drive the cursor up and down through the pinned project layout |
| `b` | return back to the main hub |
| `q` / `ctrl+c` | exit and terminate the app |