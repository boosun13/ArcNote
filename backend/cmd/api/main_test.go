package main

import "testing"

func TestServerAddrUsesDefaultPort(t *testing.T) {
	t.Setenv("PORT", "")

	if got := serverAddr(); got != ":8080" {
		t.Fatalf("serverAddr() = %q, want %q", got, ":8080")
	}
}

func TestServerAddrUsesEnvironmentPort(t *testing.T) {
	t.Setenv("PORT", "9090")

	if got := serverAddr(); got != ":9090" {
		t.Fatalf("serverAddr() = %q, want %q", got, ":9090")
	}
}
