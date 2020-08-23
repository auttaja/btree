package btree

import (
	"github.com/andersfylling/snowflake"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tree := NewBinaryTree()
	val := tree.Find(1)
	if val != nil {
		t.Fatal("tree is not empty!")
	}
}

func TestBinaryTree_Insert(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(&TreeValue{
		key:  snowflake.NewSnowflake(1),
		item: true,
	})

	if !tree.root.Value.item.(bool) {
		t.Fatal("did not insert properly")
	}
}

func TestBinaryTree_Find(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(&TreeValue{
		key:  snowflake.NewSnowflake(1),
		item: true,
	})

	if !tree.root.Value.item.(bool) {
		t.Fatal("did not insert properly")
	}

	if !tree.Find(snowflake.NewSnowflake(1)).(bool) {
		t.Fatal("failed to find key")
	}
}

func TestBinaryTree_Delete(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(&TreeValue{
		key:  snowflake.NewSnowflake(1),
		item: true,
	})

	if !tree.root.Value.item.(bool) {
		t.Fatal("did not insert properly")
	}

	tree.Delete(snowflake.NewSnowflake(1))

	if tree.Find(snowflake.NewSnowflake(1)) != nil {
		t.Fatal("found key after delete")
	}
}

func TestBinaryTree_Delete2(t *testing.T) {
	tree := NewBinaryTree()
	for i := 0; i < 100; i++ {
		tree.Insert(&TreeValue{
			key:  snowflake.NewSnowflake(uint64(i)),
			item: true,
		})
	}

	tree.Delete(snowflake.NewSnowflake(50))

	if tree.Find(snowflake.NewSnowflake(50)) != nil {
		t.Fatal("found key after delete")
	}
}

func BenchmarkBinaryTree_Insert(b *testing.B) {
	tree := NewBinaryTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(&TreeValue{
			key:  snowflake.Snowflake(i),
			item: true,
		})
	}
}

func BenchmarkBinaryTree_Find(b *testing.B) {
	tree := NewBinaryTree()
	for i := 0; i < b.N; i++ {
		tree.Insert(&TreeValue{
			key:  snowflake.Snowflake(i),
			item: true,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Find(snowflake.Snowflake(i))
	}
}

func BenchmarkBinaryTree_Delete(b *testing.B) {
	tree := NewBinaryTree()
	for i := 0; i < b.N; i++ {
		tree.Insert(&TreeValue{
			key:  snowflake.Snowflake(i),
			item: true,
		})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Delete(snowflake.Snowflake(i))
	}
}
