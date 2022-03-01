# go-algorithms

[![GoDoc](https://godoc.org/github.com/irenicaa/go-algorithms?status.svg)](https://godoc.org/github.com/irenicaa/go-algorithms)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-algorithms)](https://goreportcard.com/report/github.com/irenicaa/go-algorithms)
[![Build Status](https://app.travis-ci.com/irenicaa/go-algorithms.svg?branch=master)](https://app.travis-ci.com/irenicaa/go-algorithms)
[![codecov](https://codecov.io/gh/irenicaa/go-algorithms/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-algorithms)

The library implementing some basic algorithms and data structures.

- algorithms:
  - search:
    - search for the minimal item in a slice
    - search for the maximal item in a slice
    - [linear search](https://en.wikipedia.org/wiki/Linear_search)
    - [binary search](https://en.wikipedia.org/wiki/Binary_search_algorithm)
  - filtering out the non-unique items in a slice:
    - universal
    - in a sorted slice only
    - via a set (a data structure); i.e., in a slice containing hashable items only
  - merging two sorted slices
  - sorting:
    - [bubble sort](https://en.wikipedia.org/wiki/Bubble_sort)
    - [cocktail sort](https://en.wikipedia.org/wiki/Cocktail_shaker_sort)
    - [comb sort](https://en.wikipedia.org/wiki/Comb_sort)
    - [insertion sort](https://en.wikipedia.org/wiki/Insertion_sort)
    - [Shell sort](https://en.wikipedia.org/wiki/Shellsort)
    - [selection sort](https://en.wikipedia.org/wiki/Selection_sort)
    - [merge sort](https://en.wikipedia.org/wiki/Merge_sort)
    - [quicksort](https://en.wikipedia.org/wiki/Quicksort)
- data structures:
  - [set](<https://en.wikipedia.org/wiki/Set_(abstract_data_type)>):
    - operations:
      - checking that a set contains an item
      - adding an item
      - removing an item
      - evaluating union with another set
      - evaluating intersection with another set
      - evaluating difference between two sets
  - [multiset](<https://en.wikipedia.org/wiki/Set_(abstract_data_type)#Multiset>):
    - operations:
      - checking that a set contains an item
      - adding an item
      - removing an item
      - evaluating sum with another set
      - evaluating union with another set
      - evaluating intersection with another set
      - evaluating difference between two sets

## Installation

```
$ go get github.com/irenicaa/go-algorithms
```

## License

The MIT License (MIT)

Copyright &copy; 2022 irenica
