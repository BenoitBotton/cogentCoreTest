package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func main() {
	b := core.NewBody()
	grid := core.NewFrame(b)
	grid.Styler(func(s *styles.Style) {
		s.Display = styles.Grid
		s.Columns = 3
	})
	for range 9 {
		bt := core.NewButton(grid).SetType(core.ButtonAction)
		bt.Styler(func(s *styles.Style) {
			s.Border.Width.Set(units.Dp(1))
			s.Border.Radius.Zero()
			s.Min.Set(units.Em(2))
		})
	}
	b.RunMainWindow()
}
