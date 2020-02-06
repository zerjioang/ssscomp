package main

import (
	"container/heap"
	"crypto/subtle"
	"log"
	"testing"
	"time"
)

var (
	letters = []byte("1234567890abcdefghijklmnopqrstuvwxyz")
)

type TestRun struct {
	Time int64
	Byte byte
}

type Times []TestRun

func (t Times) Len() int           { return len(t) }
func (t Times) Less(i, j int) bool { return t[i].Time > t[j].Time }
func (t Times) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func (t *Times) Push(v interface{}) {
	*t = append(*t, v.(TestRun))
}

func (t *Times) Pop() interface{} {
	a := *t
	n := len(a)
	v := a[n-1]
	*t = a[0 : n-1]
	return v
}

type Compare func(x, y []byte) int

func BrokenCompare(x, y []byte) int {
	for i := range x {
		if x[i] != y[i] {
			return 0
		}
	}
	return 1
}

func Crack(password []byte, comp Compare) []byte {
	n := len(password)
	guess := make([]byte, n)
	for index := range password {
		times := make(Times, 0)
		for _, letter := range letters {
			guess[index] = letter
			result := testing.Benchmark(func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					comp(password, guess)
				}
			})
			heap.Push(&times, TestRun{
				Time: result.NsPerOp(),
				Byte: letter,
			})
			log.Printf("took %s (%d ns/op) to try %q for index %d", result.T, result.NsPerOp(), letter, index)
		}
		tr := heap.Pop(&times).(TestRun)
		guess[index] = tr.Byte
		log.Printf("best guess is %q for index %d", tr.Byte, index)
		log.Printf("guess is now: %s", guess)
	}
	return guess
}

func ConstantTimeCrack(pw []byte) []byte {
	return Crack(pw, subtle.ConstantTimeCompare)
}

func BrokenCrack(pw []byte) []byte {
	return Crack(pw, BrokenCompare)
}

func main() {
	pw := []byte("cr4ck")
	start := time.Now()
	guess := BrokenCrack(pw)
	end := time.Now()
	dur := end.Sub(start)
	log.Printf("password guess after %s is: %s", dur, guess)
}
