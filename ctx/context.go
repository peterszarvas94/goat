package ctx

import (
	"context"
	"net/http"
)

type KV map[string]any

func Add(r *http.Request, items KV) *http.Request {
	var newR = r
	for key, value := range items {
		ctx := context.WithValue(newR.Context(), key, value)
		newR = newR.WithContext(ctx)
	}
	return newR
}

func Get[T any](r *http.Request, key string) (*T, bool) {
	val := r.Context().Value(key)
	if typedVal, ok := val.(*T); ok && typedVal != nil {
		return typedVal, true
	}
	return nil, false
}

func Delete(r *http.Request, key string) *http.Request {
	ctx := context.WithValue(context.Background(), key, nil)
	return r.WithContext(ctx)
}
