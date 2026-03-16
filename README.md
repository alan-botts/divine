# divine

A general-purpose divination CLI. Draw cards from tarot, I Ching, runes, koans, wisdom literature, and more. All deck data is embedded in the binary — no external files needed at runtime.

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

### Draw cards

```bash
divine draw                              # draw one card from the full pool
divine draw -n 3                         # draw three cards
divine draw --deck rider_waite_tarot     # draw from a specific deck
divine draw --deck tao_te_ching -n 3     # three chapters of the Tao Te Ching
divine draw --deck stoic_meditations     # a Stoic meditation
```

### List available decks

```bash
divine decks
```

### Validate deck data

```bash
divine tools validate-decks
```

### Version

```bash
divine version
```

## Included Decks

950 cards across 20 decks.

| Deck | Dir name | Cards | License |
|------|----------|-------|---------|
| Aesop's Fables | `aesops_fables` | 50 | Public Domain |
| Botts Koans | `botts_koans` | 87 | CC BY 4.0 |
| Botts Playing Cards (52 + 2 Jokers) | `botts_playing_cards` | 54 | CC BY 4.0 |
| Botts Tarot (full 78-card deck) | `botts_tarot` | 78 | MIT |
| Creative Prompts | `creative_prompts` | 25 | CC BY 4.0 |
| Decision Heuristics | `decision_heuristics` | 52 | CC BY 4.0 |
| Elder Futhark Runes | `elder_futhark` | 24 | Public Domain |
| The Gateless Gate (Mumonkan) | `gateless_gate` | 48 | Public Domain |
| Geomancy | `geomancy` | 16 | Public Domain |
| I Ching (Book of Changes) | `i_ching` | 64 | Public Domain |
| Magic 8 Ball | `magic_8_ball` | 20 | Public Domain |
| Mark Twain's Wit | `mark_twain` | 55 | Public Domain |
| Ogham (Celtic Tree Alphabet) | `ogham` | 20 | Public Domain |
| Oscar Wilde's Epigrams | `oscar_wilde` | 55 | Public Domain |
| Poor Richard's Almanac | `poor_richards` | 55 | Public Domain |
| Rider-Waite-Smith Tarot (Major Arcana) | `rider_waite_tarot` | 22 | Public Domain |
| Stoic Meditations | `stoic_meditations` | 52 | Public Domain |
| Tao Te Ching | `tao_te_ching` | 81 | Public Domain |
| The Prince (Machiavelli) | `the_prince` | 42 | Public Domain |
| World Proverbs | `world_proverbs` | 50 | Public Domain |

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

3. **Card files** (`.md` or `.mdx`) — markdown with YAML frontmatter:
   ```markdown
   ---
   title: "Card Name"
   number: 1
   keywords: ["keyword1", "keyword2"]
   ---

   Card description and interpretation text here.
   ```

Run `divine tools validate-decks` to check your deck. Decks are embedded at build time via `go:embed`, so rebuild after adding or modifying deck files.

## License

This project is licensed under **Creative Commons Attribution 4.0 International (CC BY 4.0)**.

Individual decks carry their own licenses — see `_deck.yaml` and `_LICENSE` in each deck directory. Public domain works include attribution to original authors. Original content (Botts Koans, Creative Prompts, Decision Heuristics) is CC BY 4.0. Botts Tarot is MIT.
