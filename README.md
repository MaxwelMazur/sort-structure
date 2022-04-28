# sort-structure

> simple and incredible as the ordering of its structure can make a difference


```go 
type carefreeStructure struct {
	Int   bool    // 1 byte
	Float float64 // 8 bytes
	Bool  int32   // 4 bytes
}

```
  | bytes | description |
  |---|---|
  | 1000-0000 | 1 byte for bool |  
  | 1111-1111 | 8 byte for float64 |
  | 1111-0000 | 4 byte for int32 |
  
---

```go 
type worriedStructure struct {
	Float float64 // 8 bytes
	Int   int32   // 4 bytes
	Bool  bool    // 1 byte
}
```
  | bytes | description |
  |---|---|
  | 1111-1111 | 8 byte for float64 |
  | 1000-1111 |  1 byte for bool + 4 byte for int32 |
  
  <br>

> for more details: https://medium.com/@felipedutratine/how-to-organize-the-go-struct-in-order-to-save-memory-c78afcf59ec2
