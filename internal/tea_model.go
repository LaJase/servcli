package internal

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/inancgumus/screen"
)

type status int
type sshFinishedMsg struct{ err error }

const (
	entities status = iota
	servers
)

const divisor = 3

var CfgGlobal ServerConfig

/* STYLING */
var (
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 1, 4)
	columnStyle   = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	lists    []list.Model
	focused  status
	loaded   bool
	quitting bool
	choice   string
	err      error
	command  bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.lists[m.focused].FilterState() == list.Filtering {
			break
		}

		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "left", "h":
			m.Prev()
		case "right", "l":
			m.Next()
		case "enter":
			item, ok := m.lists[servers].SelectedItem().(item)
			if ok {
				m.choice = string(item.title)
				return m, runSshCommand(generateOutput(m.choice))
			}
			// return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		columnStyle.Width(msg.Width / divisor)
		focusedStyle.Width(msg.Width / divisor)
		m.lists[entities].SetSize(msg.Width, 2*msg.Height/divisor)
		m.lists[servers].SetSize(msg.Width, 2*msg.Height/divisor)

		m.loaded = true
	case sshFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			m.quitting = true
			return m, tea.Quit
			// return m, nil
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	m.lists[m.focused].SetShowHelp(true)

	if m.focused == entities {
		selectedItem := m.lists[entities].SelectedItem().(item)
		alcSelected := selectedItem.title

		// Init servers list
		serverItems := []list.Item{}
		for _, group := range CfgGlobal.ServerList[alcSelected].Elements {
			for _, server := range group {
				serverItems = append(serverItems, item{title: server.Name, desc: getDescription(server)})
			}
		}
		m.lists[servers].SetItems(serverItems)
	}

	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		if m.choice != "" {
			return quitTextStyle.Render(generateOutput(m.choice))
		}
		return quitTextStyle.Render("No server has been chosen... Bye")
	}
	if m.command {
		return quitTextStyle.Render("No server has been chosen... Bye")
	}
	if m.loaded {
		airlineView := m.lists[entities].View()
		serverView := m.lists[servers].View()

		switch m.focused {
		case servers:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				columnStyle.Render(airlineView),
				focusedStyle.Render(serverView),
			)
		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				focusedStyle.Render(airlineView),
				columnStyle.Render(serverView),
			)
		}
	} else {
		return quitTextStyle.Render("Loading...")
	}

}

func (m *Model) Prev() {
	m.lists[m.focused].SetShowHelp(false)
	if m.focused == servers {
		m.focused = entities
	}
}

func (m *Model) Next() {
	m.lists[m.focused].SetShowHelp(false)
	if m.focused == entities {
		m.focused = servers
	}
}

func (m *Model) InitLists() {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList}

	// Init airlines list
	m.lists[entities].Title = "Entity"
	items := []list.Item{}
	serverItems := []list.Item{}
	for key, value := range CfgGlobal.ServerList {
		items = append(items, item{title: key, desc: value.Description})
	}
	m.lists[entities].SetItems(items)

	// Init servers list
	for _, group := range CfgGlobal.ServerList["Entity1"].Elements {
		for _, server := range group {
			serverItems = append(serverItems, item{title: server.Name, desc: getDescription(server)})
		}
	}
	m.lists[servers].Title = "Servers"
	m.lists[servers].SetItems(serverItems)
}

func getDescription(server Server) string {
	if server.IsAWS {
		return "AWS"
	}
	return "OLD"
}

func generateOutput(serverName string) string {
	return fmt.Sprintf(CfgGlobal.SshCommand, serverName)
}

func runSshCommand(command string) tea.Cmd {
	screen.Clear()
	screen.MoveTopLeft()
	parts := strings.Fields(command)
	c := exec.Command(parts[0], parts[1:]...) //nolint:gosec
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return sshFinishedMsg{err}
	})
}
