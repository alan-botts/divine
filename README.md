# divine

A general-purpose divination CLI. Draw cards from tarot, I Ching, creative prompts, and more.

## Installation

```bash
go install github.com/alan-botts/divine@latest
```

Or clone and build:

```bash
git clone https://github.com/alan-botts/divine.git
cd divine
go build -o divine .
```

## Usage

### Draw a card

```bash
divine draw              # draw one card from all decks
divine draw -n 3         # draw three cards
divine draw --deck rider_waite_tarot   # draw from a specific deck
divine draw --deck i_ching -n 1        # one I Ching hexagram
```

### List available decks

```bash
divine decks
```

### Validate deck data

```bash
divine tools validate-decks
```

## Included Decks

| Deck | Cards | License |
|------|-------|---------|
| Rider-Waite-Smith Tarot (Major Arcana) | 22 | Public Domain |
| I Ching (Book of Changes) | 64 | Public Domain |
| Creative Prompts | 25 | CC0 |
| Magic 8 Ball | 20 | Public Domain |

## Adding a Deck

Create a new directory under `cards/` with:

1. **`index.yaml`** — deck metadata:
   ```yaml
   name: "My Deck"
   description: "A description of the deck"
   author: "Your Name"
   source: "https://example.com"
   license_type: "CC0"
   card_count: 10
   tags: ["custom"]
   ```

2. **`LICENSE`** — full license text for the card content.

3. **Card files** (`*.md`) — markdown files with YAML frontmatter:
   ```markdown
   ---
   title: "Card Name"
   number: 1
   keywords: ["keyword1", "keyword2"]
   asset_url: ""
   ---

   Card description and interpretation text here.
   ```

Run `divine tools validate-decks` to check your deck.

## License

MIT (CLI code). Individual deck content is licensed as noted in each deck's `LICENSE` file.
