package counter

import (
	"testing"
)

func TestCounter_Down(t *testing.T) {
	tests := map[string]struct {
		c    *Counter
		want int64
	}{
		"over 0": {
			func() *Counter {
				c := New()
				c.Set(20)
				return c
			}(),
			19,
		},
		"0 when Down from 1": {
			func() *Counter {
				c := New()
				c.Set(1)
				return c
			}(),
			0,
		},
		"0 when Down from 0": {
			func() *Counter {
				c := New()
				c.Set(0)
				return c
			}(),
			0,
		},
		"0 when Down from -1": {
			func() *Counter {
				c := New()
				c.Set(-1)
				return c
			}(),
			0,
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := tt.c.Down(); got != tt.want {
				t.Errorf("Counter.Down() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_Subtract(t *testing.T) {
	type args struct {
		val int64
	}
	tests := map[string]struct {
		c    *Counter
		args args
		want int64
	}{
		"over 0": {
			func() *Counter {
				c := New()
				c.Set(20)
				return c
			}(),
			args{10},
			10,
		},
		"0 when 10 - 10": {
			func() *Counter {
				c := New()
				c.Set(10)
				return c
			}(),
			args{10},
			0,
		},
		"0 when 1 - 10": {
			func() *Counter {
				c := New()
				c.Set(1)
				return c
			}(),
			args{10},
			0,
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := tt.c.Subtract(tt.args.val); got != tt.want {
				t.Errorf("Counter.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCounter_Set(t *testing.T) {
	type args struct {
		v int64
	}
	tests := map[string]struct {
		c    *Counter
		args args
		want int64
	}{
		"10 when Set(10)": {
			New(),
			args{10},
			10,
		},
		"0 when Set(0)": {
			New(),
			args{0},
			0,
		},
		"0 when Set(-1)": {
			New(),
			args{-1},
			0,
		},
	}
	for name, tt := range tests {
		tt := tt
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			tt.c.Set(tt.args.v)
			if got := tt.c.Get(); got != tt.want {
				t.Errorf("Counter.Set() set to %v, want %v", got, tt.want)
			}
		})
	}
}
