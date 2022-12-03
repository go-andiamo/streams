# Streams
[![GoDoc](https://godoc.org/github.com/go-andiamo/streams?status.svg)](https://pkg.go.dev/github.com/go-andiamo/streams)
[![Latest Version](https://img.shields.io/github/v/tag/go-andiamo/streams.svg?sort=semver&style=flat&label=version&color=blue)](https://github.com/go-andiamo/streams/releases)
[![codecov](https://codecov.io/gh/go-andiamo/streams/branch/main/graph/badge.svg?token=igjnZdgh0e)](https://codecov.io/gh/go-andiamo/streams)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-andiamo/streams)](https://goreportcard.com/report/github.com/go-andiamo/streams)

A very light Streams implementation in Golang

## Installation
To install Streams, use go get:

    go get github.com/go-andiamo/streams

To update Streams to the latest version, run:

    go get -u github.com/go-andiamo/streams

## Interfaces

<details>
    <summary><strong>Stream Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>AllMatch(p Predicate[T])</code><br>
                <ul>
                    returns whether all elements of this stream match the provided predicate<br>
                    <em>if the provided predicate is nil or the stream is empty, always returns false</em>
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>AnyMatch(p Predicate[T])</code><br>
                <ul>
                    returns whether any elements of this stream match the provided predicate<br>
                    <em>if the provided predicate is nil or the stream is empty, always returns false</em>
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Append(items ...T)</code><br>
                <ul>
                    creates a new stream with all the elements of this stream followed by the specified elements
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Concat(add Stream[T])</code><br>
                <ul>
                    creates a new stream with all the elements of this stream followed by all the elements of the added stream
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Count(p Predicate[T])</code><br>
                <ul>
                    returns the count of elements that match the provided predicate<br>
                    <em>If the predicate is nil, returns the count of all elements</em>
                </ul>
            </td>
            <td>
                <code>int</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Difference(other Stream[T], c Comparator[T])</code><br>
                <ul>
                    creates a new stream that is the set difference between this and the supplied other stream<br>
                    <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)</em>
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Distinct()</code><br>
                <ul>
                    creates a new stream of distinct elements in this stream
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Filter(p Predicate[T])</code><br>
                <ul>
                    creates a new stream of elements in this stream that match the provided predicate<br>
                    <em>if the provided predicate is nil, all elements in this stream are returned</em>
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>FirstMatch(p Predicate[T])</code><br>
                <ul>
                    returns an optional of the first element that matches the provided predicate<<br>
                    if no elements match the provided predicate, an empty (not present) optional is returned<br>
                    <em>if the provided predicate is nil, the first element in this stream is returned</em>
                </ul>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>ForEach(c Consumer[T])</code><br>
                <ul>
                    performs an action on each element of this stream<br>
                    the action to be performed is defined by the provided consumer<br>
                    <em>if the provided consumer is nil, nothing is performed</em>
                </ul>
            </td>
            <td>
                <code>error</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Has(v T, c Comparator[T])</code><br>
                <ul>
                    returns whether this stream contains an element that is equal to the element value provided<br>
                    equality is determined using the provided comparator<br>
                    <em>if the provided comparator is nil, always returns false</em>
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Intersection(other Stream[T], c Comparator[T])</code><br>
                <ul>
                    creates a new stream that is the set intersection of this and the supplied other stream<br>
                    <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)</em>
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Iterator(ps ...Predicate[T])</code><br>
                <ul>
                    returns an iterator (pull) function<br>
                    the iterator function can be used in for loops, for example<br>
                    <pre>
next := strm.Iterator()
for v, ok := next(); ok; v, ok = next() {
    fmt.Println(v)
}</pre>
                    <code>Iterator</code> can also optionally accept varargs of Predicate - which, if specified, are logically OR-ed on each pull to ensure that pulled elements match
                </ul>
            </td>
            <td>
                <code>func() (T, bool)</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>LastMatch(p Predicate[T])</code><br>
                <ul>
                    returns an optional of the last element that matches the provided predicate<br>
                    if no elements match the provided predicate, an empty (not present) optional is returned<br>
                    <em>if the provided predicate is nil, the last element in this stream is returned</em>
                </ul>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Len()</code><br>
                <ul>
                    returns the length (number of elements) of this stream
                </ul>
            </td>
            <td>
                <code>int</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Limit(maxSize int)</code><br>
                <ul>
                    creates a new stream whose number of elements is limited to the value provided<br>
                    if the maximum size is greater than the length of this stream, all elements are returned
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Max(c Comparator[T])</code><br>
                <ul>
                    returns the maximum element of this stream according to the provided comparator<br>
                    <em>if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned</em>
                </ul>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Min(c Comparator[T])</code><br>
                <ul>
                    returns the minimum element of this stream according to the provided comparator<br>
                    <em>if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned</em>
                </ul>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>MinMax(c Comparator[T])</code><br>
                <ul>
                    returns the minimum and maximum element of this stream according to the provided comparator<br>
                    <em>if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned for both</em>
                </ul>
            </td>
            <td>
                <code>(Optional[T], Optional[T])</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>NoneMatch(p Predicate[T])</code><br>
                <ul>
                    returns whether none of the elements of this stream match the provided predicate<br>
                    <em>if the provided predicate is nil or the stream is empty, always returns true</em>
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>NthMatch(p Predicate[T], nth int)</code><br>
                <ul>
                    returns an optional of the nth matching element (1 based) according to the provided predicate<br>
                    if the nth argument is negative, the nth is taken as relative to the last<br>
                    <em>if the provided predicate is nil, any element is taken as matching</em>
                    <em>if no elements match in the specified position, an empty (not present) optional is returned</em>
                </ul>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Reverse()</code><br>
                <ul>
                    creates a new stream composed of elements from this stream but in reverse order
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Skip(n int)</code><br>
                <ul>
                    creates a new stream consisting of this stream after discarding the first <em><strong>n</strong></em> elements<br>
                    if the specified n to skip is equal to or greater than the number of elements in this stream, an empty stream is returned
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Slice(start int, count int)</code><br>
                <ul>
                    creates a new stream composed of elements from this stream starting at the specified start and including the specified count (or to the end)<br>
                    the start is zero based (and less than zero is ignored)<br>
                    if the specified count is negative, items are selected from the start and then backwards by the count
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Sorted(c Comparator[T])</code><br>
                <ul>
                    creates a new stream consisting of the elements of this stream, sorted according to the provided comparator<br>
                    <em>if the provided comparator is nil, the elements are not sorted</em>
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>SymmetricDifference(other Stream[T], c Comparator[T])</code><br>
                <ul>
                    creates a new stream that is the set symmetric difference between this and the supplied other stream<br>
                    <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)</em>        
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Union(other Stream[T], c Comparator[T])</code><br>
                <ul>
                    creates a new stream that is the set union of this and the supplied other stream<br>
                    <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is always empty)</em>
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Unique(c Comparator[T])</code>
                <ul>
                    creates a new stream of unique elements in this stream<br>
                    uniqueness is determined using the provided comparator<br>
                    if provided comparator is nil but the value type of elements in this stream are directly mappable (i.e. primitive or non-pointer types) then
                    <code>Distinct</code> is used as the result, otherwise returns an empty stream
                </ul>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>Of[T any](values ...T) Stream[T]</code><br>
                <ul>
                    creates a new stream of the values provided
                </ul>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>OfSlice[T any](s []T) Stream[T]</code><br>
                <ul>
                    creates a new stream around a slice<br>
                    <em>Note: Once created, If the slice changes the stream does not</em>
                </ul>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewStreamableSlice[T any](sl *[]T) Stream[T]</code><br>
                <ul>
                    creates a Stream from a pointer to a slice<br>
                    <em>It differs from casting a slice to Streamable in that if the underlying slice changes, so does the Stream</em>
                </ul>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <strong>Casting as Streamable</strong><br>
                <pre>
sl := []string{"a", "b", "c"}
s := Streamable[string](sl)</pre>
                <ul>
                    casts a slice as a <code>Stream</code><br>
                    <em>Note: Once cast, if the slice changes the stream does not</em>
                </ul>
            </td>
        </tr>
    </table>
</details>

<details>
    <summary><strong>Comparator Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Compare(v1, v2 T)</code><br>
                <ul>
                    compares the two values lexicographically, i.e.:
                    <ul>
                        <li>the result should be <strong>0</strong> if v1 == v2</li>
                        <li>the result should be <strong>-1</strong> if v1 < v2</li>
                        <li>the result should be <strong>1</strong> if v1 > v2</li>
                    </ul>
                </ul>
            </td>
            <td>
                <code>int</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Greater(v1, v2 T)</code><br>
                <ul>
                    returns true if v1 > v2, otherwise false
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>GreaterOrEqual(v1, v2 T)</code><br>
                <ul>
                    returns true if v1 >= v2, otherwise false
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Equal(v1, v2 T)</code><br>
                <ul>
                    returns true if v1 == v2, otherwise false
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Less(v1, v2 T)</code><br>
                <ul>
                    returns true if v1 < v2, otherwise false
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>LessOrEqual(v1, v2 T)</code><br>
                <ul>
                    returns true if v1 <= v2, otherwise false
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>NotEqual(v1, v2 T)</code><br>
                <ul>
                    returns true if v1 != v2, otherwise false
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Reversed()</code><br>
                <ul>
                    creates a new comparator that imposes the reverse ordering to this comparator<br>
                    the reversal is against less/greater as well as against equality/non-equality
                </ul>
            </td>
            <td>
                <code>Comparator[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Then(other Comparator[T])</code><br>
                <ul>
                    creates a new comparator from this comparator, with a following then comparator
                    that is used when the initial comparison yields equal
                </ul>
            </td>
            <td>
                <code>Comparator[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewComparator[T any](f ComparatorFunc[T]) Comparator[T]</code><br>
                <ul>
                    creates a new <code>Comparator</code> from the function provided<br>
                    where the comparator function is:<br>
                    <code>type ComparatorFunc[T any] func(v1, v2 T) int</code><br>
                    which returns:
                    <ul>
                        <li><strong>0</strong> if v1 == v2</li>
                        <li><strong>-1</strong> if v1 < v2</li>
                        <li><strong>1</strong> if v1 > v2</li>
                    </ul>
                </ul>
            </td>
        </tr>
    </table>
</details>

<details>
    <summary><strong>Consumer Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Accept(v T)</code><br>
                <ul>
                    is called by the user of the consumer to supply a value
                </ul>
            </td>
            <td>
                <code>error</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>AndThen(after Consumer[T])</code><br>
                <ul>
                    creates a new consumer from the current with a subsequent action to be performed<br>
                    <em>multiple consumers can be chained together as one using this method</em>
                </ul>
            </td>
            <td>
                <code>Consumer[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewConsumer[T any](f ConsumerFunc[T]) Consumer[T]</code><br>
                <ul>
                    creates a new <code>Consumer</code> from the function provided<br>
                    where the consumer function is:<br>
                    <code>type ConsumerFunc[T any] func(v T) error</code>    
                </ul>
            </td>
        </tr>
    </table>
</details>

<details>
    <summary><strong>Predicate Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Test(v T)</code><br>
                <ul>
                    evaluates this predicate against the supplied value
                </ul>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>And(other Predicate[T])</code><br>
                <ul>
                    creates a composed predicate that represents a short-circuiting logical AND of this predicate and another
                </ul>
            </td>
            <td>
                <code>Predicate[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Or(other Predicate[T])</code><br>
                <ul>
                    creates a composed predicate that represents a short-circuiting logical OR of this predicate and another
                </ul>
            </td>
            <td>
                <code>Predicate[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Negate()</code><br>
                <ul>
                    creates a composed predicate that represents a logical NOT of this predicate
                </ul>
            </td>
            <td>
                <code>Predicate[T]</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewPredicate[T any](f PredicateFunc[T]) Predicate[T]</code><br>
                <ul>
                    creates a new <code>Predicate</code> from the function provided<br>
                    where the predicate function is:<br>
                    <code>type PredicateFunc[T any] func(v T) bool</code>
                </ul>
            </td>
        </tr>
    </table>
</details>

### Mapper Interfaces
<details>
    <summary><strong>Mapper Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Map(in Stream[T])</code><br>
                <ul>
                    converts the values in the input <code>Stream</code> and produces a <code>Stream</code> of output types
                </ul>
            </td>
            <td>
                <code>(Stream[R], error)</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewMapper[T any, R any](c Converter[T, R]) Mapper[T, R]</code><br>
                <ul>
                    creates a new <code>Mapper</code> that will use the provided <code>Converter</code><br>
                    <em><code>NewMapper</code> panics if a nil <code>Converter</code> is supplied</em>
                </ul>
            </td>
        </tr>        
    </table>
</details>

<details>
    <summary><strong>Converter Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Convert(v T)</code><br>
                <ul>
                    converts a value of type <strong>T</strong> and returns a value of type <strong>R</strong>
                </ul>
            </td>
            <td>
                <code>(R, error)</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewConverter[T any, R any](f ConverterFunc[T, R]) Converter[T, R]</code><br>
                <ul>
                    creates a new <code>Converter</code> from the function provided<br>
                    where the converter function is:<br>
                    <code>type ConverterFunc[T any, R any] func(v T) (R, error)</code>
                </ul>
            </td>
        </tr>        
    </table>
</details>

### Reducer Interfaces
<details>
    <summary><strong>Reducer Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Reduce(s Stream[T])</code><br>
                <ul>
                    performs a reduction of the supplied <code>Stream</code>
                </ul>
            </td>
            <td>
                <code>R</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewReducer[T any, R any](accumulator Accumulator[T, R]) Reducer[T, R]</code><br>
                <ul>
                    creates a new <code>Reducer</code> that will use the supplied <code>Accumulator</code><br>
                    <em><code>NewReducer</code> panics if a nil <code>Accumulator</code> is supplied</em>
                </ul>
            </td>
        </tr>        
    </table>
</details>

<details>
    <summary><strong>Accumulator Interface</strong></summary>
    <table>
        <tr>
            <th>Method and description</th>
            <th>Returns</th>
        </tr>
        <tr></tr>
        <tr>
            <td>
                <code>Apply(t T, r R)</code><br>
                <ul>
                    adds the value of <strong>T</strong> to <strong>R</strong>, and returns the new <strong>R</strong>
                </ul>
            </td>
            <td>
                <code>R</code>
            </td>
        </tr>
        <tr></tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr></tr>
        <tr>
            <td colspan="2">
                <code>NewAccumulator[T any, R any](f AccumulatorFunc[T, R]) Accumulator[T, R]</code><br>
                <ul>
                    creates a new <code>Accumulator</code> from the function provided<br>
                    where the accumulator function is:<br>
                    <code>type AccumulatorFunc[T any, R any] func(t T, r R) R</code>
                </ul>
            </td>
        </tr>        
    </table>
</details>

## Examples
<details>
    <summary><strong>Find first match...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
    "strings"
)

func main() {
    s := Of("a", "B", "c", "D", "e", "F")
    upperPredicate := NewPredicate(func(v string) bool {
        return strings.ToUpper(v) == v
    })
    o := s.FirstMatch(upperPredicate)
    o.IfPresentOtherwise(
        func(v string) {
            println(`Found: "` + v + `"`)
        },
        func() {
            println(`Did not find an uppercase`)
        },
    )
}
```
[try on go-playground](https://go.dev/play/p/C-GYuInfkNm)

</details>

<details>
    <summary><strong>Find last match...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
    "strings"
)

func main() {
    s := Of("a", "B", "c", "D", "e", "F")
    upperPredicate := NewPredicate(func(v string) bool {
        return strings.ToUpper(v) == v
    })
    o := s.LastMatch(upperPredicate.Negate())
    o.IfPresentOtherwise(
        func(v string) {
            println(`Found: "` + v + `"`)
        },
        func() {
            println(`Did not find a lowercase`)
        },
    )
}
```
[try on go-playground](https://go.dev/play/p/2UcwpZEuV-L)

</details>

<details>
    <summary><strong>Nth match...</strong></summary>

```go
package main

import (
    "fmt"
    . "github.com/go-andiamo/streams"
    "strings"
)

func main() {
    s := Of("a", "B", "c", "D", "e", "F", "g", "H", "i", "J")
    upperPredicate := NewPredicate(func(v string) bool {
        return strings.ToUpper(v) == v
    })

    for nth := -6; nth < 7; nth++ {
        s.NthMatch(upperPredicate, nth).IfPresentOtherwise(
            func(v string) {
                fmt.Printf("Found \"%s\" at nth pos %d\n", v, nth)
            },
            func() {
                fmt.Printf("No match at nth pos %d\n", nth)
            },
        )
    }
}
```
[try on go-playground](https://go.dev/play/p/I_DrKK4ZXAT)

</details>

<details>
    <summary><strong>Sort descending & print...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
)

func main() {
    s := Of("311", "AAAA", "30", "3", "1", "Baaa", "4000", "0400", "40", "Aaaa", "BBBB", "4", "01", "2", "0101", "201", "20")
    _ = s.Sorted(StringComparator.Reversed()).ForEach(NewConsumer(func(v string) error {
        println(v)
        return nil
    }))
}
```
[try on go-playground](https://go.dev/play/p/bU6UZ479pF1)

</details>

<details>
    <summary><strong>Compound sort...</strong></summary>

```go
package main

import (
    "fmt"
    . "github.com/go-andiamo/streams"
)

func main() {
    type myStruct struct {
        value    int
        priority int
    }
    myComparator := NewComparator(func(v1, v2 myStruct) int {
        return IntComparator.Compare(v1.value, v2.value)
    }).Then(NewComparator(func(v1, v2 myStruct) int {
        return IntComparator.Compare(v1.priority, v2.priority)
    }).Reversed())
    s := Of(
        myStruct{
            value:    2,
            priority: 2,
        },
        myStruct{
            value:    2,
            priority: 0,
        },
        myStruct{
            value:    2,
            priority: 1,
        },
        myStruct{
            value:    1,
            priority: 2,
        },
        myStruct{
            value:    1,
            priority: 1,
        },
        myStruct{
            value:    1,
            priority: 0,
        },
    )
    _ = s.Sorted(myComparator).ForEach(NewConsumer(func(v myStruct) error {
        fmt.Printf("Value: %d, Priority: %d\n", v.value, v.priority)
        return nil
    }))
}
```
[try on go-playground](https://go.dev/play/p/TYeVcgjdAB3)

</details>

<details>
    <summary><strong>Min and max...</strong></summary>

```go
package main

import (
    "fmt"
    . "github.com/go-andiamo/streams"
)

func main() {
    type myStruct struct {
        value    int
        priority int
    }
    myComparator := NewComparator(func(v1, v2 myStruct) int {
        return IntComparator.Compare(v1.value, v2.value)
    }).Then(NewComparator(func(v1, v2 myStruct) int {
        return IntComparator.Compare(v1.priority, v2.priority)
    }))
    s := Of(
        myStruct{
            value:    2,
            priority: 2,
        },
        myStruct{
            value:    2,
            priority: 0,
        },
        myStruct{
            value:    2,
            priority: 1,
        },
        myStruct{
            value:    1,
            priority: 2,
        },
        myStruct{
            value:    1,
            priority: 1,
        },
        myStruct{
            value:    1,
            priority: 0,
        },
    )
    min := s.Min(myComparator)
    min.IfPresentOtherwise(
        func(v myStruct) {
            fmt.Printf("Min... Value: %d, Priority: %d\n", v.value, v.priority)
        },
        func() {
            println("No min found!")
        })
    max := s.Max(myComparator)
    max.IfPresentOtherwise(
        func(v myStruct) {
            fmt.Printf("Max... Value: %d, Priority: %d\n", v.value, v.priority)
        },
        func() {
            println("No max found!")
        })
}
```
[try on go-playground](https://go.dev/play/p/7uUKh5Qg-7L)

</details>

<details>
    <summary><strong>Set intersection, union, difference and symmetric difference...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
)

func main() {
    s1 := Of("a", "B", "c", "C", "d", "D", "d")
    s2 := Of("e", "E", "a", "A", "b")
    println("Intersection...")
    _ = s1.Unique(StringInsensitiveComparator).Intersection(s2.Unique(StringInsensitiveComparator), StringInsensitiveComparator).
        ForEach(NewConsumer(func(v string) error {
            println(v)
            return nil
        }))
    println("Union...")
    _ = s1.Unique(StringInsensitiveComparator).Union(s2.Unique(StringInsensitiveComparator), StringInsensitiveComparator).
        ForEach(NewConsumer(func(v string) error {
            println(v)
            return nil
        }))
    println("Symmetric Difference...")
    _ = s1.Unique(StringInsensitiveComparator).SymmetricDifference(s2.Unique(StringInsensitiveComparator), StringInsensitiveComparator).
        ForEach(NewConsumer(func(v string) error {
            println(v)
            return nil
        }))
    println("Difference (s1 to s2)...")
    _ = s1.Unique(StringInsensitiveComparator).Difference(s2.Unique(StringInsensitiveComparator), StringInsensitiveComparator).
        ForEach(NewConsumer(func(v string) error {
            println(v)
            return nil
        }))
    println("Difference (s2 to s1)...")
    _ = s2.Unique(StringInsensitiveComparator).Difference(s1.Unique(StringInsensitiveComparator), StringInsensitiveComparator).
        ForEach(NewConsumer(func(v string) error {
            println(v)
            return nil
        }))
}
```
[try on go-playground](https://go.dev/play/p/IMpym38xHXV)

</details>

<details>
    <summary><strong>Map...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
)

func main() {
    type character struct {
        name string
        age  int
    }
    characters := OfSlice([]character{
        {
            `Frodo Baggins`,
            50,
        },
        {
            `Samwise Gamgee`,
            38,
        },
        {
            `Gandalf`,
            2000,
        },
        {
            `Aragorn`,
            87,
        },
        {
            `Legolas`,
            200,
        },
        {
            `Gimli`,
            139,
        },
        {
            `Meridoc Brandybuck`,
            36,
        },
        {
            `Peregrin Took`,
            28,
        },
        {
            `Boromir`,
            40,
        },
    })

    m := NewMapper(NewConverter[character, string](func(v character) (string, error) {
        return v.name, nil
    }))
    names, _ := m.Map(characters)
    _ = names.Sorted(StringComparator).ForEach(NewConsumer(func(v string) error {
        println(v)
        return nil
    }))
}
```
[try on go-playground](https://go.dev/play/p/oe5ZTSbzUyy)

</details>

<details>
    <summary><strong>Reduce...</strong></summary>

```go
package main

import (
    "fmt"
    . "github.com/go-andiamo/streams"
)

func main() {
    type account struct {
        currency string
        acNo     string
        balance  float64
    }
    accounts := OfSlice([]account{
        {
            `GBP`,
            `1051065`,
            50.01,
        },
        {
            `USD`,
            `1931132`,
            259.98,
        },
        {
            `EUR`,
            `1567807`,
            313.25,
        },
        {
            `EUR`,
            `1009321`,
            50.01,
        },
        {
            `USD`,
            `1573756`,
            12.02,
        },
        {
            `GBP`,
            `1456044`,
            99.99,
        },
    })

    accum := NewAccumulator[account, map[string]float64](func(v account, r map[string]float64) map[string]float64 {
        if r == nil {
            r = map[string]float64{}
        }
        if cv, ok := r[v.currency]; ok {
            r[v.currency] = cv + v.balance
        } else {
            r[v.currency] = v.balance
        }
        return r
    })
    r := NewReducer(accum)
    for k, v := range r.Reduce(accounts) {
        fmt.Printf("%s %f\n", k, v)
    }
}
```
[try on go-playground](https://go.dev/play/p/HwgwlkNFUTQ)

</details>

<details>
    <summary><strong>Filter with composed predicate...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
    "regexp"
    "strings"
)

func main() {
    s := Of("aaa", "", "AAA", "012", "bBbB", "Ccc", "CCC", "D", "EeE", "eee", " ", "  ", "A12")

    pNotEmpty := NewPredicate(func(v string) bool {
        return len(strings.Trim(v, " ")) > 0
    })
    rxNum := regexp.MustCompile(`^[0-9]+$`)
    pNumeric := NewPredicate(func(v string) bool {
        return rxNum.MatchString(v)
    })
    rxUpper := regexp.MustCompile(`^[A-Z]+$`)
    pAllUpper := NewPredicate(func(v string) bool {
        return rxUpper.MatchString(v)
    })
    rxLower := regexp.MustCompile(`^[a-z]+$`)
    pAllLower := NewPredicate(func(v string) bool {
        return rxLower.MatchString(v)
    })
    // only want strings that are non-empty and all numeric, all upper or all lower... 
    pFinal := pNotEmpty.And(pNumeric.Or(pAllUpper).Or(pAllLower))

    _ = s.Filter(pFinal).ForEach(NewConsumer(func(v string) error {
        println(v)
        return nil
    }))
}
```
[try on go-playground](https://go.dev/play/p/WxOOpEv-kI0)

</details>

<details>
    <summary><strong>Distinct vs Unique...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
)

func main() {
    s := Of("a", "A", "b", "B", "c", "C")

    println("Distinct...")
    _ = s.Distinct().ForEach(NewConsumer(func(v string) error {
        println(v)
        return nil
    }))
    println("Unique (case insensitive)...")
    _ = s.Unique(StringInsensitiveComparator).ForEach(NewConsumer(func(v string) error {
        println(v)
        return nil
    }))
}
```
[try on go-playground](https://go.dev/play/p/JZY9b6o6OLd)

</details>

<details>
    <summary><strong>Distinct vs Unique (structs)...</strong></summary>

```go
package main

import (
    "fmt"
    . "github.com/go-andiamo/streams"
    "strings"
)

type MyStruct struct {
    value    string
    priority int
}

func main() {
    s1 := OfSlice([]MyStruct{
        {
            "A",
            1,
        },
        {
            "A",
            2,
        },
        {
            "A",
            1,
        },
        {
            "A",
            2,
        },
    })

    println("\nStruct Distinct...")
    _ = s1.Distinct().ForEach(NewConsumer(func(v MyStruct) error {
        fmt.Printf("Value: %s, Priority: %d\n", v.value, v.priority)
        return nil
    }))
    println("\nStruct Unique (no comparator)...")
    _ = s1.Unique(nil).ForEach(NewConsumer(func(v MyStruct) error {
        fmt.Printf("Value: %s, Priority: %d\n", v.value, v.priority)
        return nil
    }))

    s2 := OfSlice([]*MyStruct{
        {
            "A",
            1,
        },
        {
            "A",
            2,
        },
        {
            "A",
            1,
        },
        {
            "A",
            2,
        },
    })

    println("\nStruct Ptr Distinct...")
    _ = s2.Distinct().ForEach(NewConsumer(func(v *MyStruct) error {
        fmt.Printf("Value: %s, Priority: %d\n", v.value, v.priority)
        return nil
    }))
    println("\nStruct Ptr Unique (no comparator)...")
    _ = s2.Unique(nil).ForEach(NewConsumer(func(v *MyStruct) error {
        fmt.Printf("Value: %s, Priority: %d\n", v.value, v.priority)
        return nil
    }))
    cmp := NewComparator(func(v1, v2 *MyStruct) int {
        return strings.Compare(v1.value, v2.value)
    }).Then(NewComparator(func(v1, v2 *MyStruct) int {
        return IntComparator.Compare(v1.priority, v2.priority)
    }))
    println("\nStruct Ptr Unique (with comparator)...")
    _ = s2.Unique(cmp).ForEach(NewConsumer(func(v *MyStruct) error {
        fmt.Printf("Value: %s, Priority: %d\n", v.value, v.priority)
        return nil
    }))
}
```
[try on go-playground](https://go.dev/play/p/9Caw_RT4hwp)

</details>

<details>
    <summary><strong>For each...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
)

var stringValuePrinter = NewConsumer(func(v string) error {
    println(v)
    return nil
})

func main() {
    s := Of("a", "b", "c", "d")

    _ = s.ForEach(stringValuePrinter)
}
```
[try on go-playground](https://go.dev/play/p/yks2s2E8czr)

</details>

<details>
    <summary><strong>Iterator...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
)

func main() {
    s := Of("a", "b", "c", "d")

    next := s.Iterator()
    for v, ok := next(); ok; v, ok = next() {
        println(v)
    }
}
```
[try on go-playground](https://go.dev/play/p/Yae6ZLi2vVj)

</details>

<details>
    <summary><strong>Iterator with predicate...</strong></summary>

```go
package main

import (
    . "github.com/go-andiamo/streams"
    "strings"
)

func main() {
    s := Of("a", "B", "c", "D", "e", "F", "g", "H", "i", "J")
    upper := NewPredicate(func(v string) bool {
        return strings.ToUpper(v) == v
    })

    next := s.Iterator(upper)
    for v, ok := next(); ok; v, ok = next() {
        println(v)
    }
}
```
[try on go-playground](https://go.dev/play/p/GDQJDsZsSY9)

</details>
