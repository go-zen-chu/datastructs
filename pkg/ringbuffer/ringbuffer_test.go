package ringbuffer

import (
	"reflect"
	"testing"
)

func TestNewRingBuffer(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name    string
		args    args
		want    RingBuffer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "check ring buffer can be created",
			args:    args{len: 5},
			want:    &ringBuffer{buf: []int{0, 0, 0, 0, 0}, len: 5, startIdx: 0, endIdx: 0},
			wantErr: false},
		{
			name:    "check ring buffer can't be created with invalid args",
			args:    args{len: 0},
			want:    nil,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRingBuffer(tt.args.len)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRingBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRingBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ringBuffer_Len(t *testing.T) {
	type fields struct {
		buf      []int
		len      int
		startIdx int
		endIdx   int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name:   "check ring buffer len is 5",
			fields: fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 3, endIdx: 2},
			want:   5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ringBuffer{
				buf:      tt.fields.buf,
				len:      tt.fields.len,
				startIdx: tt.fields.startIdx,
				endIdx:   tt.fields.endIdx,
			}
			if got := r.Len(); got != tt.want {
				t.Errorf("ringBuffer.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ringBuffer_IsEmpty(t *testing.T) {
	type fields struct {
		buf      []int
		len      int
		startIdx int
		endIdx   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "check ring buffer is not empty",
			fields: fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 3, endIdx: 2},
			want:   false},
		{
			name:   "check ring buffer is empty",
			fields: fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 2, endIdx: 2},
			want:   true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ringBuffer{
				buf:      tt.fields.buf,
				len:      tt.fields.len,
				startIdx: tt.fields.startIdx,
				endIdx:   tt.fields.endIdx,
			}
			if got := r.IsEmpty(); got != tt.want {
				t.Errorf("ringBuffer.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ringBuffer_IsFull(t *testing.T) {
	type fields struct {
		buf      []int
		len      int
		startIdx int
		endIdx   int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "check ring buffer is full",
			fields: fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 2, endIdx: 1},
			want:   true},
		{
			name:   "check ring buffer is not full",
			fields: fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 2, endIdx: 0},
			want:   false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ringBuffer{
				buf:      tt.fields.buf,
				len:      tt.fields.len,
				startIdx: tt.fields.startIdx,
				endIdx:   tt.fields.endIdx,
			}
			if got := r.IsFull(); got != tt.want {
				t.Errorf("ringBuffer.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ringBuffer_ShowQueue(t *testing.T) {
	type fields struct {
		buf      []int
		len      int
		startIdx int
		endIdx   int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "check ring buffer's queue data",
			fields: fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 2, endIdx: 0},
			want:   []int{3, 4, 5, 1}},
		{
			name:   "check ring buffer's queue data when empty",
			fields: fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 2, endIdx: 2},
			want:   []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ringBuffer{
				buf:      tt.fields.buf,
				len:      tt.fields.len,
				startIdx: tt.fields.startIdx,
				endIdx:   tt.fields.endIdx,
			}
			if got := r.ShowQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ringBuffer.ShowQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ringBuffer_Enqueue(t *testing.T) {
	type fields struct {
		buf      []int
		len      int
		startIdx int
		endIdx   int
	}
	type args struct {
		data int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "check ring buffer can enqueue",
			fields:  fields{buf: []int{0, 0, 0, 0, 0}, len: 5, startIdx: 0, endIdx: 0},
			args:    args{data: 1},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ringBuffer{
				buf:      tt.fields.buf,
				len:      tt.fields.len,
				startIdx: tt.fields.startIdx,
				endIdx:   tt.fields.endIdx,
			}
			if err := r.Enqueue(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ringBuffer.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_ringBuffer_Dequeue(t *testing.T) {
	type fields struct {
		buf      []int
		len      int
		startIdx int
		endIdx   int
	}
	tests := []struct {
		name     string
		fields   fields
		wantData int
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     "check ring buffer can dequeue",
			fields:   fields{buf: []int{1, 2, 3, 4, 5}, len: 5, startIdx: 2, endIdx: 4},
			wantData: 3,
			wantErr:  false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ringBuffer{
				buf:      tt.fields.buf,
				len:      tt.fields.len,
				startIdx: tt.fields.startIdx,
				endIdx:   tt.fields.endIdx,
			}
			gotData, err := r.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("ringBuffer.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotData != tt.wantData {
				t.Errorf("ringBuffer.Dequeue() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
