#!/usr/bin/env bash
#
# Runs the exercises in order and stops at the first one that is not solved yet.
#
#   exercises/run.sh              # up to the first unsolved exercise
#   exercises/run.sh --all        # every exercise, keep going
#   exercises/run.sh 02_strings   # one topic
#   exercises/run.sh --solutions  # check the reference solutions still pass
#
# An exercise is solved when what it prints matches exercises/expected/. The
# expectations are generated from the reference solutions by
# `go test ./exercises -update`, so an exercise cannot promise output that is
# impossible to produce.
set -uo pipefail

cd "$(dirname "$0")/.."

BIN="${ROCKET_LANG:-}"
if [[ -z "$BIN" ]]; then
  BIN="$(mktemp -d)/rocket-lang"
  go build -o "$BIN" . || { echo "could not build the interpreter"; exit 1; }
fi

ALL=0
FROM_SOLUTIONS=0
FILTER=""

for arg in "$@"; do
  case "$arg" in
    --all)       ALL=1 ;;
    --solutions) FROM_SOLUTIONS=1; ALL=1 ;;
    -h|--help)   sed -n '2,14p' "$0" | sed 's|^# \{0,1\}||'; exit 0 ;;
    *)           FILTER="$arg" ;;
  esac
done

# The interpreter prints the program's own final value, which is a nil as soon
# as a program ends in puts(). The exercise did not ask for it, so one trailing
# nil is dropped -- the same rule generate.go applies to the expectations.
strip_final_value() {
  awk 'BEGIN { held = 0 }
       held { print line }
       { line = $0; held = 1 }
       END { if (held && line != "nil") print line }'
}

passed=0
failed=0
total=0

for expected in $(find exercises/expected -name '*.txt' | sort); do
  id="${expected#exercises/expected/}"
  id="${id%.txt}"

  if [[ -n "$FILTER" && "$id" != *"$FILTER"* ]]; then
    continue
  fi

  if (( FROM_SOLUTIONS )); then
    source_file="exercises/solutions/${id}.rl"
  else
    source_file="exercises/${id}.rl"
  fi

  total=$((total + 1))
  actual="$("$BIN" "$source_file" 2>&1 | strip_final_value)"
  want="$(cat "$expected")"

  if [[ "$actual" == "$want" ]]; then
    printf '  \033[32m✓\033[0m %s\n' "$id"
    passed=$((passed + 1))
    continue
  fi

  failed=$((failed + 1))
  printf '  \033[31m✗\033[0m %s\n\n' "$id"
  printf '    edit %s\n\n' "$source_file"
  diff <(printf '%s\n' "$want") <(printf '%s\n' "$actual") \
    | sed 's/^</    expected:/; s/^>/    got:     /; /^---$/d; /^[0-9]/d'
  echo

  if (( ! ALL )); then
    printf '  %d of %d done. Fix the one above, then run this again.\n' "$passed" "$(find exercises/expected -name '*.txt' | wc -l | tr -d ' ')"
    exit 1
  fi
done

echo
if (( failed == 0 )); then
  printf '  \033[32mall %d solved\033[0m\n' "$passed"
  exit 0
fi

printf '  %d solved, %d to go (of %d)\n' "$passed" "$failed" "$total"
exit 1
