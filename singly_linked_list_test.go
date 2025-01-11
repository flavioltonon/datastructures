package datastructures

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSinglyLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *SinglyLinkedList[int]
	}{
		{
			name: "It should create a new SinglyLinkedList with the type that has been declared",
			want: new(SinglyLinkedList[int]),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSinglyLinkedList[int]()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSinglyLinkedList_PrependNode(t *testing.T) {
	type fields struct {
		head   *Node[int]
		length int
	}

	type args struct {
		node *Node[int]
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SinglyLinkedList[int]
	}{
		{
			name: "It should insert the new node at the head of the linked list",
			fields: fields{
				head: &Node[int]{
					key:   "b",
					value: 2,
					next:  nil,
				},
				length: 1,
			},
			args: args{
				node: &Node[int]{
					key:   "a",
					value: 1,
				},
			},
			want: &SinglyLinkedList[int]{
				head: &Node[int]{
					key:   "a",
					value: 1,
					next: &Node[int]{
						key:   "b",
						value: 2,
						next:  nil,
					},
				},
				length: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList[int]{
				head:   tt.fields.head,
				length: tt.fields.length,
			}

			l.PrependNode(tt.args.node)

			require.Equal(t, tt.want, l)
		})
	}
}

func TestSinglyLinkedList_AppendNode(t *testing.T) {
	type fields struct {
		head   *Node[int]
		length int
	}

	type args struct {
		node *Node[int]
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SinglyLinkedList[int]
	}{
		{
			name: "It should insert the new node at the tail of the linked list",
			fields: fields{
				head: &Node[int]{
					key:   "b",
					value: 2,
					next:  nil,
				},
				length: 1,
			},
			args: args{
				node: &Node[int]{
					key:   "a",
					value: 1,
				},
			},
			want: &SinglyLinkedList[int]{
				head: &Node[int]{
					key:   "b",
					value: 2,
					next: &Node[int]{
						key:   "a",
						value: 1,
						next:  nil,
					},
				},
				length: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList[int]{
				head:   tt.fields.head,
				length: tt.fields.length,
			}

			l.AppendNode(tt.args.node)

			require.Equal(t, tt.want, l)
		})
	}
}

func TestSinglyLinkedList_DeleteNode(t *testing.T) {
	type fields struct {
		head   *Node[int]
		length int
	}

	type args struct {
		key string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SinglyLinkedList[int]
	}{
		{
			name: "If the key belongs to the head node, it should delete that node",
			fields: fields{
				head: &Node[int]{
					key:   "b",
					value: 2,
					next:  nil,
				},
				length: 1,
			},
			args: args{
				key: "b",
			},
			want: &SinglyLinkedList[int]{
				head:   nil,
				length: 0,
			},
		},
		{
			name: "If the key belongs to the tail node, it should delete that node",
			fields: fields{
				head: &Node[int]{
					key:   "a",
					value: 1,
					next: &Node[int]{
						key:   "b",
						value: 2,
						next:  nil,
					},
				},
				length: 2,
			},
			args: args{
				key: "b",
			},
			want: &SinglyLinkedList[int]{
				head: &Node[int]{
					key:   "a",
					value: 1,
					next:  nil,
				},
				length: 1,
			},
		},
		{
			name: "If the key does not belong to any nodes, no changes should be made to the linked list",
			fields: fields{
				head: &Node[int]{
					key:   "a",
					value: 1,
					next:  nil,
				},
				length: 1,
			},
			args: args{
				key: "b",
			},
			want: &SinglyLinkedList[int]{
				head: &Node[int]{
					key:   "a",
					value: 1,
					next:  nil,
				},
				length: 1,
			},
		},
		{
			name: "If the key belongs to more than one node, only the first one should be deleted",
			fields: fields{
				head: &Node[int]{
					key:   "a",
					value: 1,
					next: &Node[int]{
						key:   "b",
						value: 2,
						next: &Node[int]{
							key:   "a",
							value: 3,
							next:  nil,
						},
					},
				},
				length: 3,
			},
			args: args{
				key: "a",
			},
			want: &SinglyLinkedList[int]{
				head: &Node[int]{
					key:   "b",
					value: 2,
					next: &Node[int]{
						key:   "a",
						value: 3,
						next:  nil,
					},
				},
				length: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &SinglyLinkedList[int]{
				head:   tt.fields.head,
				length: tt.fields.length,
			}

			l.DeleteNode(tt.args.key)

			require.Equal(t, tt.want, l)
		})
	}
}
