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
| Creative Prompts | 25 | CC BY 4.0 |
| Alan's Koans | 27 | CC BY 4.0 |
| Magic 8 Ball | 20 | Public Domain |

## Adding a Deck

Create a new directory under `decks/` with:

1. **`_deck.yaml`** — deck metadata:
   ```yaml
   name: "My Deck"
   description: "A description of the deck"
   author: "Your Name"
   source: "https://example.com"
   license_type: "CC BY 4.0"
   card_count: 10
   tags: ["custom"]
   ```

2. **`_LICENSE`** — full license text for the card content.

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

This project is licensed under **Creative Commons Attribution 4.0 International (CC BY 4.0)**.

### Deck Licensing

The repository as a whole is CC BY 4.0, but individual decks may carry their own licenses
appropriate to their content. Each deck has a `_LICENSE` file in its directory and a
`license_type` field in its `_deck.yaml`. Specifically:

- **Original content** (creative_prompts, alan_koans): CC BY 4.0
- **Public domain works** (rider_waite_tarot, i_ching): Public Domain, with attribution to original authors
- **Uncopyrightable facts** (magic_8_ball): Public Domain

See each deck's `_LICENSE` file for full details and attribution.
