package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want Queue
	}{
		{
			name: "check queue can be created",
			args: args{cap: 5},
			want: &queue{data: make([]int, 0, 5), size: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Queue(t *testing.T) {
	t.Run("check queue works as expected", func(t *testing.T) {
		q := NewQueue(5)
		if !q.IsEmpty() {
			t.Error("queue supposed to be empty\n")
		}
		q.Enqueue(4)
		q.Enqueue(321514)
		q.Enqueue(329023)
		q.Enqueue(0)
		if q.IsEmpty() {
			t.Error("queue supposed not to be empty\n")
		}
		isDequeued, val := q.Dequeue()
		wantVal := 4
		if !isDequeued {
			t.Error("should be dequeued\n")
		}
		if val != wantVal {
			t.Errorf("dequeued value supposed to be %d, gt %d\n", wantVal, val)
		}
		q.Dequeue()
		q.Dequeue()
		q.Dequeue()
		isDequeued, _ = q.Dequeue()
		if isDequeued {
			t.Error("queue should return nothing because queue is empty\n")
		}
	})
}

func Test_queue_String(t *testing.T) {
	type fields struct {
		data []int
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "check queue stringer works",
			fields: fields{
				data: []int{1, 2, 3, 4},
				size: 4,
			},
			want: "[1 2 3 4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
				size: tt.fields.size,
			}
			if got := q.String(); got != tt.want {
				t.Errorf("queue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
