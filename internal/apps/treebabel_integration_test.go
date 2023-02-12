package apps_test

import (
	"testing"
	"time"

	"github.com/edualb-challenge/treebabel/internal/apps"
	"github.com/edualb-challenge/treebabel/internal/models"
)

func TestTreeBabelApp(t *testing.T) {
	type input struct {
		inputFile  string
		windowSize uint64
	}
	type output struct {
		avg []models.AverageDeliveryTime
	}
	type testCase struct {
		name string
		in   input
		out  output
	}

	defaultTime, err := time.Parse("2006-01-02 15:04:05.000000", "2018-12-26 18:11:00.000000")
	if err != nil {
		t.Errorf("unexpected error when parsing default date, got: %v", err)
		t.FailNow()
	}

	// In order to facilitates the manutenability, we include the 'tree' comment
	// that represents the current tree the test is related with
	tests := []testCase{
		{
			name: "should succeed when use the challenge input",
			in: input{
				inputFile:  "../../testdata/treebabel/challenge-input.json",
				windowSize: 10,
			},
			out: output{
				avg: []models.AverageDeliveryTime{
					{
						Date: models.Timestamp{
							Time: defaultTime,
						},
						Average: 0,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(1 * time.Minute),
						},
						Average: 20,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(2 * time.Minute),
						},
						Average: 20,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(3 * time.Minute),
						},
						Average: 20,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(4 * time.Minute),
						},
						Average: 20,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(5 * time.Minute),
						},
						Average: 25.5,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(6 * time.Minute),
						},
						Average: 25.5,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(7 * time.Minute),
						},
						Average: 25.5,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(8 * time.Minute),
						},
						Average: 25.5,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(9 * time.Minute),
						},
						Average: 25.5,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(10 * time.Minute),
						},
						Average: 25.5,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(11 * time.Minute),
						},
						Average: 31,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(12 * time.Minute),
						},
						Average: 31,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(13 * time.Minute),
						},
						Average: 42.5,
					},
				},
			},
		},
		{
			name: "should succeed when use the complex input",
			in: input{
				inputFile:  "../../testdata/treebabel/complex.json",
				windowSize: 3,
			},
			out: output{
				avg: []models.AverageDeliveryTime{
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(-1 * time.Minute),
						},
						Average: 0,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime,
						},
						Average: 10,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(1 * time.Minute),
						},
						Average: 15,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(2 * time.Minute),
						},
						Average: 15,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(3 * time.Minute),
						},
						Average: 30,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(4 * time.Minute),
						},
						Average: 40,
					},
					{
						Date: models.Timestamp{
							Time: defaultTime.Add(5 * time.Minute),
						},
						Average: 35.5,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			app, err := apps.NewTreeBabel(tc.in.inputFile, tc.in.windowSize)
			if err != nil {
				t.Errorf("unexpected error when getting app, got: %v", err)
				t.FailNow()
			}

			got, err := app.Run()
			if err != nil {
				t.Errorf("unexpected error when running app, got: %v", err)
				t.FailNow()
			}

			if len(got) != len(tc.out.avg) {
				t.Errorf("unexpected average delivery time segment length, got %d, wants %d", len(got), len(tc.out.avg))
				t.FailNow()
			}

			for i, avg := range got {
				if avg.Average != tc.out.avg[i].Average {
					t.Errorf("unexpected average [index %d], got %f, wants %f", i, avg.Average, tc.out.avg[i].Average)
					t.FailNow()
				}
				if avg.Date != tc.out.avg[i].Date {
					t.Errorf("unexpected date [index %d], got %v, wants %v", i, avg.Date, tc.out.avg[i].Date)
					t.FailNow()
				}
			}
		})
	}
}
