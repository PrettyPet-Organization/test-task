package repository

import (
	"small_warehouse/internal/model"
	"testing"
)

func TestStore(t *testing.T) {
	st := New[model.Good]()

	good := model.NewGood("", "")

	st.Put(good.Id.String(), *good)
	if len(st.m) != 1 {
		t.Fatal("put has failed")
	}

	st.Delete(good.Id.String())
	if len(st.m) != 0 {
		t.Fatal("deletion has failed")
	}

	st.Delete("key")
}
