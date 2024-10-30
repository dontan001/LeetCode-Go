package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	a      = 1
	b      = 2
	c      = 0b0001
	d      = 0b_0001_1100
	gopher = `         
           ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
)

func main() {
	defer Trace()

	/*var n int
	fmt.Scanf("%d\n", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	quickSort(nums, 0, n-1)

	fmt.Printf("there are %d numbers. \n", len(nums))
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}

	fmt.Printf("\nrevert order \n")
	for idx, _ := range nums {
		fmt.Printf("%d ", nums[len(nums)-idx-1])
	}*/

	/*var container = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3}
	k := "four"
	v, ok := container[k]
	if ok {
		fmt.Printf("%s", v)
	} else {
		fmt.Println("not found")
	}*/

	/*ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	element := <-ch
	fmt.Println(element)*/

	/*var p = People{
		name:   "cherry",
		gender: "female",
	}
	fmt.Println(p)

	var c = Chinese{}
	c.People.String()*/

	/*totalTask := 10
	check := make(chan int, totalTask)
	for i := 0; i < totalTask; i++ {
		go func() {
			fmt.Printf("process # %d \n", i)
			check <- i
		}()
	}

	for i := 0; i < totalTask; i++ {
		fmt.Printf("done # %d \n", <-check)
	}

	fmt.Printf("the main goroutine is done\n")*/

	/*numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, i := range numbers {
		fmt.Println(idx, i)
	}
	coordinateWithChan()*/

	/*fmt.Printf("%b \n", 7)*/

	produceAndConsume()
}

type Human interface {
	SetName(name string)
	SetGender(gender string)
}

type People struct {
	name    string
	gender  string
	address string
}

func (p People) String() string {
	return p.name + " " + p.gender + " " + p.address
}

func (p People) SetName(name string) {
	p.name = name
}

func (p People) SetGender(gender string) {
	p.gender = gender
}

type Chinese struct {
	Province int
	People
}

func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}

func GetToken() {
	mu := sync.Mutex{}
	defer func() { mu.Unlock() }()

	mu.Lock()
	fmt.Printf("you have a token @ %s\n", time.Now().String())
}

func coordinateWithChan() {
	defer Trace()

	end := make(chan struct{}, 2)
	go func() {
		fmt.Printf("first done \n")
		end <- struct{}{}
	}()
	go func() {
		fmt.Printf("second done \n")
		end <- struct{}{}
	}()
	<-end
	<-end
	fmt.Println("all done")
}

func coordinateWithWaitGroup() {
	finish := sync.WaitGroup{}
	finish.Add(2)
	go func() {
		fmt.Printf("first done \n")
		finish.Done()
	}()
	go func() {
		fmt.Printf("second done \n")
		finish.Done()
	}()
	finish.Wait()
	fmt.Println("all done")
}

func init() {
	fmt.Println("main: first init func invoked")
}

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := curGoroutineID()
	fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
	return func() { fmt.Printf("g[%05d]: exit: [%s]\n", gid, name) }
}

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}
