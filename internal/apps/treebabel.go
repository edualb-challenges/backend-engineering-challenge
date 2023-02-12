package apps

import (
	"fmt"
	"time"

	"github.com/edualb-challenge/treebabel/internal/iofiles"
	"github.com/edualb-challenge/treebabel/internal/models"
)

// TreeBabel is an application responsible to process the average delivery time
type TreeBabel struct {
	TranslationDeliveredFile string
	WindowSize               uint64
}

// Run executes the treebabel application
func (app TreeBabel) Run() error {
	firstTreeTime, err := app.getFirstTreeTime()
	if err != nil {
		return err
	}

	lastTreeTime, err := app.getLastTreeTime()
	if err != nil {
		return err
	}

	leavesQty := int(lastTreeTime.Sub(firstTreeTime).Minutes())

	leaves := make([]float64, leavesQty)
	leafTimeMap := make(map[int]time.Time, leavesQty)

	for i := range leaves {
		leafTimeMap[i] = firstTreeTime.Add(time.Duration(i) * time.Minute)
	}

	for i := range leaves {
		t, ok := leafTimeMap[i]
		if !ok {
			return fmt.Errorf("invalid index of %d", i)
		}
		fmt.Printf("time in %d size: %v\n", i, t)
	}

	fmt.Printf("leaves size: %d\n\n", len(leaves))

	return nil
}

func (app TreeBabel) getLastTreeTime() (time.Time, error) {
	var lastTime time.Time
	lastLine, err := iofiles.GetLastLine(app.TranslationDeliveredFile)
	if err != nil {
		return lastTime, err
	}

	lastTD, err := models.GetTranslationDeliveredFromBytes(lastLine)
	if err != nil {
		return lastTime, nil
	}
	lastTimestamp := lastTD.Timestamp

	minute := lastTimestamp.Minute()
	if lastTimestamp.Second() > 0 {
		minute = lastTimestamp.Minute() + 1
	}

	lastTime = time.Date(lastTimestamp.Year(), lastTimestamp.Month(), lastTimestamp.Day(), lastTimestamp.Hour(), minute, 0, 0, lastTimestamp.Location())

	return lastTime, nil
}

func (app TreeBabel) getFirstTreeTime() (time.Time, error) {
	var firstTime time.Time
	firstLine, err := iofiles.GetFirstLine(app.TranslationDeliveredFile)
	if err != nil {
		return firstTime, err
	}

	firstTD, err := models.GetTranslationDeliveredFromBytes(firstLine)
	if err != nil {
		return firstTime, err
	}
	firstTimestamp := firstTD.Timestamp
	firstTime = time.Date(firstTimestamp.Year(), firstTimestamp.Month(), firstTimestamp.Day(), firstTimestamp.Hour(), firstTimestamp.Minute()-int(app.WindowSize), 0, 0, firstTimestamp.Location())

	return firstTime, nil
}
