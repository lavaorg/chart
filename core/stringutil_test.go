package chart

import (
	"testing"
)

func TestSplitCSV(t *testing.T) {
	// replaced new assertions helper

	gotest.AssertEmpty(t, SplitCSV(""))
	gotest.AssertEqual(t, []string{"foo"}, SplitCSV("foo"))
	gotest.AssertEqual(t, []string{"foo", "bar"}, SplitCSV("foo,bar"))
	gotest.AssertEqual(t, []string{"foo", "bar"}, SplitCSV("foo, bar"))
	gotest.AssertEqual(t, []string{"foo", "bar"}, SplitCSV(" foo , bar "))
	gotest.AssertEqual(t, []string{"foo", "bar", "baz"}, SplitCSV("foo,bar,baz"))
	gotest.AssertEqual(t, []string{"foo", "bar", "baz,buzz"}, SplitCSV("foo,bar,\"baz,buzz\""))
	gotest.AssertEqual(t, []string{"foo", "bar", "baz,'buzz'"}, SplitCSV("foo,bar,\"baz,'buzz'\""))
	gotest.AssertEqual(t, []string{"foo", "bar", "baz,'buzz"}, SplitCSV("foo,bar,\"baz,'buzz\""))
	gotest.AssertEqual(t, []string{"foo", "bar", "baz,\"buzz\""}, SplitCSV("foo,bar,'baz,\"buzz\"'"))
}
