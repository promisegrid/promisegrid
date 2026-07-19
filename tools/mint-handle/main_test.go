package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestMintWidthOne checks that proquint-1 mints are 5 chars and pass the
// proquint shape regex.
func TestMintWidthOne(t *testing.T) {
	got, err := mint(1, map[string]string{}, 0, false)
	if err != nil {
		t.Fatalf("mint: %v", err)
	}
	if len(got) != 5 {
		t.Fatalf("len(%q) = %d, want 5", got, len(got))
	}
	if !isProquint1String(got) {
		t.Fatalf("mint %q is not proquint-1", got)
	}
}

// TestMintWidthTwo checks that proquint-2 mints are 11 chars (5+1+5) and
// contain exactly one hyphen.
func TestMintWidthTwo(t *testing.T) {
	got, err := mint(2, map[string]string{}, 0, false)
	if err != nil {
		t.Fatalf("mint: %v", err)
	}
	if len(got) != 11 {
		t.Fatalf("len(%q) = %d, want 11", got, len(got))
	}
	if strings.Count(got, "-") != 1 {
		t.Fatalf("hyphen count in %q = %d, want 1", got, strings.Count(got, "-"))
	}
	if !isProquint2String(got) {
		t.Fatalf("mint %q is not proquint-2", got)
	}
}

// TestMintAvoidsCollision pre-populates the corpus with the value that
// seed=42 would produce on first attempt, then mints with a clock-based
// retry. The returned handle must be different from the seeded one, proving
// that the retry path actually fires.
func TestMintAvoidsCollision(t *testing.T) {
	corpus := map[string]string{}
	first, err := mint(1, corpus, 42, false)
	if err != nil {
		t.Fatalf("seed mint: %v", err)
	}
	corpus[first] = "fake/path.md"
	got, err := mint(1, corpus, 42, false)
	if err != nil {
		t.Fatalf("mint after collision: %v", err)
	}
	if got == first {
		t.Fatalf("mint returned colliding handle %q", got)
	}
}

// TestMintDryRunCollidesErrors verifies dry-run is deterministic and reports
// a collision instead of retrying with wall-clock entropy.
func TestMintDryRunCollidesErrors(t *testing.T) {
	first, err := mint(1, map[string]string{}, 42, true)
	if err != nil {
		t.Fatalf("seed dry-run: %v", err)
	}
	_, err = mint(1, map[string]string{first: "fake/path.md"}, 42, true)
	if err == nil {
		t.Errorf("dry-run with seeded collision: want error, got nil")
	}
}

// TestScanCorpusEmpty verifies an empty repo scan returns an empty corpus
// instead of a nil-map error.
func TestScanCorpusEmpty(t *testing.T) {
	root := t.TempDir()
	corpus, err := scanCorpus(root)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if len(corpus) != 0 {
		t.Fatalf("corpus size = %d, want 0", len(corpus))
	}
}

// TestScanCorpusScansArbitraryNestedNames proves the scanner does not depend
// on particular coordination directories or filename prefixes.
func TestScanCorpusScansArbitraryNestedNames(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "docs/thought-experiments/TE-vapoj-substrate.md", "# placeholder\n")
	writeFile(t, root, "arbitrary/not-a-coordination-place/prefixjodonpostfix.md", "# placeholder\n")
	writeFile(t, root, "nested/bahor-nisam.txt", "# placeholder\n")
	writeFile(t, root, "notes/plain.md", "# placeholder\n\nID: DI-dafut\n")
	makeDir(t, root, "scratch/prefixpilohpostfix")

	corpus, err := scanCorpus(root)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	for _, handle := range []string{"vapoj", "jodon", "bahor", "nisam", "bahor-nisam", "dafut", "piloh"} {
		if _, ok := corpus[handle]; !ok {
			t.Errorf("missing handle %q in %v", handle, corpus)
		}
	}
}

// TestScanCorpusAllowsDuplicateNameReservations verifies duplicate ordinary
// substrings make a handle occupied without making the repo unscannable.
func TestScanCorpusAllowsDuplicateNameReservations(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "alpha/manifest.json", "{}\n")
	writeFile(t, root, "beta/manifest.json", "{}\n")

	corpus, err := scanCorpus(root)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if _, ok := corpus["manif"]; !ok {
		t.Fatalf("missing duplicate-reserved handle manif in %v", corpus)
	}
}

// TestScanCorpusIgnoresGitDirectory verifies .git internals are the only
// explicit path exclusion from the working-tree name scan.
func TestScanCorpusIgnoresGitDirectory(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, ".git/objects/hidden/TODO-vapoj-hidden.md", "# placeholder\n")
	writeFile(t, root, ".git/objects/hidden/di.md", "ID: DI-bahor\n")
	writeFile(t, root, "visible/keep-jodon.txt", "# placeholder\n")

	corpus, err := scanCorpus(root)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if _, ok := corpus["vapoj"]; ok {
		t.Fatalf(".git filename handle leaked into corpus: %v", corpus)
	}
	if _, ok := corpus["bahor"]; ok {
		t.Fatalf(".git DI handle leaked into corpus: %v", corpus)
	}
	if _, ok := corpus["jodon"]; !ok {
		t.Fatalf("visible handle missing from corpus: %v", corpus)
	}
}

// TestScanCorpusFindsDIOwnerLinesEverywhere proves DI owner lines participate
// in the occupied set without requiring a TODO directory.
func TestScanCorpusFindsDIOwnerLinesEverywhere(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "research/plain-note.md", "# placeholder\n\nID: DI-fonuz\n")

	corpus, err := scanCorpus(root)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if _, ok := corpus["fonuz"]; !ok {
		t.Fatalf("DI owner handle missing from corpus: %v", corpus)
	}
}

func makeDir(t *testing.T, root, path string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(root, path), 0o755); err != nil {
		t.Fatal(err)
	}
}

func writeFile(t *testing.T, root, path, contents string) {
	t.Helper()
	full := filepath.Join(root, path)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(contents), 0o644); err != nil {
		t.Fatal(err)
	}
}
