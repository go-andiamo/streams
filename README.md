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
        <tr>
            <td>
                <code>AllMatch(p Predicate[T])</code><br>
                returns whether all elements of this stream match the provided predicate<br>
                <em>if the provided predicate is nil or the stream is empty, always returns false</em>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>AnyMatch(p Predicate[T])</code><br>
                returns whether any elements of this stream match the provided predicate<br>
                <em>if the provided predicate is nil or the stream is empty, always returns false</em>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Append(items ...T)</code><br>
                creates a new stream with all the elements of this stream followed by the specified elements
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Concat(add Stream[T])</code><br>
                creates a new stream with all the elements of this stream followed by all the elements of the added stream
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Count(p Predicate[T])</code><br>
                returns the count of elements that match the provided predicate<br>
                <em>If the predicate is nil, returns the count of all elements</em>
            </td>
            <td>
                <code>int</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Difference(other Stream[T], c Comparator[T])</code><br>
                creates a new stream that is the set difference between this and the supplied other stream<br>
                <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is unpredictable)</em>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Distinct()</code><br>
                creates a new stream of distinct elements in this stream
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Filter(p Predicate[T])</code><br>
                creates a new stream of elements in this stream that match the provided predicate<br>
                <em>if the provided predicate is nil, all elements in this stream are returned</em>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>FirstMatch(p Predicate[T])</code><br>
                returns an optional of the first element that matches the provided predicate<<br>
                if no elements match the provided predicate, an empty (not present) optional is returned<br>
                <em>if the provided predicate is nil, the first element in this stream is returned</em>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>ForEach(c Consumer[T])</code><br>
                performs an action on each element of this stream<br>
                the action to be performed is defined by the provided consumer<br>
                <em>if the provided consumer is nil, nothing is performed</em>
            </td>
            <td>
                <code>error</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Has(v T, c Comparator[T])</code><br>
                returns whether this stream contains an element that is equal to the element value provided<br>
                equality is determined using the provided comparator<br>
                <em>if the provided comparator is nil, always returns false</em>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Intersection(other Stream[T], c Comparator[T])</code><br>
                creates a new stream that is the set intersection of this and the supplied other stream<br>
                <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is unpredictable)</em>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>LastMatch(p Predicate[T])</code><br>
                returns an optional of the last element that matches the provided predicate<br>
                if no elements match the provided predicate, an empty (not present) optional is returned<br>
                <em>if the provided predicate is nil, the last element in this stream is returned</em>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Len()</code><br>
                returns the length (number of elements) of this stream
            </td>
            <td>
                <code>int</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Limit(maxSize int)</code><br>
                creates a new stream whose number of elements is limited to the value provided<br>
                if the maximum size is greater than the length of this stream, all elements are returned
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Max(c Comparator[T])</code><br>
                returns the maximum element of this stream according to the provided comparator<br>
                <em>if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned</em>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Min(c Comparator[T])</code><br>
                returns the minimum element of this stream according to the provided comparator<br>
                <em>if the provided comparator is nil or the stream is empty, an empty (not present) optional is returned</em>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>NoneMatch(p Predicate[T])</code><br>
                returns whether none of the elements of this stream match the provided predicate<br>
                <em>if the provided predicate is nil or the stream is empty, always returns true</em>
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>NthMatch(p Predicate[T], nth int)</code><br>
                returns an optional of the nth matching element (1 based) according to the provided predicate<br>
                if the nth argument is negative, the nth is taken as relative to the last<br>
                <em>if the provided predicate is nil, any element is taken as matching</em>
                <em>if no elements match in the specified position, an empty (not present) optional is returned</em>
            </td>
            <td>
                <code>Optional[T]</code>
            </td>
        <tr>
            <td>
                <code>Reverse()</code><br>
                creates a new stream composed of elements from this stream but in reverse order
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Skip(n int)</code><br>
                creates a new stream consisting of this stream after discarding the first <em><strong>n</strong></em> elements<br>
                if the specified n to skip is equal to or greater than the number of elements in this stream, an empty stream is returned
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Slice(start int, count int)</code><br>
                creates a new stream composed of elements from this stream starting at the specified start and including the specified count (or to the end)<br>
                the start is zero based (and less than zero is ignored)<br>
                if the specified count is negative, items are selected from the start and then backwards by the count
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Sorted(c Comparator[T])</code><br>
                creates a new stream consisting of the elements of this stream, sorted according to the provided comparator<br>
                <em>if the provided comparator is nil, the elements are not sorted</em>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>SymmetricDifference(other Stream[T], c Comparator[T])</code><br>
                creates a new stream that is the set symmetric difference between this and the supplied other stream<br>
                <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is unpredictable)</em>        
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Union(other Stream[T], c Comparator[T])</code><br>
                creates a new stream that is the set union of this and the supplied other stream<br>
                <em>equality of elements is determined using the provided comparator (if the provided comparator is nil, the result is unpredictable)</em>
            </td>
            <td>
                <code>Stream[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Unique(c Comparator[T])</code>
                creates a new stream of unique elements in this stream<br>
                uniqueness is determined using the provided comparator<br>
                if provided comparator is nil but the value type of elements in this stream are directly mappable (i.e. primitive or non-pointer types) then
                <code>Distinct</code> is used as the result, otherwise returns an empty stream
            </td>
            <td>
                <code>Stream[T]</code>
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
        <tr>
            <td>
                <code>Compare(v1, v2 T)</code><br>
                compares the two values lexicographically, i.e.:
                <ul>
                    <li>the result should be 0 if v1 == v2</li>
                    <li>the result should be -1 if v1 < v2</li>
                    <li>the result should be 1 if v1 > v2</li>
                </ul>
            </td>
            <td>
                <code>int</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Greater(v1, v2 T)</code><br>
                returns true if v1 > v2, otherwise false
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>GreaterOrEqual(v1, v2 T)</code><br>
                returns true if v1 >= v2, otherwise false
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Equal(v1, v2 T)</code><br>
                returns true if v1 == v2, otherwise false
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Less(v1, v2 T)</code><br>
                returns true if v1 < v2, otherwise false
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>LessOrEqual(v1, v2 T)</code><br>
                returns true if v1 <= v2, otherwise false
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>NotEqual(v1, v2 T)</code><br>
                returns true if v1 != v2, otherwise false
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Reversed()</code><br>
                creates a new comparator that imposes the reverse ordering to this comparator<br>
                the reversal is against less/greater as well as against equality/non-equality
            </td>
            <td>
                <code>Comparator[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Then(other Comparator[T])</code><br>
                creates a new comparator from this comparator, with a following then comparator
                that is used when the initial comparison yields equal
            </td>
            <td>
                <code>Comparator[T]</code>
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
        <tr>
            <td>
                <code>Accept(v T)</code><br>
                is called by the user of the consumer to supply a value
            </td>
            <td>
                <code>error</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>AndThen(after Consumer[T])</code><br>
                creates a new consumer from the current with a subsequent action to be performed<br>
                <em>multiple consumers can be chained together as one using this method</em>
            </td>
            <td>
                <code>Consumer[T]</code>
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
        <tr>
            <td>
                <code>Test(v T)</code><br>
                evaluates this predicate against the supplied value
            </td>
            <td>
                <code>bool</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>And(other Predicate[T])</code><br>
                creates a composed predicate that represents a short-circuiting logical AND of this predicate and another
            </td>
            <td>
                <code>Predicate[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Or(other Predicate[T])</code><br>
                creates a composed predicate that represents a short-circuiting logical OR of this predicate and another
            </td>
            <td>
                <code>Predicate[T]</code>
            </td>
        </tr>
        <tr>
            <td>
                <code>Negate()</code><br>
                creates a composed predicate that represents a logical NOT of this predicate
            </td>
            <td>
                <code>Predicate[T]</code>
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
        <tr>
            <td>
                <code>Map(in Stream[T])</code><br>
                converts the values in the input <code>Stream</code> and produces a <code>Stream</code> of output types
            </td>
            <td>
                <code>(Stream[R], error)</code>
            </td>
        </tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr>
            <td colspan="2">
                <code>NewMapper[T any, R any](c Converter[T, R]) Mapper[T, R]</code><br>
                creates a new <code>Mapper</code> that will use the provided <code>Converter</code>
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
        <tr>
            <td>
                <code>Convert(v T)</code><br>
                converts a value of type <strong>T</strong> and returns a value of type <strong>R</strong>
            </td>
            <td>
                <code>(R, error)</code>
            </td>
        </tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr>
            <td colspan="2">
                <code>NewConverter[T any, R any](f ConverterFunc[T, R]) Converter[T, R]</code><br>
                creates a new <code>Converter</code> from the function provided
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
        <tr>
            <td>
                <code>Reduce(s Stream[T])</code><br>
                performs a reduction of the supplied <code>Stream</code>
            </td>
            <td>
                <code>R</code>
            </td>
        </tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr>
            <td colspan="2">
                <code>NewReducer[T any, R any](accumulator Accumulator[T, R]) Reducer[T, R]</code><br>
                creates a new <code>Reducer</code> that will use the supplied <code>Accumulator</code>
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
        <tr>
            <td>
                <code>Apply(t T, r R)</code><br>
                adds the value of <strong>T</strong> to <strong>R</strong>, and returns the new <strong>R</strong>
            </td>
            <td>
                <code>R</code>
            </td>
        </tr>
        <tr>
            <th colspan="2">Constructors</th>
        </tr>
        <tr>
            <td colspan="2">
                <code>NewAccumulator[T any, R any](f AccumulatorFunc[T, R]) Accumulator[T, R]</code><br>
                creates a new <code>Accumulator</code> from the function provided
            </td>
        </tr>        
    </table>
</details>

## Examples
#### Find first match...
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

#### Find last match...
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

#### Sort descending & print...
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

#### Compound sort...
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

#### Set intersection, union, difference and symmetric difference...
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