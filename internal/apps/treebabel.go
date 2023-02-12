package apps

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/edualb-challenge/treebabel/internal/iofiles"
	"github.com/edualb-challenge/treebabel/internal/models"
	"github.com/edualb-challenge/treebabel/internal/tree"
)

// TreeBabel is an application responsible to process the average delivery time
type TreeBabel struct {
	EventsFile string
	WindowSize uint64

	firstTreeTime time.Time
	lastTreeTime  time.Time
	leaves        []float64
	timeLeafMap   map[time.Time]int64
}

// Run executes the treebabel application returning the average delivery time segment
func (app TreeBabel) Run() ([]models.AverageDeliveryTime, error) {
	err := app.setRangeOfTimes()
	if err != nil {
		return nil, err
	}

	err = app.setLeaves()
	if err != nil {
		return nil, err
	}

	segTree := tree.NewSegment(app.leaves)
	err = app.buildTree(segTree)
	if err != nil {
		return nil, err
	}

	avarageDeliveryTimeSegment, err := app.getAverageDeliveryTime(segTree)
	if err != nil {
		return nil, err
	}

	return avarageDeliveryTimeSegment, nil
}

func (app *TreeBabel) setRangeOfTimes() error {
	firstTreeTime, err := app.getFirstTreeTime()
	if err != nil {
		return err
	}

	lastTreeTime, err := app.getLastTreeTime()
	if err != nil {
		return err
	}

	app.firstTreeTime = firstTreeTime
	app.lastTreeTime = lastTreeTime
	return nil
}

func (app TreeBabel) getFirstTreeTime() (time.Time, error) {
	var firstTime time.Time
	firstLine, err := iofiles.GetFirstLine(app.EventsFile)
	if err != nil {
		return firstTime, err
	}

	firstTD, err := models.GetEventFromBytes(firstLine)
	if err != nil {
		return firstTime, err
	}
	firstTimestamp := firstTD.Timestamp
	firstTime = time.Date(firstTimestamp.Year(), firstTimestamp.Month(), firstTimestamp.Day(), firstTimestamp.Hour(), firstTimestamp.Minute()-int(app.WindowSize), 0, 0, firstTimestamp.Location())

	return firstTime, nil
}

func (app TreeBabel) getLastTreeTime() (time.Time, error) {
	var lastTime time.Time
	lastLine, err := iofiles.GetLastLine(app.EventsFile)
	if err != nil {
		return lastTime, err
	}

	lastTD, err := models.GetEventFromBytes(lastLine)
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

func (app *TreeBabel) setLeaves() error {
	leavesQty := int(app.lastTreeTime.Sub(app.firstTreeTime).Minutes()) + 1
	leaves := make([]float64, leavesQty)
	timeLeafMap := make(map[time.Time]int64, leavesQty)

	for i := range leaves {
		timeLeafMap[app.firstTreeTime.Add(time.Duration(i)*time.Minute)] = int64(i)
	}

	app.leaves = leaves
	app.timeLeafMap = timeLeafMap

	return nil
}

func (app TreeBabel) buildTree(segTree *tree.Segment) error {
	file, err := os.Open(app.EventsFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		event, err := models.GetEventFromBytes(scanner.Bytes())
		if err != nil {
			return err
		}

		timestamp := event.Timestamp

		minute := timestamp.Minute()
		if timestamp.Second() > 0 {
			minute += 1
		}

		eventTime := time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), minute, 0, 0, timestamp.Location())

		index, ok := app.timeLeafMap[eventTime]
		if !ok {
			return fmt.Errorf("time '%v' is not valid", eventTime.Format("2006-12-26 18:11:08"))
		}
		segTree.Set(index, event.Duration)
	}
	return nil
}

func (app TreeBabel) getAverageDeliveryTime(segTree *tree.Segment) ([]models.AverageDeliveryTime, error) {
	averageDeliveryTime := []models.AverageDeliveryTime{}

	lastIndex, ok := app.timeLeafMap[app.lastTreeTime]
	if !ok {
		return nil, fmt.Errorf("failing getting the last time index for '%s", app.lastTreeTime.Format("2006-12-26 18:11:08"))
	}

	beginsTime := app.firstTreeTime.Add(time.Duration(app.WindowSize) * time.Minute)
	beginsAt, ok := app.timeLeafMap[app.firstTreeTime.Add(time.Duration(app.WindowSize)*time.Minute)]
	if !ok {
		return nil, fmt.Errorf("failing getting the last time index for '%s", app.firstTreeTime)
	}

	for i := 0; int(i) < len(app.leaves); i++ {
		index := int64(i)
		nextIndex := beginsAt + int64(i)

		if nextIndex-1 >= lastIndex {
			break
		}

		avg := segTree.Query(index, nextIndex)
		indexTime := beginsTime.Add(time.Duration(index) * time.Minute)

		var adt models.AverageDeliveryTime
		adt.Average = avg
		adt.Date = models.Timestamp{
			Time: indexTime,
		}
		averageDeliveryTime = append(averageDeliveryTime, adt)
	}
	return averageDeliveryTime, nil
}
