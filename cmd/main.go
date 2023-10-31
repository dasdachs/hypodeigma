package main

import (
	"fmt"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type model int
type tickMsg struct{}

func (m model) Init() bubbletea.Cmd {
	return bubbletea.Tick(tickMsg{}, 0, 0)
}

func (m model) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	switch msg.(type) {
	case tickMsg:
		return m + 1, nil
	default:
		return m, nil
	}
}

func (m model) View() string {
	return fmt.Sprintf("Count: %d", m)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "This is an example app",
	Long:  `This is a longer description of the example app`,
	Run: func(cmd *cobra.Command, args []string) {
		p := bubbletea.NewProgram(model(0))
		if err := p.Start(); err != nil {
			fmt.Println("Could not start program:", err)
		}
	},
}

func main() {
	rootCmd.Execute()
}
