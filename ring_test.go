package ring

import (
	// "fmt"
	"testing"
)

func TestSetsSize(t *testing.T) {
	r := &Ring{}
	r.SetCapacity(10)
	if r.Capacity() != 10 {
		t.Fatal("Size of ring was not 10", r.Capacity())
	}
}

func TestSavesSomeData(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 7; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		r.Enqueue(val)
	}
	for i := 0; i < 7; i++ {
		x := r.Dequeue()
		val := make([]string, 1)
		val[0] = string(i)
		if x[0] != val[0] {
			t.Fatal("Unexpected response", x, "wanted", val)
		}
	}
}

func TestReusesBuffer(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 7; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		r.Enqueue(val)
	}
	for i := 0; i < 7; i++ {
		r.Dequeue()
	}
	for i := 7; i < 14; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		r.Enqueue(val)
	}
	for i := 7; i < 14; i++ {
		x := r.Dequeue()
		val := make([]string, 1)
		val[0] = string(i)
		if x[0] != val[0] {
			t.Fatal("Unexpected response", x, "wanted", val)
		}
	}
}

func TestOverflowsBuffer(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 20; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		r.Enqueue(val)
	}
	for i := 10; i < 20; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		x := r.Dequeue()
		if x[0] != val[0] {
			t.Fatal("Unexpected response", x, "wanted", val)
		}
	}
}

func TestPartiallyOverflows(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 15; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		r.Enqueue(val)
	}
	for i := 5; i < 15; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		x := r.Dequeue()
		if x[0] != val[0] {
			t.Fatal("Unexpected response", x, "wanted", val)
		}
	}
}

func TestPeeks(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	for i := 0; i < 10; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		r.Enqueue(val)
	}
	for i := 0; i < 10; i++ {
		r.Peek()
		r.Peek()
		x1 := r.Peek()
		x := r.Dequeue()
		val := make([]string, 1)
		val[0] = string(i)
		if x[0] != val[0] {
			t.Fatal("Unexpected response", x, "wanted", val)
		}
		if x1[0] != x[0] {
			t.Fatal("Unexpected response", x1, "wanted", x)
		}
	}
}

func TestConstructsArr(t *testing.T) {
	r := Ring{}
	r.SetCapacity(10)
	v := r.Values()
	if len(v) != 0 {
		t.Fatal("Unexpected values", v, "wanted len of", 0)
	}
	for i := 1; i < 21; i++ {
		val := make([]string, 1)
		val[0] = string(i)
		r.Enqueue(val)
		l := i
		if l > 10 {
			l = 10
		}
		v = r.Values()
		if len(v) != l {
			t.Fatal("Unexpected values", v, "wanted len of", l, "index", i)
		}
	}
}
