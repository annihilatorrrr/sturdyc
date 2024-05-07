package sturdyc_test

import (
	"testing"
	"time"

	"github.com/creativecreature/sturdyc"
)

func TestPanicsIfCapacityIsLessThanZero(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when capacity is less than zero")
		}
	}()
	sturdyc.New[string](-1, 1, time.Hour, 5)
}

func TestPanicsIfTheNumberOfShardsIsLessThanOne(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use zero shards")
		}
	}()
	sturdyc.New[string](100, 0, time.Hour, 5)
}

func TestPanicsIfTTLIsLessThanOne(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use zero as TTL")
		}
	}()
	sturdyc.New[string](100, 2, 0, 5)
}

func TestPanicsIfEvictionPercentageIsLessThanZero(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use -1 as eviction percentage")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, -1)
}

func TestPanicsIfEvictionPercentageIsGreaterThanOneHundred(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use 101 as eviction percentage")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 101)
}

func TestPanicsIfRefreshBufferingIsEnabledWithoutStampedeProtection(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use buffered refreshes without stampede protection")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithRefreshBuffering(10, time.Minute),
	)
}

func TestPanicsIfTheRefreshBufferSizeIsLessThanOne(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use buffered refreshes with zero as size")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithStampedeProtection(time.Minute, time.Hour, time.Second, true),
		sturdyc.WithRefreshBuffering(0, time.Minute),
	)
}

func TestPanicsIfTheRefreshBufferTimeoutIsLessThanOne(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use 0 as buffer timeout")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithStampedeProtection(time.Minute, time.Hour, time.Second, true),
		sturdyc.WithRefreshBuffering(10, 0),
	)
}

func TestPanicsIfTheEvictionIntervalIsSetToLessThanOne(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use 0 as eviction interval")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithEvictionInterval(0),
	)
}

func TestPanicsIfTheMinRefreshTimeIsGreaterThanTheMaxRefreshTime(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use hour as min refresh time and minute as max")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithStampedeProtection(time.Hour, time.Minute, time.Second, true),
	)
}

func TestPanicsIfTheRetryBaseDelayIsLessThanZero(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use -1 as retry base delay")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithStampedeProtection(time.Minute, time.Hour, -1, true),
	)
}

func TestPanicsIfPassthroughPercentageIsLessThanOne(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use -1 as passthrough percentage")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithPassthroughPercentage(-1),
	)
}

func TestPanicsIfPassthroughPercentageIsGreaterThanOneHundred(t *testing.T) {
	t.Parallel()

	defer func() {
		err := recover()
		if err == nil {
			t.Error("expected a panic when trying to use 101 as passthrough percentage")
		}
	}()
	sturdyc.New[string](100, 10, time.Minute, 5,
		sturdyc.WithPassthroughPercentage(101),
	)
}
