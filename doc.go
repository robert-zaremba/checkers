/*
Package checkers is an extension to github.com/go-check/check library.
It provides additional usefull checkers:

  * Between - check if a number is between given 2 other numbers
  * Contains (check if a slice/array/string contains specified element)
  * CloseTo - an alias for EqualsWithTolerance
  * IsIn (check if an element is in a slice/array/string)
  * DoesNotExist
  * DurationLessThan
  * WithinDuration
  * EqualsWithTolerance - check if two numbers are "close enough"
  * HasPrefix, HasSuffix
  * IsDirectory
  * IsEmpty - check if specified object is empty (nil, [], {}, "", 0)
  * IsNonEmptyFile
  * IsSymlink, SymlinkDoesNotExist
  * IsTrue, IsFalse
  * MapEquals - check if 2 maps contain the same elements
  * SameContents
  * SamePath
  * Satisfies
  * SliceEquals - check if 2 slices contain the same elements
*/
package checkers
