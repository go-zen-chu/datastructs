package bst

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want *BinarySearchTree
	}{
		{
			name: "BinarySearchTree should be created",
			args: args{val: 3},
			want: &BinarySearchTree{root: &TreeNode{Val: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinarySearchTree(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinarySearchTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Add_Find_Delete(t *testing.T) {
	t.Run("Add, Find, Delete values in BST", func(t *testing.T) {
		bt := NewBinarySearchTree(50)
		for _, val := range []int{25, 100, 12, 37, 150, 6, 18, 31, 43, 9, 1, 8, 40, 48, 125, 135} {
			bt.Add(val)
		}
		// Print current tree
		fmt.Println(bt.String())

		// Find test
		wantVal := 40
		if got := bt.Find(wantVal); got == nil || got.Val != wantVal {
			t.Errorf("Should get %d, but got %v", wantVal, got)
		}
		wantVal = 42
		if got := bt.Find(wantVal); got != nil {
			t.Errorf("Should not get %d, but got %v", wantVal, got)
		}

		// Delete test
		delVal := 42
		if got := bt.Delete(delVal); got != false {
			t.Errorf("Should not find %d, but got %v", delVal, got)
		}
		for _, delVal := range []int{8, 1, 25, 125, 100, 50, 37, 9, 40, 135, 48, 31, 6, 150, 43, 12, 18} {
			if got := bt.Delete(delVal); got != true {
				t.Errorf("Should delete %d, but got %v", delVal, got)
			}
			fmt.Println(bt.String())
		}
	})
}

func TestBinarySearchTree_Add_and_Print(t *testing.T) {
	t.Run("Add values to BinarySearchTree and print as a string", func(t *testing.T) {
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
			t.Errorf("Adding value to BinarySearchTree\ngot:\n%s\nwant:\n%s\n", got, want)
		}
	})
}

func TestBinarySearchTree_Is_Same(t *testing.T) {
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
			t.Errorf("BinarySearchTree should be equal:\n%s\n%s\n", bt1.String(), bt2.String())
		}
		if bt1.Equal(bt3) == false {
			t.Errorf("BinarySearchTree should be equal:\n%s\n%s\n", bt1.String(), bt3.String())
		}
		if bt1.Equal(nbt) == true {
			t.Errorf("BinarySearchTree should not be equal:\n%s\n%s\n", bt1.String(), nbt.String())
		}
	})
}
