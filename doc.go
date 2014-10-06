/*
Package checkers is an extension to github.com/go-check/check library.
It provides additional usefull checkers:

  * Contains (check if a slice/array/string contains specified element)
  * EqualsWithTolerance - check if two numbers are "close enough"
  * IsTrue
  * IsEmpty - check if specified object is empty (nil, [], {}, "", 0)
  * SliceEquals - check if 2 slices contain the same elements
  * MapEquals - check if 2 maps contain the same elements
*/
package checkers
