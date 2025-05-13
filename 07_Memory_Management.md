# Memory Management in Go

Memory allocation and deallocation happens automatically in Go.

## Two Methods of Memory Management

### make()

- allocates memory , initializes slices, maps, and channels and returns a value of the type not a pointer.
- The syntax is `make(Type, size, capacity)` for slices and `make(Type)` for maps and channels.
- Gives non-zeroed storage meaning the memory is allocated and initialized.

- Example:
  ```go
  slice := make([]int, 5) // creates a slice of int with length 5
  m := make(map[string]int) // creates a map with string keys and int values
  ch := make(chan int) // creates a channel of int
  ```

### new()

- allocates memory, no initialization, for arrays and structs and returns a pointer to the type.
- The syntax is `new(Type)` for arrays and structs.
- Gives zeroed storage meaning the memory is allocated but not initialized.
- Example:
  ```go
  slice := new([]int)
  m := new(map[string]int)
  ch := new(chan int)
  ```

## Important Notes

- Garbage collection is automatic in Go, meaning the programmer does not need to manually free memory.
- OutofScope or Nil variables are automatically garbage collected.
