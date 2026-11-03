package param

// Ptr returns a pointer to v. Use when no type-specific helper exists.
func Ptr[T any](v T) *T { return &v }

// String returns a pointer to v.
func String(v string) *string { return &v }

// Int64 returns a pointer to v.
func Int64(v int64) *int64 { return &v }

// Float64 returns a pointer to v.
func Float64(v float64) *float64 { return &v }

// Bool returns a pointer to v.
func Bool(v bool) *bool { return &v }
