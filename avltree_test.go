package btree

import (
	"github.com/andersfylling/snowflake"
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	tree := NewAVLTree()
	val := tree.Find(1)
	if val != nil {
		t.Fatal("tree is not empty!")
	}
}

func TestAVLTree_Insert(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(&TreeValue{
		key:  snowflake.NewSnowflake(1),
		item: true,
	})

	if !tree.root.Value.item.(bool) {
		t.Fatal("did not insert properly")
	}
}

func TestAVLTree_Find(t *testing.T) {
	tree := NewAVLTree()
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

func TestAVLTree_Delete(t *testing.T) {
	tree := NewAVLTree()
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

func TestAVLTree_Delete2(t *testing.T) {
	tree := NewAVLTree()
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

func BenchmarkAVLTree_Insert(b *testing.B) {
	tree := NewAVLTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(&TreeValue{
			key:  snowflake.Snowflake(i),
			item: true,
		})
	}
}

func BenchmarkAVLTree_Find(b *testing.B) {
	tree := NewAVLTree()
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

func BenchmarkAVLTree_Delete(b *testing.B) {
	tree := NewAVLTree()
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
