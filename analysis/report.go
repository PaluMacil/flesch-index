package analysis

import (
	"fmt"
	"github.com/PaluMacil/flesch-index/flesch"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func toPNGPath(originalFilename string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user home directory: %w", err)
	}
	baseChartPath := path.Join(homeDir, ".flesch-index-data")
	_, filename := filepath.Split(originalFilename)
	noPath := strings.TrimSuffix(filename, path.Ext(filename))
	newFilename := noPath + ".png"

	return path.Join(baseChartPath, newFilename), nil
}

func ensureOutputDirectory() error {
	exampleFilepath, err := toPNGPath("fake.ext")
	if err != nil {

		return fmt.Errorf("determining output directory: %w", err)
	}
	folder, _ := path.Split(exampleFilepath)
	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {

		return fmt.Errorf("ensuring creation of output directory: %w", err)
	}

	return nil
}

type SyllableAnalysis struct {
	// Number of words that have a number of syllables
	// e.g. [4]6 would mean there are 6 words with 4 syllables
	SyllableDistribution map[int]int
	ChartPath            string
}

func BuildSyllableAnalysis(document flesch.Document) (SyllableAnalysis, error) {
	pngPath, err := toPNGPath(document.Name())
	if err != nil {
		return SyllableAnalysis{}, fmt.Errorf("generating PNG path for chart: %w", err)
	}
	analysis := SyllableAnalysis{
		SyllableDistribution: make(map[int]int),
		ChartPath:            pngPath,
	}
	for _, word := range document.Words() {
		syllables := word.Syllables()
		value, exists := analysis.SyllableDistribution[syllables]
		if exists {
			analysis.SyllableDistribution[syllables] = value + 1
		} else {
			analysis.SyllableDistribution[syllables] = 1
		}
	}

	p, err := plot.New()
	if err != nil {
		return SyllableAnalysis{}, fmt.Errorf("creating a new plot: %w", err)
	}
	p.Title.Text = "Words per Syllable Count"
	p.Y.Label.Text = "Word Count"
	p.X.Label.Text = "Syllables"

	w := vg.Points(30)

	// Store the syllable count keys and word count values in slices in sorted order
	var syllableCounts []int
	var wordCountLabels []string
	var wordCounts []float64
	for syllableCount := range analysis.SyllableDistribution {
		syllableCounts = append(syllableCounts, syllableCount)
	}
	sort.Ints(syllableCounts)
	for _, syllableCount := range syllableCounts {
		wordCounts = append(wordCounts, float64(analysis.SyllableDistribution[syllableCount]))
		wordCountLabels = append(wordCountLabels, strconv.Itoa(syllableCount))
	}

	barValue := plotter.Values(wordCounts)
	bar, err := plotter.NewBarChart(barValue, w)
	if err != nil {
		return SyllableAnalysis{}, fmt.Errorf("creating bar with values %v: %w", []float64(barValue), err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(0)

	p.Add(bar)
	p.Legend.Top = true
	p.NominalX(wordCountLabels...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, pngPath); err != nil {
		return SyllableAnalysis{}, fmt.Errorf("saving chart png: %w", err)
	}

	return analysis, nil
}

type Report struct {
	SyllableAnalysis SyllableAnalysis
}

func Build(document flesch.Document) (Report, error) {
	// ensure output directory exists
	if err := ensureOutputDirectory(); err != nil {
		return Report{}, err
	}

	syllableAnalysis, err := BuildSyllableAnalysis(document)
	if err != nil {
		return Report{}, fmt.Errorf("building syllable analysis: %w", err)
	}
	return Report{
		SyllableAnalysis: syllableAnalysis,
	}, nil
}
