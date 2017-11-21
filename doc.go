/*
Package checkers is an extension to github.com/go-check/check library.
It provides additional usefull checkers:

  * Between - checks if a number is between given 2 other numbers
  * Contains (checks if a slice/array/string contains specified element)
  * CloseTo - an alias for EqualsWithTolerance
  * IsIn (checks if an element is in a slice/array/string)
  * DoesNotExist
  * DurationLessThan
  * WithinDuration
  * EqualsWithTolerance - checks if two numbers are "close enough"
  * HasPrefix, HasSuffix
  * IsDirectory
  * IsEmpty - checks if specified object is empty (nil, [], {}, "", 0)
  * IsNonEmptyFile
  * IsSymlink, SymlinkDoesNotExist
  * IsTrue, IsFalse
  * MapEquals - checks if 2 maps contain the same elements
  * SameContents
  * SamePath
  * Satisfies
  * SliceEquals - checks if 2 slices contain the same elements
  * StrEquals - checks if fmt.Sprint values of objects are equal
*/
package checkers
