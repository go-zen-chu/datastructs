package bst

import (
	"reflect"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want *BinarySearchTree
	}{
		{
			name: "BinaryTree should be created",
			args: args{val: 3},
			want: &BinarySearchTree{root: &TreeNode{Val: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinarySearchTree(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_Add_and_Print(t *testing.T) {
	t.Run("Add values to BinaryTree and print as a string", func(t *testing.T) {
		bt := NewBinarySearchTree(5)
		bt.Add(3)
		bt.Add(4)
		bt.Add(7)
		bt.Add(6)
		bt.Add(9)

		want := `└ 5
 └ 3
  └ 4
 └ 7
  └ 6
  └ 9
`
		if got := bt.String(); got != want {
			t.Errorf("Adding value to BinaryTree\ngot:\n%s\nwant:\n%s\n", got, want)
		}
	})
}

func TestBinaryTree_Is_Same(t *testing.T) {
	t.Run("Check binary tree is same", func(t *testing.T) {
		bt1 := NewBinarySearchTree(5)
		bt1.Add(1)
		bt1.Add(7)
		bt1.Add(6)
		bt2 := NewBinarySearchTree(5)
		bt2.Add(7)
		bt2.Add(6)
		bt2.Add(1)
		bt3 := NewBinarySearchTree(5)
		bt3.Add(7)
		bt3.Add(1)
		bt3.Add(6)
		nbt := NewBinarySearchTree(1)
		nbt.Add(7)
		nbt.Add(5)
		nbt.Add(6)

		if bt1.Equal(bt2) == false {
			t.Errorf("BinaryTree should be equal:\n%s\n%s\n", bt1.String(), bt2.String())
		}
		if bt1.Equal(bt3) == false {
			t.Errorf("BinaryTree should be equal:\n%s\n%s\n", bt1.String(), bt3.String())
		}
		if bt1.Equal(nbt) == true {
			t.Errorf("BinaryTree should not be equal:\n%s\n%s\n", bt1.String(), nbt.String())
		}
	})
}
