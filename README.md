# pipe

Concurrent, sequential, transformations along Golang channels.

## Usage

```
import "github.com/paulbellamy/pipe"
```

## Interface

Pipes are created with the ```NewPipe(input, output chan interface{}) *Pipe``` method.

After that there are several chaining methods to build up the processing. Once the pipe is prepared, simply pipe items into the input channel and retrieve the results from the output channel.

Be careful, because some of the transformations (e.g. Reduce, Skip) result in channels which are 'leaky'. Meaning that one item in may not equal one item out.

For example, to count the number of items passing through a channel:

```Go
// Define our counter
counter := 0
counter_func := func(item interface{}) {
  counter++
}

// Set up our pipe
input := make(chan interface{}, 5)

// Add our counter
output := ForEach(input, counter_func)

// Now we send some items
input <- true
input <- true
input <- true

// Check how many have gone through
fmt.Println(counter) // prints "3"
```

  You can, of course, modify the items flowing through the pipe:

```Go
// Set up our pipe
input := make(chan interface{}, 5)
output := NewPipe(input).
  Filter(func(item interface{}) bool {
    // Only allow ints divisible by 5
    return (item.(int) % 5) == 0
  }).
  Map(func(item interface{}) interface{} {
    // Add 2 to each
    return item.(int) + 2
  }).
  Output

// Now we send some items
input <- 1 // will be dropped
input <- 5 // will come through as 7
```

## Available Transformations

* Filter(func(item interface{}) bool)
* ForEach(func(item interface{}))
* Map(func(item interface{}) interface{})
* Reduce(initial interface{}, func(accumulator interface{}, item interface{}) interface{})
* Skip(n int64)
* SkipWhile(func(item interface{}) bool)
* Take(n int64)
* TakeWhile(func(item interface{}) bool)
* Zip(other chan interface{})

## Godoc

```
func Filter(input chan interface{}, fn FilterFunc) chan interface{}
    Apply a filtering function to a channel, which will only pass through
    items when the filter func returns true.

func ForEach(input chan interface{}, fn ForEachFunc) chan interface{}
    Execute a function for each item (without modifying the item). Useful
    for monitoring, logging, or causing some side-effect.

func Map(input chan interface{}, fn MapFunc) chan interface{}
    Pass through the result of the map function for each item

func Reduce(input chan interface{}, initial interface{}, fn ReduceFunc) chan interface{}
    Accumulate the result of the reduce function being called on each item,
    then when the input channel is closed, pass the result to the output
    channel

func Skip(input chan interface{}, num int64) chan interface{}
    Skip a given number of items from the input pipe. After that number has
    been dropped, the rest are passed straight through.

func SkipWhile(input chan interface{}, fn SkipWhileFunc) chan interface{}
    Skip the items from the input pipe until the given function returns
    true. After that , the rest are passed straight through.

func Take(input chan interface{}, num int64) chan interface{}
    Accept only the given number of items from the input pipe. After that
    number has been received, all input messages will be ignored and the
    output channel will be closed.

func TakeWhile(input chan interface{}, fn TakeWhileFunc) chan interface{}
    Accept items from the input pipe until the given function returns false.
    After that, all input messages will be ignored and the output channel
    will be closed.

func Zip(input chan interface{}, other chan interface{}) chan interface{}
    Group each message from the input channel with it's corresponding
    message from the other channel. This will block on the first channel
    until it receives a message, then block on the second until it gets one
    from there. At that point an array containing both will be sent to the
    output channel.

    For example, if channel a is being zipped with channel b, and output on
    channel c:

	a <- 1
	b <- 2
	result := <-c // result will equal []interface{}{1, 2}


TYPES

type FilterFunc func(item interface{}) bool

type ForEachFunc func(item interface{})
    A function which foreachs

type MapFunc func(item interface{}) interface{}

type Pipe struct {
    Output chan interface{}
}
    A Pipe is a set of transforms being applied along the channel. We use
    this as a helper while constructing a chained pipe. It lets us use a
    nicer syntax.

func NewPipe(input chan interface{}) *Pipe
    Return a new Pipe object which echoes input to output

func (p *Pipe) Filter(fn FilterFunc) *Pipe
    Helper for chained construction

func (p *Pipe) ForEach(fn ForEachFunc) *Pipe
    Execute a function for each item (without modifying the item)

func (p *Pipe) Map(fn MapFunc) *Pipe
    Helper for chained construction

func (p *Pipe) Reduce(initial interface{}, fn ReduceFunc) *Pipe
    Accumulate the result of the reduce function being called on each item,
    then when the input channel is closed, pass the result to the output
    channel

func (p *Pipe) Skip(num int64) *Pipe
    Helper for chained constructor

func (p *Pipe) SkipWhile(fn SkipWhileFunc) *Pipe
    Helper function for chained constructor

func (p *Pipe) Take(num int64) *Pipe
    Helper for the chained constructor

func (p *Pipe) TakeWhile(fn TakeWhileFunc) *Pipe
    Helper for the chained constructor

func (p *Pipe) Zip(other chan interface{}) *Pipe
    Helper for the chained constructor

type ReduceFunc func(result, item interface{}) interface{}

type SkipWhileFunc func(item interface{}) bool

type TakeWhileFunc func(item interface{}) bool
```
