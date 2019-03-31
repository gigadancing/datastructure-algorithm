package linkedlist

// 460. All O'one Data Structure
// Implement a data structure supporting the following operations:
//
// Inc(Key) - Inserts a new key with value 1. Or increments an existing key by 1. Key is guaranteed to be a non-empty string.
// Dec(Key) - If Key's value is 1, remove it from the data structure. Otherwise decrements an existing key by 1.
//            If the key does not exist, this function does nothing. Key is guaranteed to be a non-empty string.
// GetMaxKey() - Returns one of the keys with maximal value. If no element exists, return an empty string "".
// GetMinKey() - Returns one of the keys with minimal value. If no element exists, return an empty string "".
// Challenge: Perform all these in O(1) time complexity.

type AllOne struct {
}

// Initialize your data structure here.
func AllOneConstructor() AllOne {
	return AllOne{}
}

// Inserts a new key <Key> with value 1. Or increments an existing key by 1.
func (this *AllOne) Inc(key string) {

}

// Decrements an existing key by 1. If Key's value is 1, remove it from the data structure.
func (this *AllOne) Dec(key string) {

}

// Returns one of the keys with maximal value.
func (this *AllOne) GetMaxKey() string {
	return ""
}

// Returns one of the keys with Minimal value.
func (this *AllOne) GetMinKey() string {
	return ""
}
