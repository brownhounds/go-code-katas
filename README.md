# Go Code Katas

Small Go code katas organized as stand-alone modules. This repo is a personal
practice space for algorithm drills, data-structure patterns, and core Go
syntax fluency.

## Why code katas

- Build speed and accuracy on common patterns.
- Isolate a single concept to practice it deeply with tests.
- Keep muscle memory fresh for interviews or real-world problem solving.
- Create a growing library of reference solutions you can revisit later.

### Slices / Arrays

- `two_sum`: Find two indices whose values add to a target using a hash map
(complement lookup).
- `valid_anagram`: Decide if two strings are anagrams using a frequency count.
- `rotate_slice_right`: Rotate a slice to the right by `k` steps in-place using
the three-reversal trick.
- `remove_duplicates`: Remove duplicates from a sorted slice in-place using a 
two-pointer pass; returns the new logical length.
- `move_zeros`: Move all zeros to the end of a slice in-place while preserving 
the order of non-zero elements (stable compaction).
- `maximum_subarray_sum`: Compute the max sum of any contiguous subarray using 
Kadane's algorithm.
- `product_of_array_except_self`: Build a result slice where each element is 
the product of all other elements using prefix/suffix products.

### Strings

- `valid_anagram`: Determine whether two strings are anagrams using character frequency counting.
- `longest_unique_substring`: Compute the length of the longest substring without repeating characters 
using a sliding window with last-seen index tracking.
- `minimum_window_substring`: Find the smallest contiguous substring containing all characters of a target
  string (including duplicates) using a sliding window and frequency map; Unicode-safe via rune processing.
