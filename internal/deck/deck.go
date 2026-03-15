package deck

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// IndexMeta represents a deck's index.yaml metadata.
type IndexMeta struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Author      string   `yaml:"author"`
	Source      string   `yaml:"source"`
	LicenseType string   `yaml:"license_type"`
	CardCount   int      `yaml:"card_count"`
	Tags        []string `yaml:"tags"`
}

// Card represents a single card parsed from a markdown file.
type Card struct {
	Title    string   `yaml:"title"`
	Number   int      `yaml:"number"`
	Keywords []string `yaml:"keywords"`
	AssetURL string   `yaml:"asset_url"`
	Body     string   `yaml:"-"`
	Filename string   `yaml:"-"`
}

// Deck is a loaded deck with metadata and cards.
type Deck struct {
	DirName string
	Meta    IndexMeta
	Cards   []Card
	Path    string
}

// LoadAll discovers and loads all decks from the given cards directory.
func LoadAll(cardsDir string) ([]Deck, error) {
	entries, err := os.ReadDir(cardsDir)
	if err != nil {
		return nil, fmt.Errorf("reading cards directory: %w", err)
	}

	var decks []Deck
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		d, err := LoadDeck(filepath.Join(cardsDir, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("loading deck %s: %w", entry.Name(), err)
		}
		decks = append(decks, d)
	}
	return decks, nil
}

// LoadDeck loads a single deck from the given directory.
func LoadDeck(dir string) (Deck, error) {
	d := Deck{
		DirName: filepath.Base(dir),
		Path:    dir,
	}

	// Parse index.yaml
	indexPath := filepath.Join(dir, "index.yaml")
	indexData, err := os.ReadFile(indexPath)
	if err != nil {
		return d, fmt.Errorf("reading index.yaml: %w", err)
	}
	if err := yaml.Unmarshal(indexData, &d.Meta); err != nil {
		return d, fmt.Errorf("parsing index.yaml: %w", err)
	}

	// Find and parse card files
	entries, err := os.ReadDir(dir)
	if err != nil {
		return d, fmt.Errorf("reading deck directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		card, err := ParseCard(filepath.Join(dir, entry.Name()))
		if err != nil {
			return d, fmt.Errorf("parsing card %s: %w", entry.Name(), err)
		}
		d.Cards = append(d.Cards, card)
	}

	return d, nil
}

// ParseCard parses a card markdown file with YAML frontmatter.
func ParseCard(path string) (Card, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Card{}, fmt.Errorf("reading file: %w", err)
	}

	var card Card
	card.Filename = filepath.Base(path)

	// Split frontmatter from body
	content := string(data)
	if !strings.HasPrefix(content, "---") {
		return Card{}, fmt.Errorf("missing YAML frontmatter delimiter")
	}

	// Find end of frontmatter
	rest := content[3:]
	idx := strings.Index(rest, "\n---")
	if idx < 0 {
		return Card{}, fmt.Errorf("missing closing YAML frontmatter delimiter")
	}

	frontmatter := rest[:idx]
	body := strings.TrimSpace(rest[idx+4:])

	if err := yaml.Unmarshal([]byte(frontmatter), &card); err != nil {
		return Card{}, fmt.Errorf("parsing frontmatter: %w", err)
	}
	card.Body = body

	return card, nil
}

// Validate checks a deck for common issues and returns a list of errors.
func (d *Deck) Validate() []string {
	var errs []string

	if d.Meta.Name == "" {
		errs = append(errs, "index.yaml: missing name")
	}
	if d.Meta.CardCount > 0 && len(d.Cards) != d.Meta.CardCount {
		errs = append(errs, fmt.Sprintf("card_count mismatch: index.yaml says %d, found %d card files", d.Meta.CardCount, len(d.Cards)))
	}

	// Check LICENSE exists
	if _, err := os.Stat(filepath.Join(d.Path, "LICENSE")); os.IsNotExist(err) {
		errs = append(errs, "missing LICENSE file")
	}

	for _, c := range d.Cards {
		if c.Title == "" {
			errs = append(errs, fmt.Sprintf("%s: missing title", c.Filename))
		}
	}

	return errs
}

// DrawRandom selects n random cards from the deck without replacement.
func (d *Deck) DrawRandom(n int) []Card {
	if n >= len(d.Cards) {
		shuffled := make([]Card, len(d.Cards))
		copy(shuffled, d.Cards)
		rand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})
		return shuffled
	}
	indices := rand.Perm(len(d.Cards))[:n]
	result := make([]Card, n)
	for i, idx := range indices {
		result[i] = d.Cards[idx]
	}
	return result
}

// FindCardsDir locates the cards/ directory relative to the executable or CWD.
func FindCardsDir() (string, error) {
	// Try relative to executable
	exe, err := os.Executable()
	if err == nil {
		dir := filepath.Join(filepath.Dir(exe), "cards")
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir, nil
		}
	}

	// Try CWD
	cwd, err := os.Getwd()
	if err == nil {
		dir := filepath.Join(cwd, "cards")
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir, nil
		}
	}

	return "", fmt.Errorf("could not find cards/ directory")
}

// RenderCard formats a card for terminal display.
func RenderCard(card Card, deckName string) string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("  %s\n", card.Title))
	buf.WriteString(fmt.Sprintf("  from %s\n", deckName))

	if len(card.Keywords) > 0 {
		buf.WriteString(fmt.Sprintf("  %s\n", strings.Join(card.Keywords, " | ")))
	}

	buf.WriteString("\n")
	// Indent body lines
	for _, line := range strings.Split(card.Body, "\n") {
		buf.WriteString(fmt.Sprintf("  %s\n", line))
	}

	return buf.String()
}
