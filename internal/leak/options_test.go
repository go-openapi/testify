package leak

/*
func TestOptionsFilters(t *testing.T) {
	opts := buildOpts()
	cur := stack.Current()
	all := getStableAll(t, cur)

	// At least one of these should be the same as current, the others should be filtered out.
	for _, s := range all {
		if s.ID() == cur.ID() {
			require.False(t, opts.filter(s), "Current test running function should not be filtered")
		} else {
			require.True(t, opts.filter(s), "Default goroutines should be filtered: %v", s)
		}
	}

	defer startBlockedG().unblock()

	// Now the filters should find something that doesn't match a filter.
	countUnfiltered := func() int {
		var unmatched int
		for _, s := range stack.All() {
			if s.ID() == cur.ID() {
				continue
			}
			if !opts.filter(s) {
				unmatched++
			}
		}
		return unmatched
	}
	require.Equal(t, 1, countUnfiltered(), "Expected blockedG goroutine to not match any filter")

	// If we add an extra filter to ignore blockTill, it shouldn't match.
	opts = buildOpts(IgnoreTopFunction("go.uber.org/goleak.(*blockedG).block"))
	require.Zero(t, countUnfiltered(), "blockedG should be filtered out. running: %v", stack.All())

	// If we ignore startBlockedG, that should not ignore the blockedG goroutine
	// because startBlockedG should be the "created by" function in the stack.
	opts = buildOpts(IgnoreAnyFunction("go.uber.org/goleak.startBlockedG"))
	require.Equal(t, 1, countUnfiltered(),
		"startBlockedG should not be filtered out. running: %v", stack.All())
}

func TestOptionsIgnoreCreatedBy(t *testing.T) {
	stopCh := make(chan struct{})
	go func() {
		<-stopCh
	}()
	defer close(stopCh)

	cur := stack.Current()
	opts := buildOpts(IgnoreCreatedBy("go.uber.org/goleak.TestOptionsIgnoreCreatedBy"))

	for _, s := range stack.All() {
		if s.ID() == cur.ID() {
			continue
		}

		if opts.filter(s) {
			continue
		}

		t.Errorf("Unexpected goroutine: %v", s)
	}
}

func TestOptionsIgnoreAnyFunction(t *testing.T) {
	cur := stack.Current()
	opts := buildOpts(IgnoreAnyFunction("go.uber.org/goleak.(*blockedG).run"))

	for _, s := range stack.All() {
		if s.ID() == cur.ID() {
			continue
		}

		if opts.filter(s) {
			continue
		}

		t.Errorf("Unexpected goroutine: %v", s)
	}
}

func TestOptionsRetry(t *testing.T) {
	opts := buildOpts()
	opts.maxRetries = 50 // initial attempt + 50 retries = 11
	opts.maxSleep = time.Millisecond

	for i := 0; i < 50; i++ {
		assert.True(t, opts.retry(i), "Attempt %v/51 should allow retrying", i)
	}
	assert.False(t, opts.retry(51), "Attempt 51/51 should not allow retrying")
	assert.False(t, opts.retry(52), "Attempt 52/51 should not allow retrying")
}
*/
