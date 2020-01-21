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

func toPNGPath(originalFilename, chartName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting user home directory: %w", err)
	}
	baseChartPath := path.Join(homeDir, ".flesch-index-data")
	_, filename := filepath.Split(originalFilename)
	noPath := strings.TrimSuffix(filename, path.Ext(filename))
	newFilename := fmt.Sprintf("%s.%s.png", noPath, chartName)

	return path.Join(baseChartPath, newFilename), nil
}

func ensureOutputDirectory() error {
	exampleFilepath, err := toPNGPath("fake.ext", "none")
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

type SyllableDistributionAnalysis struct {
	// Number of words that have a number of syllables
	// e.g. [4]6 would mean there are 6 words with 4 syllables
	SyllableDistribution map[int]int
	KeyOrder             []int
	ChartPath            string
}

func BuildSyllableAnalysis(document flesch.Document) (SyllableDistributionAnalysis, error) {
	pngPath, err := toPNGPath(document.Name(), "SyllableDistribution")
	if err != nil {
		return SyllableDistributionAnalysis{}, fmt.Errorf("generating PNG path for chart: %w", err)
	}
	analysis := SyllableDistributionAnalysis{
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
		return SyllableDistributionAnalysis{}, fmt.Errorf("creating a new plot: %w", err)
	}
	p.Title.Text = "Syllable Distribution"
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
	analysis.KeyOrder = syllableCounts
	for _, syllableCount := range syllableCounts {
		wordCounts = append(wordCounts, float64(analysis.SyllableDistribution[syllableCount]))
		wordCountLabels = append(wordCountLabels, strconv.Itoa(syllableCount))
	}

	barValue := plotter.Values(wordCounts)
	bar, err := plotter.NewBarChart(barValue, w)
	if err != nil {
		return SyllableDistributionAnalysis{}, fmt.Errorf("creating bar with values %v: %w", []float64(barValue), err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(0)

	p.Add(bar)
	p.Legend.Top = true
	p.NominalX(wordCountLabels...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, pngPath); err != nil {
		return SyllableDistributionAnalysis{}, fmt.Errorf("saving chart png: %w", err)
	}

	return analysis, nil
}

type Report struct {
	SyllableAnalysis      SyllableDistributionAnalysis
	SyllableRatioAnalysis SyllableRatioAnalysis
}

func Build(document flesch.Document) (Report, error) {
	// ensure output directory exists
	if err := ensureOutputDirectory(); err != nil {
		return Report{}, err
	}

	syllableAnalysis, err := BuildSyllableAnalysis(document)
	if err != nil {
		return Report{}, fmt.Errorf("building syllable distribution analysis: %w", err)
	}

	syllableRatioAnalysis, err := BuildSyllableRatioAnalysis(document)
	if err != nil {
		return Report{}, fmt.Errorf("building syllable ratio analysis: %w", err)
	}
	return Report{
		SyllableAnalysis:      syllableAnalysis,
		SyllableRatioAnalysis: syllableRatioAnalysis,
	}, nil
}

type charactersToSyllables struct {
	word       string
	characters int
	syllables  int
}

func (ratio charactersToSyllables) Ratio() float64 {
	return float64(ratio.characters) / float64(ratio.syllables)
}

// SyllableRatioAnalysis shows the top words with the highest ratio of characters to syllables
type SyllableRatioAnalysis struct {
	SyllableRatio map[string]float64
	KeyOrder      []string
	ChartPath     string
}

func BuildSyllableRatioAnalysis(document flesch.Document) (SyllableRatioAnalysis, error) {
	const numberCharted = 8
	pngPath, err := toPNGPath(document.Name(), "SyllableRatio")
	if err != nil {
		return SyllableRatioAnalysis{}, fmt.Errorf("generating PNG path for chart: %w", err)
	}
	analysis := SyllableRatioAnalysis{
		SyllableRatio: make(map[string]float64),
		ChartPath:     pngPath,
	}
	words := document.UniqueWords()
	var keyOrder []charactersToSyllables
	for _, word := range words {
		syllables := word.Syllables()
		ratio := charactersToSyllables{word.String(), len(word.Runes()), syllables}
		keyOrder = append(keyOrder, ratio)
		analysis.SyllableRatio[word.String()] = ratio.Ratio()
	}
	sort.SliceStable(keyOrder, func(i, j int) bool {
		return keyOrder[i].Ratio() > keyOrder[j].Ratio()
	})
	for i, ratio := range keyOrder {
		if i >= numberCharted {
			break
		}
		analysis.KeyOrder = append(analysis.KeyOrder, ratio.word)
		// limit displayed words to top several
	}
	var ratios []float64
	for _, word := range analysis.KeyOrder {
		if len(ratios) >= numberCharted {
			break
		}
		ratios = append(ratios, analysis.SyllableRatio[word])
	}

	// build chart
	p, err := plot.New()
	if err != nil {
		return SyllableRatioAnalysis{}, fmt.Errorf("creating a new plot: %w", err)
	}
	p.Title.Text = fmt.Sprintf("Top %d Syllable Ratio (characters to syllables)", numberCharted)
	p.Y.Label.Text = "Ratio"
	p.X.Label.Text = "Words"

	w := vg.Points(20)

	barValue := plotter.Values(ratios)
	bar, err := plotter.NewBarChart(barValue, w)
	if err != nil {
		return SyllableRatioAnalysis{}, fmt.Errorf("creating bar with values %v: %w", []float64(barValue), err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(1)

	p.Add(bar)
	p.Legend.Top = true
	p.NominalX(analysis.KeyOrder[:numberCharted]...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, pngPath); err != nil {
		return SyllableRatioAnalysis{}, fmt.Errorf("saving chart png: %w", err)
	}

	return analysis, nil
}
