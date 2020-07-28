package main

import "testing"

const (
	IMAGE_HOST_HTTPS_BOOK  = "book"
	IMAGE_HOST_HTTPS_TING  = "book"
	IMAGE_HOST_HTTPS_USER  = "book"
	IMAGE_HOST_HTTPS_OTHER = "book"
	IMAGE_HOST_HTTP_BOOK   = "book"
	IMAGE_HOST_HTTP_TING   = "book"
	IMAGE_HOST_HTTP_USER   = "book"
	IMAGE_HOST_HTTP_OTHER  = "book"
)

func getHost(Scheme, namespace string) (path string) {
	if Scheme == "https" {
		switch namespace {
		case "book":
			path = IMAGE_HOST_HTTPS_BOOK
		case "ting":
			path = IMAGE_HOST_HTTPS_TING
		case "user":
			path = IMAGE_HOST_HTTPS_USER
		case "other":
			path = IMAGE_HOST_HTTPS_OTHER
		default:
			path = IMAGE_HOST_HTTPS_BOOK
		}
	} else {
		switch namespace {
		case "book":
			path = IMAGE_HOST_HTTP_BOOK
		case "ting":
			path = IMAGE_HOST_HTTP_TING
		case "user":
			path = IMAGE_HOST_HTTP_USER
		case "other":
			path = IMAGE_HOST_HTTP_OTHER
		default:
			path = IMAGE_HOST_HTTP_BOOK
		}
	}
	return
}

func getHostV2(Scheme, namespace string) string {
	if Scheme == "https" {
		data := map[string]string{
			"book":  IMAGE_HOST_HTTPS_BOOK,
			"ting":  IMAGE_HOST_HTTPS_TING,
			"user":  IMAGE_HOST_HTTPS_USER,
			"other": IMAGE_HOST_HTTPS_OTHER,
		}
		if path, ok := data[namespace]; ok {
			return path
		}
		return IMAGE_HOST_HTTPS_BOOK
	} else {
		data := map[string]string{
			"book":  IMAGE_HOST_HTTPS_BOOK,
			"ting":  IMAGE_HOST_HTTPS_TING,
			"user":  IMAGE_HOST_HTTPS_USER,
			"other": IMAGE_HOST_HTTPS_OTHER,
		}
		if path, ok := data[namespace]; ok {
			return path
		}
		return IMAGE_HOST_HTTPS_BOOK
	}
}

func BenchmarkHost(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getHost("https", "book")
	}
}

func BenchmarkHostV2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = getHostV2("https", "book")
	}
}
