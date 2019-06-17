package server

import (
	"fmt"
	"testing"
)

func TestParseUserInfo(t *testing.T) {
	user := "eyJpZCI6NzUsInVzZXJuYW1lIjoid2VuemhlbmdsaW4iLCJlbWFpbCI6IndlbnpoZW5nbGluQGhhb2RhaS5uZXQiLCJuYW1lIjoid2VuemhlbmdsaW4iLCJzdGF0ZSI6ImFjdGl2ZSIsImNyZWF0ZWRfYXQiOiIyMDE4LTEyLTEwVDAzOjExOjQzLjQ3WiIsImJpbyI6IiIsImxvY2F0aW9uIjoiIiwicHVibGljX2VtYWlsIjoiIiwic2t5cGUiOiIiLCJsaW5rZWRpbiI6IiIsInR3aXR0ZXIiOiIiLCJ3ZWJzaXRlX3VybCI6IiIsIm9yZ2FuaXphdGlvbiI6IiIsImV4dGVybl91aWQiOiIiLCJwcm92aWRlciI6IiIsInRoZW1lX2lkIjoxLCJsYXN0X2FjdGl2aXR5X29uIjoiMjAxOS0wMi0yMCIsImNvbG9yX3NjaGVtZV9pZCI6MSwiaXNfYWRtaW4iOnRydWUsImF2YXRhcl91cmwiOiJodHRwczovL3d3dy5ncmF2YXRhci5jb20vYXZhdGFyLzhhOGYwNzkxMDZkZDFlNTcxYmJkYjQ2ZDA0OThhYWJlP3M9ODBcdTAwMjZkPWlkZW50aWNvbiIsImNhbl9jcmVhdGVfZ3JvdXAiOnRydWUsImNhbl9jcmVhdGVfcHJvamVjdCI6dHJ1ZSwicHJvamVjdHNfbGltaXQiOjEwMDAwMCwiY3VycmVudF9zaWduX2luX2F0IjoiMjAxOS0wMi0yMFQwNDoyNDozMi44NloiLCJsYXN0X3NpZ25faW5fYXQiOiIyMDE5LTAyLTIwVDAyOjAwOjU1LjI4MloiLCJjb25maXJtZWRfYXQiOiIyMDE4LTEyLTEwVDAzOjExOjQzLjQyM1oiLCJ0d29fZmFjdG9yX2VuYWJsZWQiOmZhbHNlLCJpZGVudGl0aWVzIjpbeyJwcm92aWRlciI6ImxkYXBtYWluIiwiZXh0ZXJuX3VpZCI6ImNuPXdlbnpoZW5nbGluLG91PeaKgOacr+S4reW/gyxvdT1wZW9wbGUsZGM9aGFvZGFpLGRjPW5ldCJ9XSwiZXh0ZXJuYWwiOmZhbHNlLCJwcml2YXRlX3Byb2ZpbGUiOmZhbHNlLCJzaGFyZWRfcnVubmVyc19taW51dGVzX2xpbWl0IjowfQ=="

	u, _, err := ParseUserInfo(user)
	if err != nil {
		t.Error("parse user err", err)
		return
	}
	fmt.Println("got user", u)
}