package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEkiConcierge(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EkiConcierge Suite")
}
