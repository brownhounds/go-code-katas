# Go Code Katas

Small Go code katas organized as stand-alone modules. This repo is a personal
practice space for algorithm drills, data-structure patterns, and core Go
syntax fluency.

## Why code katas

- Build speed and accuracy on common patterns.
- Isolate a single concept to practice it deeply with tests.
- Keep muscle memory fresh for interviews or real-world problem solving.
- Create a growing library of reference solutions you can revisit later.

## Modules / Exercises

- `two_sum`: Find two indices whose values add to a target using a hash map
  (complement lookup).
- `rotate_slice_right`: Rotate a slice to the right by `k` steps in-place using
  the three-reversal trick.
- `remove_duplicates`: Remove duplicates from a sorted slice in-place using a
  two-pointer pass; returns the new logical length.
- `move_zeros`: Move all zeros to the end of a slice in-place while preserving
  the order of non-zero elements (stable compaction).
- `maximum_subarray_sum`: Compute the max sum of any contiguous subarray using
  Kadane's algorithm.

## Running tests

Each module is its own Go module. Run tests from the module directory, e.g.:

```bash
cd two_sum
go test
```
