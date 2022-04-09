package parallel

// ParallelOption provides some useful configure options for parallel, such as conurrency limit.
type ParallelOption struct {
	concurrency int
}

type optFn func(*ParallelOption)

// WithConcurrency set the maximum number of concurrent `iteratee` goroutines running at the same time.
func WithConcurrency(concurrency int) optFn {
	return func(po *ParallelOption) {
		po.concurrency = concurrency
	}
}

func mergeOptions(optionFns []optFn) ParallelOption {
	opts := ParallelOption{
		concurrency: 0,
	}

	for _, fn := range optionFns {
		fn(&opts)
	}

	return opts
}
