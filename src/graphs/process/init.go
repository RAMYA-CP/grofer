package process

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// PerProcPage holds the ui elements rendered by the command grofer proc -p PID
type PerProcPage struct {
	Grid             *ui.Grid
	CPUChart         *widgets.Gauge
	MemChart         *widgets.Gauge
	PIDTable         *widgets.Table
	ChildProcsList   *widgets.List
	CTXSwitchesChart *widgets.BarChart
	PageFaultsChart  *widgets.BarChart
	MemStatsChart    *widgets.BarChart
}

// NewProcPage initializes a new page from the PerProcPage struct and returns it
func NewPerProcPage() *PerProcPage {
	page := &PerProcPage{
		Grid:             ui.NewGrid(),
		CPUChart:         widgets.NewGauge(),
		MemChart:         widgets.NewGauge(),
		PIDTable:         widgets.NewTable(),
		ChildProcsList:   widgets.NewList(),
		CTXSwitchesChart: widgets.NewBarChart(),
		PageFaultsChart:  widgets.NewBarChart(),
		MemStatsChart:    widgets.NewBarChart(),
	}
	page.InitPerProc()
	return page
}

// InitPerProc initializes and sets the ui and grid for grofer proc -p PID
func (page *PerProcPage) InitPerProc() {
	// Initialize Gauge for CPU Chart
	page.CPUChart.Title = " CPU % "
	page.CPUChart.BarColor = ui.ColorGreen
	page.CPUChart.BorderStyle.Fg = ui.ColorCyan
	page.CPUChart.TitleStyle.Fg = ui.ColorWhite

	// Initialize Gauge for Memory Chart
	page.MemChart.Title = " Mem % "
	page.MemChart.BarColor = ui.ColorGreen
	page.MemChart.BorderStyle.Fg = ui.ColorCyan
	page.MemChart.TitleStyle.Fg = ui.ColorWhite

	// Initialize Table for PID Details Table
	page.PIDTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	page.PIDTable.TextAlignment = ui.AlignCenter
	page.PIDTable.RowSeparator = false
	page.PIDTable.Title = " PID "
	page.PIDTable.BorderStyle.Fg = ui.ColorCyan
	page.PIDTable.TitleStyle.Fg = ui.ColorWhite

	// Initialize List for Child Processes list
	page.ChildProcsList.Title = " Child Processes "
	page.ChildProcsList.BorderStyle.Fg = ui.ColorCyan
	page.ChildProcsList.TitleStyle.Fg = ui.ColorWhite

	// Initialize Bar Chart for CTX Swicthes Chart
	page.CTXSwitchesChart.Data = []float64{0, 0}
	page.CTXSwitchesChart.Labels = []string{"Volun", "Involun"}
	page.CTXSwitchesChart.Title = " Ctx switches "
	page.CTXSwitchesChart.BorderStyle.Fg = ui.ColorCyan
	page.CTXSwitchesChart.TitleStyle.Fg = ui.ColorWhite
	page.CTXSwitchesChart.BarWidth = 10
	page.CTXSwitchesChart.BarColors = []ui.Color{ui.ColorGreen, ui.ColorCyan}
	page.CTXSwitchesChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	page.CTXSwitchesChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	// Initialize Bar Chart for Page Faults Chart
	page.PageFaultsChart.Data = []float64{0, 0}
	page.PageFaultsChart.Labels = []string{"minr", "mjr"}
	page.PageFaultsChart.Title = " Page Faults "
	page.PageFaultsChart.BorderStyle.Fg = ui.ColorCyan
	page.PageFaultsChart.TitleStyle.Fg = ui.ColorWhite
	page.PageFaultsChart.BarWidth = 10
	page.PageFaultsChart.BarColors = []ui.Color{ui.ColorGreen, ui.ColorCyan}
	page.PageFaultsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	page.PageFaultsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	// Initialize Bar Chart for Memory Stats Chart
	page.MemStatsChart.Data = []float64{0, 0, 0, 0}
	page.MemStatsChart.Labels = []string{"RSS", "Data", "Stack", "Swap"}
	page.MemStatsChart.Title = " Mem Stats (mb) "
	page.MemStatsChart.BorderStyle.Fg = ui.ColorCyan
	page.MemStatsChart.TitleStyle.Fg = ui.ColorWhite
	page.MemStatsChart.BarWidth = 10
	page.MemStatsChart.BarColors = []ui.Color{ui.ColorGreen, ui.ColorMagenta, ui.ColorYellow, ui.ColorCyan}
	page.MemStatsChart.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorWhite)}
	page.MemStatsChart.NumStyles = []ui.Style{ui.NewStyle(ui.ColorBlack)}

	// Initialize Grid layout
	page.Grid.Set(
		ui.NewCol(0.5,
			ui.NewRow(0.125, page.CPUChart),
			ui.NewRow(0.125, page.MemChart),
			ui.NewRow(0.35, page.PIDTable),
			ui.NewRow(0.4, page.ChildProcsList),
		),
		ui.NewCol(0.5,
			ui.NewRow(0.6,
				ui.NewCol(0.5, page.CTXSwitchesChart),
				ui.NewCol(0.5, page.PageFaultsChart),
			),
			ui.NewRow(0.4, page.MemStatsChart),
		),
	)

	w, h := ui.TerminalDimensions()
	page.Grid.SetRect(0, 0, w, h)
}

// AllProcPage struct holds the ui elements rendered by the grofer proc command
type AllProcPage struct {
	Grid         *ui.Grid
	HeadingTable *widgets.Table
	BodyList     *widgets.List
}

// NewAllProcsPage initializes a new page from the AllProcPage struct and returns it
func NewAllProcsPage() *AllProcPage {
	page := &AllProcPage{
		Grid:         ui.NewGrid(),
		HeadingTable: widgets.NewTable(),
		BodyList:     widgets.NewList(),
	}
	page.InitAllProc()
	return page
}

// InitAllProc initializes and sets the ui and grid for grofer proc
func (page *AllProcPage) InitAllProc() {
	page.HeadingTable.TextStyle = ui.NewStyle(ui.ColorWhite)
	page.HeadingTable.Rows = [][]string{[]string{" PID",
		" Command",
		" CPU",
		" Memory",
		" Status",
		" Foreground",
		" Creation Time",
		" Thread Count",
	}}
	page.HeadingTable.ColumnWidths = []int{10, 40, 10, 10, 8, 12, 23, 15}
	page.HeadingTable.TextAlignment = ui.AlignLeft
	page.HeadingTable.RowSeparator = false

	page.BodyList.TextStyle = ui.NewStyle(ui.ColorWhite)
	page.BodyList.TitleStyle.Fg = ui.ColorCyan

	page.Grid.Set(
		ui.NewRow(0.12, page.HeadingTable),
		ui.NewRow(0.88, page.BodyList),
	)

	w, h := ui.TerminalDimensions()
	page.Grid.SetRect(0, 0, w, h)
}
