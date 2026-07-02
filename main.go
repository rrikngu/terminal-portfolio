package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type page int

const (
	techPage page = iota
	collectPage
)

type viewMode int

const (
	menuView viewMode = iota
	contentView
)

var (
	pinkStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4"))
	boldStyle = lipgloss.NewStyle().Bold(true)
)

const (
	taglineBold   = "  creative programmer displaying\n  her interests in fun ways."
	taglineNormal = "\n  building things with golang.\n"
)

// Project models an active repo entry in the portfolio list
type Project struct {
	Name        string
	Description string
	URL         string
}

// houses state and structural components
type model struct {
	mode          viewMode
	activePage    page
	headerArt     string
	projects      []Project
	selectedIndex int
}

func initialModel() model {
	const githubURL = "https://github.com/rrikngu"
	return model{
		mode:          menuView,
		activePage:    techPage,
		selectedIndex: 0,
		projects: []Project{
			{
				Name:        "pennypoca (python)",
				Description: "conversational ai agent for photocard market pricing",
				URL:         githubURL,
			},
			{
				Name:        "terminal portfolio (go)",
				Description: "the site you are currently viewing: a bubble tea tui portfolio",
				URL:         githubURL,
			},
			{
				Name:        "cupidity (c#)",
				Description: "a tactical simultaneous trading card game",
				URL:         githubURL,
			},
			{
				Name:        "dream loop (python)",
				Description: "ai agent fork for autonomous vehicle training",
				URL:         githubURL,
			},
		},
	}
}

func openURL(url string) tea.Cmd {
	return func() tea.Msg {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", url)
		case "darwin":
			cmd = exec.Command("open", url)
		default:
			cmd = exec.Command("xdg-open", url)
		}
		_ = cmd.Start()
		return nil
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) enterPage() (model, tea.Cmd) {
	m.mode = contentView
	return m, nil
}

func (m model) switchToPage(p page) (model, tea.Cmd) {
	m.activePage = p
	return m, nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "`":
			if m.mode == menuView {
				return m.enterPage()
			}
			if m.mode == contentView && m.activePage == collectPage {
				return m, openURL(m.projects[m.selectedIndex].URL)
			}

		case "b":
			if m.mode == contentView {
				m.mode = menuView
			}

		case "left":
			return m.switchToPage(techPage)

		case "right":
			return m.switchToPage(collectPage)

		case "up":
			if m.mode == contentView && m.activePage == collectPage && m.selectedIndex > 0 {
				m.selectedIndex--
			}

		case "down":
			if m.mode == contentView && m.activePage == collectPage && m.selectedIndex < len(m.projects)-1 {
				m.selectedIndex++
			}
		}
	}

	return m, nil
}

func (m model) renderNavBar() string {
	tab1 := "1. technical bio"
	tab2 := "2. about me"
	if m.activePage == techPage {
		tab1 = pinkStyle.Render("> 1. technical bio")
	} else {
		tab2 = pinkStyle.Render("> 2. about me")
	}
	return fmt.Sprintf(" %s      %s \n", tab1, tab2)
}

func (m model) renderHelp() string {
	if m.mode == menuView {
		return " [← → to select | ` to open | q to quit] \n"
	}
	if m.activePage == collectPage {
		return " [b to return home | ← → to switch | ↑ ↓ to scroll | ` to open link | q to quit]\n"
	}
	return " [b to return home | ← → to switch | q to quit]\n"
}

func (m model) renderPageContent() string {
	if m.activePage == techPage {
		s := boldStyle.Render("\n  TECHNICAL BACKGROUND & EXPERIENCE")
		s += "\n"	
		s += "  • tech stack:  golang, typescript/react, c#, sql, python\n"
		s += "  • systems:     linux, docker, tcp networking, aws, redis\n\n"
		s += "  focus on full-stack architecture and creative backend engineering.\n"
		s += "  building things with responsive frontends and asynchronus data pipelines.\n"
		return s
	}

	s := boldStyle.Render("\n  ABOUT ME")
	s += "\n"
	s += "  • likes:       creative programming, collecting, and learning\n"
	s += "  • working on:  ai pricing agent + terminal portfolio + more\n\n"
	s += boldStyle.Render("  PINNED PROJECTS (↑ ↓ to scroll, ` to open)")
	s += "\n"
	for i, project := range m.projects {
		pointer := " "
		if i == m.selectedIndex {
			pointer = ">"
		}
		line := fmt.Sprintf(" %s %s", pointer, project.Name)
		if i == m.selectedIndex {
			line = pinkStyle.Render(line)
			line += pinkStyle.Render(fmt.Sprintf("\n    └─ %s", project.Description))
		}
		s += line + "\n"
	}
	return s
}

func (m model) renderNameArt() string {
	return pinkStyle.Render(strings.TrimRight(m.headerArt, "\n"))
}

func (m model) renderHeader() string {
	s := m.renderNameArt()
	s += "\n\n" + boldStyle.Render(taglineBold)
	s += "\n" + taglineNormal
	return s
}

func (m model) View() string {
	var s string
	if m.mode == menuView {
		s = m.renderHeader()
	} else {
		s = m.renderNameArt() + "\n" + m.renderPageContent()
	}

	s += "\n----------------------------------------------------------\n"
	s += " " + m.renderNavBar()
	s += " " + m.renderHelp()
	return s
}

func main() {
	bannerBytes, err := os.ReadFile("name.txt")
	var header string

	if err != nil {
		header = "kayla nguyen"
	} else {
		header = strings.TrimRight(strings.ReplaceAll(string(bannerBytes), "\r", ""), "\n")
	}

	m := initialModel()
	m.headerArt = header

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
