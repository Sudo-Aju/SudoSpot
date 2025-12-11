package ui

import (
	"context"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zmb3/spotify/v2"
)

var (
	windowStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#1DB954")).
			Padding(1, 2).
			Width(60)

	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
	artistStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAAAA"))
	statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#1DB954"))
)

type Model struct {
	Client    *spotify.Client
	TrackName string
	Artist    string
	IsPlaying bool
	Progress  int
	Duration  int
	Error     string
}

func NewModel(client *spotify.Client) Model {
	return Model{
		Client:    client,
		TrackName: "Waiting for data...",
		Artist:    "Unknown Artist",
		IsPlaying: false,
	}
}

func (m Model) Init() tea.Cmd {
	return tickCmd()
}

type spotifyMsg *spotify.CurrentlyPlaying

type errMsg error

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.Error = msg.Error()
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case " ":
			if m.IsPlaying {
				return m, pause(m.Client)
			} else {
				return m, play(m.Client)
			}
		case "left":
			newPos := m.Progress - 10000
			if newPos < 0 {
				newPos = 0
			}
			return m, seek(m.Client, newPos)
		case "right":
			newPos := m.Progress + 10000
			if newPos > m.Duration {
				newPos = m.Duration - 100
			}
			return m, seek(m.Client, newPos)

		case "n", ">":
			return m, next(m.Client)

		case "p", "<":
			return m, previous(m.Client)
		}

	case tickMsg:
		return m, tea.Batch(
			checkSpotify(m.Client),
			tickCmd(),
		)

	case spotifyMsg:
		if msg == nil || msg.Item == nil {
			m.IsPlaying = false
			m.TrackName = "Not Playing"
			m.Artist = ""
			return m, nil
		}

		m.IsPlaying = msg.Playing
		m.TrackName = msg.Item.Name
		m.Artist = ""
		m.Progress = int(msg.Progress)
		m.Duration = int(msg.Item.Duration)

		if len(msg.Item.Artists) > 0 {
			m.Artist = msg.Item.Artists[0].Name
		} else {
			m.Artist = "Unknown"
		}

		return m, nil
	}

	return m, nil
}


func (m Model) View() string {
	header := titleStyle.Render("  ♫  " + m.TrackName)
	artist := artistStyle.Render(m.Artist)

	width := 40
	percent := 0.0
	if m.Duration > 0 {
		percent = float64(m.Progress) / float64(m.Duration)
	}

	if percent > 1.0 {
		percent = 1.0
	}

	filled := int(percent * float64(width))

	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "━"
		} else if i == filled {
			bar += "⬤"
		} else {
			bar += "─"
		}
	}

	currentMin := m.Progress / 1000 / 60
	currentSec := (m.Progress / 1000) % 60
	status := statusStyle.Render(fmt.Sprintf("%d:%02d", currentMin, currentSec))

	content := lipgloss.JoinVertical(lipgloss.Left,
		header,
		artist,
		"\n",
		bar+" "+status,
		"\n",
		artistStyle.Render("[q] to quit, [space] to play/pause"),
		artistStyle.Render("[<] prev, [>] next, [left/right] seek"),
	)
	
	if m.Error != "" {
		content = lipgloss.JoinVertical(lipgloss.Left, content, "\n", lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render("Error: "+m.Error))
	}

	return windowStyle.Render(content)
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func checkSpotify(client *spotify.Client) tea.Cmd {
	return func() tea.Msg {
		if client == nil {
			return nil
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		playing, err := client.PlayerCurrentlyPlaying(ctx)
		if err != nil {
			return nil
		}

		return spotifyMsg(playing)
	}
}

func play(client *spotify.Client) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := client.Play(ctx)
		if err != nil {
			return errMsg(err)
		}
		return nil
	}
}

func pause(client *spotify.Client) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err := client.Pause(ctx)
		if err != nil {
			return errMsg(err)
		}
		return nil
	}
}

func seek(client *spotify.Client, positionMs int) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := client.Seek(ctx, positionMs)
		if err != nil {
			return errMsg(err)
		}
		return nil
	}
}

func next(client *spotify.Client) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := client.Next(ctx)
		if err != nil {
			return errMsg(err)
		}
		return nil
	}
}

func previous(client *spotify.Client) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := client.Previous(ctx)
		if err != nil {
			return errMsg(err)
		}
		return nil
	}
}