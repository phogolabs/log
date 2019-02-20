package log

import "context"

// The MIT License (MIT)

// Copyright (c) 2018 Go Playground

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

var (
	// WriterCtxKey is the context.Context key to store the request log entry.
	WriterCtxKey = &contextKey{"writer"}
)

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type contextKey struct {
	name string
}

// String represent the context key
func (k *contextKey) String() string {
	return "log: context value " + k.name
}

// SetContext sets a log entry into the provided context
func SetContext(ctx context.Context, e Writer) context.Context {
	return context.WithValue(ctx, WriterCtxKey, e)
}

// GetContext returns the log Entry found in the context,
// or a new Default log Entry if none is found
func GetContext(ctx context.Context) Writer {
	v := ctx.Value(WriterCtxKey)

	if v == nil {
		v = std
	}

	return v.(Writer)
}
