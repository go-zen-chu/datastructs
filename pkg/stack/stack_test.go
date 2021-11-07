package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want Stack
	}{
		// TODO: Add test cases.
		{
			name: "check stack can be created",
			args: args{cap: 5},
			want: &stack{data: make([]int, 0, 5), size: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Stack(t *testing.T) {
	t.Run("check stack works as expected", func(t *testing.T) {
		s := NewStack(5)
		if !s.IsEmpty() {
			t.Error("stack supposed to be empty\n")
		}
		s.Push(10)
		s.Push(21301)
		s.Push(123429)
		s.Push(0)
		if s.IsEmpty() {
			t.Errorf("stack supporsed not to be empty\n")
		}
		top := s.Top()
		wantTop := 0
		if top != wantTop {
			t.Errorf("top value supposed to be %d, got %d\n", wantTop, top)
		}
		isPopped, pop := s.Pop()
		wantPop := 0
		if !isPopped {
			t.Error("should be popped\n")
		}
		if pop != wantPop {
			t.Errorf("popped value supposed to be %d, got %d\n", wantPop, pop)
		}
		s.Pop()
		s.Pop()
		s.Pop()
		isPopped, _ = s.Pop()
		if isPopped {
			t.Error(("pop should return nothing because stack is empty\n"))
		}
	})
}

func Test_stack_String(t *testing.T) {
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
			name: "check stack stringer works",
			fields: fields{
				data: []int{1, 2, 3, 4},
				size: 4,
			},
			want: "[1 2 3 4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &stack{
				data: tt.fields.data,
				size: tt.fields.size,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("stack.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
