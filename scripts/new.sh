#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 ]]; then
  echo "Usage: $0 <slug> [func_name]" >&2
  echo "Example: $0 p0123_best_time_to_buy_sell maxProfit" >&2
  exit 1
fi

slug="$1"
func="${2:-solve}"

# Convert function name to CamelCase for exported functions
# First letter should be uppercase
exported_func=$(echo "${func}" | sed 's/^\(.\)/\U\1/')

repo_root="$(cd "$(dirname "$0")/.." && pwd)"
problems_dir="$repo_root/problems"
problem_dir="$problems_dir/${slug}"
impl_file="$problem_dir/${slug}.go"
test_file="$problem_dir/${slug}_test.go"

mkdir -p "$problem_dir"

if [[ -e "$impl_file" ]] || [[ -e "$test_file" ]]; then
  echo "Refusing to overwrite existing files in: $problem_dir" >&2
  exit 2
fi

# Create the implementation file from template
cat > "$impl_file" << EOF
// __ID__ __TITLE__
// Link: __URL__
// Pattern: __PATTERN__
// Invariants: __INVARIANTS__

package ${slug}

// ${exported_func} solves the problem
// Add parameters and return types according to the problem requirements
func ${exported_func}() {
	// Implementation goes here
}
EOF

# Create the test file
cat > "$test_file" << EOF
package ${slug}

import (
	"testing"
)

func Test${exported_func}(t *testing.T) {
	// Add test cases here
	// Example:
	// result := ${exported_func}(/* input */)
	// expected := /* expected */
	// if result != expected {
	//     t.Errorf("Expected %v, got %v", expected, result)
	// }
}
EOF

# Replace placeholders in the implementation file
sed -i '' \
  -e "s/__ID__/$(echo "$slug" | sed 's/^p\([0-9]\+\).*/\1/')/" \
  -e "s/__TITLE__/$(echo "$slug" | sed 's/^p[0-9]\+_//; s/_/ /g')/" \
  -e "s|__URL__|https://leetcode.com/problems/$(echo "$slug" | sed 's/^p[0-9]\+_//')/|" \
  -e "s/__PATTERN__/TBD/" \
  -e "s/__INVARIANTS__/TBD/" \
  -e "s/__FUNC__/${exported_func}/g" \
  -e "s/__PKG_NAME__/${slug}/g" \
  "$impl_file"

# Format the files with gofmt
gofmt -w "$impl_file"
gofmt -w "$test_file"

echo "Created $problem_dir"
echo "Files created:"
echo "  $impl_file"
echo "  $test_file"
echo ""
echo "Next: edit signatures, fill implementations and tests, then run:"
echo "  go test ./$slug"