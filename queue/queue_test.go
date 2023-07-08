package queue

import "testing"

func TestCreateQueue(t *testing.T) {
	q := CreateQueue()
	if q == nil {
		t.Error("CreateQueue() returned nil queue")
	}
}

func TestEnqueue(t *testing.T) {
	q := CreateQueue()
	for i := 0; i < 20; i++ {
		q.Enqueue(i)
	}
}

func TestDequeue(t *testing.T) {
	q := CreateQueue()
	q.Enqueue(1)
	if q.Dequeue() != 1 {
		t.Error("Dequeue() failed")
	}
}

func TestQueue(t *testing.T) {
	q := CreateQueue()
	for i := 0; i < 20; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 20; i++ {
		val := q.Dequeue()
		if val != i {
			t.Error("expected ", i, " got ", val)
		}
	}
}

func TestEnd(t *testing.T) {
	q := CreateQueue()
	if q.NotEmpty() {
		t.Error("Queue is not empty, expected empty")
	}
	q.Enqueue(0)
	if !q.NotEmpty() {
		t.Error("Queue is empty, expected is not empty")
	}
}
