# golang_sort_algorithms

This repository is for implementation for learning sort algorithm

## Bubble Sort

Given a range k = 0 to Array length -1

each time from range k to Array length -1

swap the larger one to the back

## Quick Sort

Procedure (A, p, r):

if p < r :
  Select a pivot element from  A[r]

  And each time seperate array into 2 group

  L: each element small than pivot value

  R: each element large than pivot value

  swap the R smallest index element with pivot

  return the index of smallest index of R(q)
  
  Procedure(A,p, q-1)  
  Procedure(A,q+1, A.length-1)


## Merge Sort
Separete Array A into 2 Array
Then Merge with order
MergeSort(A, p, r):
if p < r:
  q = Math.floor((p+r)/2)
  MergeSort(A, p , q)
  MergeSort(A, q+1, r)
  Merge(A, p, q,r)

## Insertion Sort
for each run find the insertion place for element selected

then swap selected place with original one