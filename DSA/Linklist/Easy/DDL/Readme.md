# Doubly Linked List Implementation in Go

This file implements a basic doubly linked list (DDL) in Go. It supports the following operations:

- **Insert**: Add a new node to the end of the list.
- **Delete**: Remove a node by value, handling head, tail, and middle cases.
- **Print**: Display the list from head to tail.

## Usage

1. Create a new node and initialize the list:
   ```go
   newNode := &Node{data: 12}
   DDList := &DDList{head: newNode}
   ```
2. Insert nodes:
   ```go
   DDList.insert(20)
   DDList.insert(30)
   DDList.insert(40)
   DDList.insert(50)
   ```
3. Print the list:
   ```go
   DDList.print()
   ```
4. Delete a node:
   ```go
   DDList.deleteNode(20)
   DDList.deleteNode(40)
   ```

## Structure
- `Node`: Represents each element in the list, with `prev`, `data`, and `next` fields.
- `DDList`: The doubly linked list, with a pointer to the `head` node.

## Example Output
```
12 -->20 -->30 -->40 -->50 -->
12 -->30 -->40 -->50 -->
12 -->30 -->50 -->
```

## Notes
- Handles edge cases for deleting head, tail, and middle nodes.
- Prints messages when deleting the last element or the entire list.

---
For more details, see the code in `DDL.go`.