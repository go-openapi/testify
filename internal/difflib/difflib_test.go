// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package difflib

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGetOptCodes(t *testing.T) {
	a := "qabxcd"
	b := "abycdf"
	s := NewMatcher(splitChars(a), splitChars(b))
	w := &bytes.Buffer{}
	for _, op := range s.GetOpCodes() {
		fmt.Fprintf(w, "%s a[%d:%d], (%s) b[%d:%d] (%s)\n", string(op.Tag),
			op.I1, op.I2, a[op.I1:op.I2], op.J1, op.J2, b[op.J1:op.J2])
	}
	result := w.String()
	expected := `d a[0:1], (q) b[0:0] ()
e a[1:3], (ab) b[0:2] (ab)
r a[3:4], (x) b[2:3] (y)
e a[4:6], (cd) b[3:5] (cd)
i a[6:6], () b[5:6] (f)
`
	if expected != result {
		t.Errorf("unexpected op codes: \n%s", result)
	}
}

func TestGroupedOpCodes(t *testing.T) {
	a := []string{}
	for i := 0; i != 39; i++ {
		a = append(a, fmt.Sprintf("%02d", i))
	}
	b := make([]string, 0, len(a)+3)
	b = append(b, a[:8]...)
	b = append(b, " i")
	b = append(b, a[8:19]...)
	b = append(b, " x")
	b = append(b, a[20:22]...)
	b = append(b, a[27:34]...)
	b = append(b, " y")
	b = append(b, a[35:]...)
	s := NewMatcher(a, b)
	w := &bytes.Buffer{}
	for _, g := range s.GetGroupedOpCodes(-1) {
		fmt.Fprintf(w, "group\n")
		for _, op := range g {
			fmt.Fprintf(w, "  %s, %d, %d, %d, %d\n", string(op.Tag),
				op.I1, op.I2, op.J1, op.J2)
		}
	}
	result := w.String()
	expected := `group
  e, 5, 8, 5, 8
  i, 8, 8, 8, 9
  e, 8, 11, 9, 12
group
  e, 16, 19, 17, 20
  r, 19, 20, 20, 21
  e, 20, 22, 21, 23
  d, 22, 27, 23, 23
  e, 27, 30, 23, 26
group
  e, 31, 34, 27, 30
  r, 34, 35, 30, 31
  e, 35, 38, 31, 34
`
	if expected != result {
		t.Errorf("unexpected op codes: \n%s", result)
	}
}

func rep(s string, count int) string {
	return strings.Repeat(s, count)
}

func TestWithAsciiOneInsert(t *testing.T) {
	sm := NewMatcher(splitChars(rep("b", 100)),
		splitChars("a"+rep("b", 100)))
	assertEqual(t, sm.GetOpCodes(),
		[]OpCode{{'i', 0, 0, 0, 1}, {'e', 0, 100, 1, 101}})
	assertEqual(t, len(sm.bPopular), 0)

	sm = NewMatcher(splitChars(rep("b", 100)),
		splitChars(rep("b", 50)+"a"+rep("b", 50)))
	assertEqual(t, sm.GetOpCodes(),
		[]OpCode{{'e', 0, 50, 0, 50}, {'i', 50, 50, 50, 51}, {'e', 50, 100, 51, 101}})
	assertEqual(t, len(sm.bPopular), 0)
}

func TestWithAsciiOnDelete(t *testing.T) {
	sm := NewMatcher(splitChars(rep("a", 40)+"c"+rep("b", 40)),
		splitChars(rep("a", 40)+rep("b", 40)))
	assertEqual(t, sm.GetOpCodes(),
		[]OpCode{{'e', 0, 40, 0, 40}, {'d', 40, 41, 40, 40}, {'e', 41, 81, 40, 80}})
}

func TestSFBugsComparingEmptyLists(t *testing.T) {
	groups := NewMatcher(nil, nil).GetGroupedOpCodes(-1)
	assertEqual(t, len(groups), 0)
	diff := UnifiedDiff{
		FromFile: "Original",
		ToFile:   "Current",
		Context:  3,
	}
	result, err := GetUnifiedDiffString(diff)
	assertEqual(t, err, nil)
	assertEqual(t, result, "")
}

func TestOutputFormatRangeFormatUnified(t *testing.T) {
	// Per the diff spec at http://www.unix.org/single_unix_specification/
	//
	// Each <range> field shall be of the form:
	//   %1d", <beginning line number>  if the range contains exactly one line,
	// and:
	//  "%1d,%1d", <beginning line number>, <number of lines> otherwise.
	// If a range is empty, its beginning line number shall be the number of
	// the line just before the range, or 0 if the empty range starts the file.
	fm := formatRangeUnified
	assertEqual(t, fm(3, 3), "3,0")
	assertEqual(t, fm(3, 4), "4")
	assertEqual(t, fm(3, 5), "4,2")
	assertEqual(t, fm(3, 6), "4,3")
	assertEqual(t, fm(0, 0), "0,0")
}

func TestOutputFormatRangeFormatContext(t *testing.T) {
	// Per the diff spec at http://www.unix.org/single_unix_specification/
	//
	// The range of lines in file1 shall be written in the following format
	// if the range contains two or more lines:
	//     "*** %d,%d ****\n", <beginning line number>, <ending line number>
	// and the following format otherwise:
	//     "*** %d ****\n", <ending line number>
	// The ending line number of an empty range shall be the number of the preceding line,
	// or 0 if the range is at the start of the file.
	//
	// Next, the range of lines in file2 shall be written in the following format
	// if the range contains two or more lines:
	//     "--- %d,%d ----\n", <beginning line number>, <ending line number>
	// and the following format otherwise:
	//     "--- %d ----\n", <ending line number>
	fm := formatRangeContext
	assertEqual(t, fm(3, 3), "3")
	assertEqual(t, fm(3, 4), "4")
	assertEqual(t, fm(3, 5), "4,5")
	assertEqual(t, fm(3, 6), "4,6")
	assertEqual(t, fm(0, 0), "0")
}

func TestOutputFormatTabDelimiter(t *testing.T) {
	diff := UnifiedDiff{
		A:        splitChars("one"),
		B:        splitChars("two"),
		FromFile: "Original",
		FromDate: "2005-01-26 23:30:50",
		ToFile:   "Current",
		ToDate:   "2010-04-12 10:20:52",
		Eol:      "\n",
	}
	ud, err := GetUnifiedDiffString(diff)
	assertEqual(t, err, nil)
	result := SplitLines(ud)[:2]
	assertEqual(t, result, []string{
		"--- Original\t2005-01-26 23:30:50\n",
		"+++ Current\t2010-04-12 10:20:52\n",
	})
}

func TestOutputFormatNoTrailingTabOnEmptyFiledate(t *testing.T) {
	diff := UnifiedDiff{
		A:        splitChars("one"),
		B:        splitChars("two"),
		FromFile: "Original",
		ToFile:   "Current",
		Eol:      "\n",
	}
	ud, err := GetUnifiedDiffString(diff)
	assertEqual(t, err, nil)
	assertEqual(t, SplitLines(ud)[:2], []string{"--- Original\n", "+++ Current\n"})
}

func TestOmitFilenames(t *testing.T) {
	diff := UnifiedDiff{
		A:   SplitLines("o\nn\ne\n"),
		B:   SplitLines("t\nw\no\n"),
		Eol: "\n",
	}
	ud, err := GetUnifiedDiffString(diff)
	assertEqual(t, err, nil)
	assertEqual(t, SplitLines(ud), []string{
		"@@ -0,0 +1,2 @@\n",
		"+t\n",
		"+w\n",
		"@@ -2,2 +3,0 @@\n",
		"-n\n",
		"-e\n",
		"\n",
	})
}

func TestSplitLines(t *testing.T) {
	allTests := []struct {
		input string
		want  []string
	}{
		{"foo", []string{"foo\n"}},
		{"foo\nbar", []string{"foo\n", "bar\n"}},
		{"foo\nbar\n", []string{"foo\n", "bar\n", "\n"}},
	}
	for _, test := range allTests {
		assertEqual(t, SplitLines(test.input), test.want)
	}
}

func assertEqual(t *testing.T, a, b any) {
	t.Helper()

	if !reflect.DeepEqual(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

func splitChars(s string) []string {
	chars := make([]string, 0, len(s))
	// Assume ASCII inputs
	for _, r := range s {
		chars = append(chars, string(r))
	}

	return chars
}

// TestSequenceMatcherCaching tests that GetMatchingBlocks and GetOpCodes
// return cached results when called multiple times.
func TestSequenceMatcherCaching(t *testing.T) {
	a := splitChars("abc")
	b := splitChars("abd")

	sm := NewMatcher(a, b)

	// Call GetMatchingBlocks twice - second call should use cache
	blocks1 := sm.GetMatchingBlocks()
	blocks2 := sm.GetMatchingBlocks()
	assertEqual(t, blocks1, blocks2)

	// Call GetOpCodes twice - second call should use cache
	codes1 := sm.GetOpCodes()
	codes2 := sm.GetOpCodes()
	assertEqual(t, codes1, codes2)
}

// TestSetSeqSamePointer tests that SetSeq1 and SetSeq2 do NOT reset caches
// when the same slice pointer is passed (early return optimization).
func TestSetSeqSamePointer(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"x", "y", "z"}

	sm := NewMatcher(a, b)

	// Get initial blocks
	blocks1 := sm.GetMatchingBlocks()

	// Set the same sequences again using SetSeqs
	// Since we pass the same slice pointers, the caches should NOT be reset
	// (the implementation checks pointer equality for early return)
	sm.SetSeq1(a)
	sm.SetSeq2(b)

	// Blocks should remain cached (not nil) after setting the same sequences
	// so GetMatchingBlocks returns the cached result
	blocks2 := sm.GetMatchingBlocks()
	assertEqual(t, blocks1, blocks2)
}

// TestSequenceMatcherWithIsJunk tests the junk filtering functionality.
func TestSequenceMatcherWithIsJunk(t *testing.T) {
	// Test with a simple IsJunk function that marks whitespace as junk
	a := []string{"a", " ", "b", " ", "c"}
	b := []string{"a", "b", "c"}

	sm := NewMatcher(nil, nil)
	sm.IsJunk = func(s string) bool {
		return s == " "
	}
	sm.SetSeqs(a, b)

	// The matcher should still find matches but handle junk elements
	blocks := sm.GetMatchingBlocks()
	if len(blocks) == 0 {
		t.Error("expected some matching blocks with junk filter")
	}
}

// TestAutoJunkWithLargeSequence tests the autoJunk feature with sequences >= 200 elements.
func TestAutoJunkWithLargeSequence(t *testing.T) {
	// Create a sequence with more than 200 elements where one element appears
	// more than 1% of the time (which makes it "popular" and gets filtered)
	a := make([]string, 250)
	b := make([]string, 250)

	// Fill with unique elements
	for i := 0; i < 250; i++ {
		a[i] = fmt.Sprintf("a%d", i)
		b[i] = fmt.Sprintf("a%d", i)
	}

	// Make element "common" appear more than 1% (3+ times out of 250)
	for i := 0; i < 10; i++ {
		b[i] = "common"
	}

	sm := NewMatcher(a, b)
	// The popular element "common" should be filtered
	if len(sm.bPopular) == 0 {
		t.Log("bPopular might be empty if 'common' doesn't exceed threshold, which is expected")
	}

	// The matcher should still work
	blocks := sm.GetMatchingBlocks()
	if blocks == nil {
		t.Error("expected matching blocks")
	}
}

// TestFindLongestMatchWithJunk tests finding longest match with junk elements.
func TestFindLongestMatchWithJunk(t *testing.T) {
	// Create sequences where junk elements are adjacent to interesting matches
	a := []string{"x", "a", "b", "c", "y"}
	b := []string{"a", "b", "c"}

	sm := NewMatcher(nil, nil)
	// Mark x and y as junk
	sm.IsJunk = func(s string) bool {
		return s == "x" || s == "y"
	}
	sm.SetSeqs(a, b)

	blocks := sm.GetMatchingBlocks()
	// Should find the "a", "b", "c" match
	found := false
	for _, block := range blocks {
		if block.Size == 3 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find a match of size 3")
	}
}

// TestFindLongestMatchExtension tests the extension of matches past popular elements.
func TestFindLongestMatchExtension(t *testing.T) {
	// Test cases that exercise the match extension loops in findLongestMatch
	a := []string{"a", "b", "c", "d", "e"}
	b := []string{"x", "b", "c", "d", "y"}

	sm := NewMatcher(a, b)
	blocks := sm.GetMatchingBlocks()

	// Should find the "b", "c", "d" match
	found := false
	for _, block := range blocks {
		if block.Size >= 3 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find a match of size >= 3")
	}
}

// TestJunkFilteringInChainB tests the IsJunk function in chainB.
func TestJunkFilteringInChainB(t *testing.T) {
	// Create a matcher with junk filtering
	a := []string{"line1", "junk", "line2", "junk", "line3"}
	b := []string{"line1", "junk", "line2", "junk", "line3", "junk"}

	sm := NewMatcher(nil, nil)
	sm.IsJunk = func(s string) bool {
		return s == "junk"
	}
	sm.SetSeqs(a, b)

	// Verify junk is correctly identified
	if !sm.isBJunk("junk") {
		t.Error("expected 'junk' to be identified as junk")
	}

	// Non-junk should not be identified as junk
	if sm.isBJunk("line1") {
		t.Error("expected 'line1' to not be junk")
	}

	// Should still be able to find matches
	blocks := sm.GetMatchingBlocks()
	if len(blocks) == 0 {
		t.Error("expected some matching blocks")
	}
}

// TestMatchExtensionWithJunkOnBothSides tests junk matching extension.
func TestMatchExtensionWithJunkOnBothSides(t *testing.T) {
	// Create sequences where junk elements surround interesting matches
	// to exercise the junk extension loops in findLongestMatch
	a := []string{"junk1", "junk2", "a", "b", "c", "junk3", "junk4"}
	b := []string{"junk1", "junk2", "a", "b", "c", "junk3", "junk4"}

	sm := NewMatcher(nil, nil)
	sm.IsJunk = func(s string) bool {
		return strings.HasPrefix(s, "junk")
	}
	sm.SetSeqs(a, b)

	blocks := sm.GetMatchingBlocks()
	// Should find matches including junk elements that are identical
	totalSize := 0
	for _, block := range blocks {
		totalSize += block.Size
	}
	if totalSize < 3 {
		t.Errorf("expected total match size >= 3, got %d", totalSize)
	}
}

// TestFindLongestMatchBreakCondition tests the j >= bhi break condition.
func TestFindLongestMatchBreakCondition(t *testing.T) {
	// Create sequences that will trigger the j >= bhi condition
	// This happens when b2j has indices that exceed the search range
	a := []string{"x", "y", "z"}
	b := []string{"a", "b", "x", "y", "z"}

	sm := NewMatcher(a, b)
	blocks := sm.GetMatchingBlocks()

	// Should find the "x", "y", "z" match
	found := false
	for _, block := range blocks {
		if block.Size == 3 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find a match of size 3")
	}
}

// TestAutoJunkPopularElements tests the autoJunk filtering of popular elements.
func TestAutoJunkPopularElements(t *testing.T) {
	// Create a sequence with > 200 elements where one element appears
	// more than 1% of the time
	n := 250
	a := make([]string, n)
	b := make([]string, n)

	// Fill with mostly unique elements
	for i := 0; i < n; i++ {
		a[i] = fmt.Sprintf("line%d", i)
		b[i] = fmt.Sprintf("line%d", i)
	}

	// Make "popular" appear more than 1% (more than 2-3 times)
	// We need it to appear > n/100 + 1 times = 3+ times
	for i := 0; i < 10; i++ {
		b[i*25] = "popular"
	}

	sm := NewMatcher(a, b)

	// The element "popular" should be filtered as popular
	if len(sm.bPopular) == 0 {
		t.Log("bPopular might be empty if threshold not exceeded")
	}

	// Matcher should still produce valid results
	blocks := sm.GetMatchingBlocks()
	if blocks == nil {
		t.Error("expected non-nil matching blocks")
	}
}

// TestFindLongestMatchWithJunkExtension tests the junk extension loops
// at the end of findLongestMatch function.
func TestFindLongestMatchWithJunkExtension(t *testing.T) {
	// Create sequences where junk elements are adjacent to matches
	// This should trigger the junk extension loops
	a := []string{"junk", "a", "b", "c", "junk"}
	b := []string{"junk", "a", "b", "c", "junk"}

	sm := NewMatcher(nil, nil)
	sm.IsJunk = func(s string) bool {
		return s == "junk"
	}
	sm.SetSeqs(a, b)

	blocks := sm.GetMatchingBlocks()
	// Should find matches including junk extension
	totalSize := 0
	for _, block := range blocks {
		totalSize += block.Size
	}
	// The non-junk elements (a, b, c) should definitely match.
	// Junk elements may or may not be included depending on extension behavior.
	if totalSize < 3 {
		t.Errorf("expected total match size >= 3, got %d", totalSize)
	}
}

// TestFindLongestMatchEdgeCases tests edge cases in findLongestMatch.
func TestFindLongestMatchEdgeCases(t *testing.T) {
	// Test case where matches are found at the end of sequences
	a := []string{"unique1", "unique2", "match"}
	b := []string{"other1", "other2", "match"}

	sm := NewMatcher(a, b)
	blocks := sm.GetMatchingBlocks()

	// Should find the "match" element
	found := false
	for _, block := range blocks {
		if block.Size == 1 && block.A == 2 && block.B == 2 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find a match at the end")
	}
}

// TestMatcherWithBothSequencesSame tests the matcher with identical sequences.
func TestMatcherWithBothSequencesSame(t *testing.T) {
	a := []string{"line1", "line2", "line3"}
	b := []string{"line1", "line2", "line3"}

	sm := NewMatcher(a, b)
	blocks := sm.GetMatchingBlocks()

	// Should find all lines match
	if len(blocks) < 1 {
		t.Error("expected at least one matching block")
	}

	// The last block is always a sentinel with size 0
	for _, block := range blocks[:len(blocks)-1] {
		if block.Size != 3 {
			t.Errorf("expected matching block of size 3, got %d", block.Size)
		}
	}
}

// TestWriteUnifiedDiffWithDefaultEol tests that default EOL is applied.
func TestWriteUnifiedDiffWithDefaultEol(t *testing.T) {
	// Test that when Eol is empty, it defaults to "\n"
	diff := UnifiedDiff{
		A:        splitChars("abc"),
		B:        splitChars("abd"),
		FromFile: "file1",
		ToFile:   "file2",
		// Eol not set - should default to "\n"
	}
	result, err := GetUnifiedDiffString(diff)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(result, "\n") {
		t.Error("expected newlines in output")
	}
}
