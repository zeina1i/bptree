package main

type BTreeInterface interface {
	Set(key any, value any) (val any)
	Remove(key any) (val any)
	Get(key any) (val any)
}
