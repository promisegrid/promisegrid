package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// diOwnerRE matches DI owner lines, not ordinary references to DI IDs. This
// distinction lets DRs and docs cite a DI without being treated as owning a
// second copy of that handle.
var diOwnerRE = regexp.MustCompile(`(?m)^ID:\s*DI-(` +
	`[bdfghjklmnprstvz][aiou][bdfghjklmnprstvz][aiou][bdfghjklmnprstvz]` +
	`(?:-[bdfghjklmnprstvz][aiou][bdfghjklmnprstvz][aiou][bdfghjklmnprstvz])?` +
	`)\s*$`)

// scanCorpus returns the occupied handle set under repoRoot. It is deliberately
// path-agnostic: every file and directory name in the working tree reserves any
// proquint-shaped substring it contains, and every exact DI owner line reserves
// its DI handle.
//
// scanCorpus makes mint-handle a whole-working-tree occupied-set guard rather
// than a layout-specific artifact parser, so future paths cannot accidentally
// hide a handle-looking name from collision checks.
func scanCorpus(repoRoot string) (map[string]string, error) {
	handles := make(map[string]string)
	root, err := filepath.Abs(repoRoot)
	if err != nil {
		return nil, err
	}
	if err := scanPathNames(root, handles); err != nil {
		return nil, err
	}
	if err := scanDIHandles(root, handles); err != nil {
		return nil, err
	}
	return handles, nil
}

// scanPathNames walks the whole working tree and records every proquint-shaped
// substring in each visited entry name. The scan uses entry basenames, not full
// relative paths, so an ancestor directory handle-looking name is recorded once
// for that directory instead of once for every descendant.
//
// This avoids hard-coded path allowlists; any future repo path can reserve a
// candidate handle simply by containing that handle-looking string.
func scanPathNames(repoRoot string, handles map[string]string) error {
	return filepath.WalkDir(repoRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return fmt.Errorf("walk %s: %w", path, walkErr)
		}
		if path == repoRoot {
			return nil
		}
		if entry.Name() == ".git" {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		relPath, err := filepath.Rel(repoRoot, path)
		if err != nil {
			return fmt.Errorf("relpath %s: %w", path, err)
		}
		rememberNameHandles(handles, entry.Name(), filepath.ToSlash(relPath))
		return nil
	})
}

// scanDIHandles walks readable regular files across the working tree and
// records exact DI owner lines. It intentionally does not treat arbitrary
// proquint-looking body text as occupied, because ordinary prose may cite or
// discuss handles without owning them.
//
// This keeps DI handles in the same occupied set as name-derived handles
// without reintroducing TODO-directory path assumptions.
func scanDIHandles(repoRoot string, handles map[string]string) error {
	return filepath.WalkDir(repoRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return fmt.Errorf("walk %s: %w", path, walkErr)
		}
		if path == repoRoot {
			return nil
		}
		if entry.Name() == ".git" {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if !entry.Type().IsRegular() {
			return nil
		}
		relPath, err := filepath.Rel(repoRoot, path)
		if err != nil {
			return fmt.Errorf("relpath %s: %w", path, err)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read %s: %w", filepath.ToSlash(relPath), err)
		}
		for _, match := range diOwnerRE.FindAllStringSubmatch(string(data), -1) {
			rememberHandle(handles, match[1], filepath.ToSlash(relPath)+"#DI-"+match[1])
		}
		return nil
	})
}

// rememberNameHandles records all proquint-1 and proquint-2 substrings in one
// file or directory name. The loops advance one byte at a time so overlapping
// CVCVC substrings are not missed.
func rememberNameHandles(handles map[string]string, name, relPath string) {
	for start := 0; start+5 <= len(name); start++ {
		candidate := name[start : start+5]
		if isProquint1String(candidate) {
			owner := fmt.Sprintf("%s#name[%d:%d]", relPath, start, start+5)
			rememberHandle(handles, candidate, owner)
		}
	}
	for start := 0; start+11 <= len(name); start++ {
		candidate := name[start : start+11]
		if isProquint2String(candidate) {
			owner := fmt.Sprintf("%s#name[%d:%d]", relPath, start, start+11)
			rememberHandle(handles, candidate, owner)
		}
	}
}

// isProquint1String reports whether candidate is exactly one CVCVC proquint
// using the local alphabet in proquint.go.
func isProquint1String(candidate string) bool {
	return len(candidate) == 5 &&
		strings.ContainsRune(proquintCons, rune(candidate[0])) &&
		strings.ContainsRune(proquintVows, rune(candidate[1])) &&
		strings.ContainsRune(proquintCons, rune(candidate[2])) &&
		strings.ContainsRune(proquintVows, rune(candidate[3])) &&
		strings.ContainsRune(proquintCons, rune(candidate[4]))
}

// isProquint2String reports whether candidate is exactly two proquint-1 words
// joined by a single hyphen.
func isProquint2String(candidate string) bool {
	return len(candidate) == 11 &&
		candidate[5] == '-' &&
		isProquint1String(candidate[:5]) &&
		isProquint1String(candidate[6:])
}

// rememberHandle inserts one occupied handle into the corpus. Repeated
// occurrences are expected under whole-name substring scanning, so the first
// owner is kept as useful diagnostic context and later owners are ignored.
func rememberHandle(handles map[string]string, handle, owner string) {
	if _, duplicate := handles[handle]; duplicate {
		return
	}
	handles[handle] = owner
}
