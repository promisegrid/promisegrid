// Command mint-handle allocates a unique proquint handle for a new
// coordination artifact.
//
// Usage:
//
//	mint-handle [-w 1|2] [-r REPO_ROOT] [-n] [-s SEED]
//
// Flags:
//
//	-w   handle width: 1 for proquint-1 (5 chars, default) or 2 for
//	     proquint-2 (11 chars). proquint-1 is the canonical width for
//	     new TODO, TE, DR, and DI artifacts.
//	-r   repo root (default: current directory). The tool walks the whole
//	     working tree from this root, excluding only .git internals, to build
//	     the occupied handle set.
//	-n   dry-run: print what would be minted without contacting the
//	     clock multiple times. Mostly useful for tests.
//	-s   override the entropy seed source. By default the seed is
//	     time.Now().UnixNano() at the moment mint decides it needs a
//	     fresh attempt. -s SEED forces the first attempt to use SEED;
//	     subsequent retries fall back to the clock. Useful for tests
//	     and for deterministic re-mints under a fixed seed.
//
// Output: a single line on stdout containing the chosen handle and
// nothing else, e.g.
//
//	vapoj
//
// Or for proquint-2:
//
//	gapip-munav
//
// Exit code 0 on success, non-zero on any error, including failure to
// scan the corpus or exhausting the requested proquint space.
//
// Why this design:
//
//   - Handles are mint-time-allocated random labels, not derivable from
//     filename or content. They are persisted by being included in file names,
//     directory names, or DI owner lines. There is no central HANDLES.md.
//   - The check is "does this proquint-looking string already appear in a
//     working-tree file or directory name, or as an exact DI owner line?" This
//     whole-tree occupied-set scan avoids hard-coded layout blind spots. Forks
//     diverge naturally because each fork's occupied set is distinct; the
//     collision check is local and merge-time, not a central registry.
//   - The entropy source is the wall clock at mint time. The hash function
//     (SHA-256) folds that into 16 or 32 bits. The choice of entropy source
//     is operationally secondary; the collision check is what guarantees
//     uniqueness within the scanned corpus.
//
// The tool provides one local, collision-checked minting path for all new
// coordination artifact IDs.
package main

import (
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var (
		width    = flag.Int("w", 1, "handle width: 1 (proquint-1) or 2 (proquint-2)")
		repoRoot = flag.String("r", ".", "repo root")
		dryRun   = flag.Bool("n", false, "dry-run (use seed once, no clock retries)")
		seed     = flag.Int64("s", 0, "entropy seed override; 0 means use clock")
	)
	flag.Parse()

	if *width != 1 && *width != 2 {
		fmt.Fprintf(os.Stderr, "mint-handle: -w must be 1 or 2, got %d\n", *width)
		os.Exit(2)
	}

	corpus, err := scanCorpus(*repoRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mint-handle: %v\n", err)
		os.Exit(1)
	}

	handle, err := mint(*width, corpus, *seed, *dryRun)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mint-handle: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(handle)
}

// mint returns a fresh handle of the requested width that is not present
// in corpus. It draws entropy from time.Now().UnixNano(), folds through
// SHA-256, and retries on collision.
//
// If seed != 0 the first attempt uses seed as the entropy source. In
// dryRun mode no retries are performed; the function returns whatever the
// first attempt produced even on collision, returning an error in that
// collision case so callers know the dry-run hit an occupied handle.
//
// In production mode the function retries up to maxAttempts before giving
// up. At 50% saturation of proquint-1, about 32,000 handles, average
// retries are 2; at 99% saturation, about 65,000 handles, average retries
// are 100; at 100% saturation no handle exists and an error is returned.
func mint(width int, corpus map[string]string, seed int64, dryRun bool) (string, error) {
	const maxAttempts = 1_000_000
	for attempt := 0; attempt < maxAttempts; attempt++ {
		var ns int64
		if attempt == 0 && seed != 0 {
			ns = seed
		} else {
			ns = time.Now().UnixNano()
		}
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(ns))
		sum := sha256.Sum256(buf[:])

		var candidate string
		switch width {
		case 1:
			candidate = proquint1FromBytes(sum[:2])
		case 2:
			candidate = proquint2FromBytes(sum[:4])
		}

		if _, taken := corpus[candidate]; !taken {
			return candidate, nil
		}
		if dryRun {
			return "", fmt.Errorf("dry-run: first attempt %q collides with corpus", candidate)
		}

		// The short sleep advances coarse clocks enough that the next
		// time-derived seed is different on systems with low timestamp
		// resolution, while keeping normal minting latency negligible.
		time.Sleep(time.Microsecond)
	}
	return "", fmt.Errorf(
		"exhausted %d attempts; corpus may be saturated (%d handles in use)",
		maxAttempts, len(corpus))
}
