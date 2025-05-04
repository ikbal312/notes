package main

type SignedInt interface {
	int8 | int16
}

func sortInts[I SignedInt](slice []I) {

}

func f[K comparable, V any](m map[K]V) {}
