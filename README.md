## go-lru
> LRU =  Least Recently Used

Simpl, fast LRU cache algorithm

Inspired by the [hashlru algorithm](https://github.com/dominictarr/hashlru#algorithm).

## Usage
```
l := lru.New(10)

l.Has(1) // false
l.Set(1, 2)
v, ok := l.Get(1) // 1, true
l.Has(1) // true
```