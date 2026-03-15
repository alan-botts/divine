#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DECK_DIR="$ROOT_DIR/decks/botts_tarot"
SOURCE_URL="https://raw.githubusercontent.com/LindseyB/tarot-api/main/tarot.json"
TMP_JSON="$(mktemp)"

cleanup() {
  rm -f "$TMP_JSON"
}
trap cleanup EXIT

curl -fsSL "$SOURCE_URL" -o "$TMP_JSON"

mkdir -p "$DECK_DIR"
find "$DECK_DIR" -maxdepth 1 -type f \( -name '*.md' -o -name '*.mdx' -o -name '_deck.yaml' -o -name '_LICENSE' \) -delete

cat > "$DECK_DIR/_deck.yaml" <<'YAML'
name: "Botts Tarot"
description: "A 78-card tarot deck generated from LindseyB/tarot-api"
author: "Alan Botts <alan.botts@strangerloops.com>"
source: "https://github.com/LindseyB/tarot-api/blob/main/tarot.json"
license_type: "MIT"
card_count: 78
tags: ["tarot", "divination", "botts_tarot"]
YAML

cat > "$DECK_DIR/_LICENSE" <<'LICENSE'
MIT License

Copyright (c) 2026 Alan Botts

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Source data attribution:
- Upstream JSON source: https://github.com/LindseyB/tarot-api/blob/main/tarot.json
- This deck includes generated formatting and derived card themes.
_LICENSE

jq -c '.cards[]' "$TMP_JSON" | nl -ba -w3 -nrz | while IFS=$'\t' read -r idx card; do
  idx_num="$((10#$idx))"
  name="$(jq -r '.name' <<< "$card")"
  slug="$(printf '%s' "$name" | tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/_/g; s/^_+|_+$//g')"
  file="$DECK_DIR/${idx}_${slug}.mdx"

  jq -r --arg idx "$idx_num" '
    def yq: gsub("\\\\";"\\\\\\\\") | gsub("\"";"\\\"");
    def rankstr: if (.rank|type) == "number" then (.rank|tostring) else .rank end;
    def arr($a): "[" + ($a | map("\"" + (. | yq) + "\"") | join(", ")) + "]";
    def themes: (.meanings | to_entries[0].value);
    [
      "---",
      "title: \"" + (.name|yq) + "\"",
      "number: " + $idx,
      "keywords: " + arr(([.suit, rankstr, .element] + .sign + themes) | map(select(. != null and . != "")) | unique),
      "asset_url: \"\"",
      "rank: \"" + (rankstr|yq) + "\"",
      "suit: \"" + (.suit|yq) + "\"",
      "planet: " + (if .planet == null then "null" else "\"" + (.planet|yq) + "\"" end),
      "element: " + (if .element == null then "null" else "\"" + (.element|yq) + "\"" end),
      "sign: " + arr(.sign),
      "---",
      "",
      (.name)
    ] | join("\n")
  ' <<< "$card" > "$file"
done

echo "Generated deck at $DECK_DIR"
