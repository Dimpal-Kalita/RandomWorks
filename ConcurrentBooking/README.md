# Concurrent Booking to handle multiple booking requests at the same time

## How to use the CLI
- To Book seats 1,2,3
 ```bash
    go run main.go -book 1 2 3
 ```
- To cancel seats 1,2,3
 ```bash
    go run main.go -cancel 1 2 3
 ```
- To show the availabel seats
 ```bash
    go run main.go -show
 ```
- The flags can be combined
 ```bash
    go run main.go -book 1 2 3 -cancel 1,2 -show
 ```

## Implementation
```go
type Seat struct{
	SeatNumber int
	IsBooked bool
}

type Booking struct{
	Seats []Seat
	TotalSeats int
	BookedSeats int
	mu sync.Mutex
}

```
- The Seat struct is used to represent a seat in the auditorium. It has two fields, SeatNumber and IsBooked.
- The Booking struct is used to represent the auditorium. It has three fields, Seats, TotalSeats, and BookedSeats and a mutex to handle concurrent requests.

```go

func (b *Booking) BookSeat(seatNumber int) error{
	b.mu.Lock()
	defer b.mu.Unlock()
	if seatNumber<1 || seatNumber > b.TotalSeats{
		return errors.New("Seat number is invalid")
	}

	if b.Seats[seatNumber-1].IsBooked{
		return errors.New(fmt.Sprintf("Seat number %d is already booked", seatNumber))
	}
	b.Seats[seatNumber-1].IsBooked = true
	b.BookedSeats++
	return nil
}
```
- The BookSeat method is used to book a seat. It takes a seat number as an argument and returns an error if the seat number is invalid or the seat is already booked.
- The method uses a mutex to ensure that only one goroutine can book a seat at a time.

```go

func (b *Booking) CancelSeat(seatNumber int) error{
	b.mu.Lock()
	defer b.mu.Unlock()
	if seatNumber<1 || seatNumber > b.TotalSeats{
		return errors.New("Seat number is invalid")
	}
	if !b.Seats[seatNumber-1].IsBooked{
		return errors.New(fmt.Sprintf("Seat number %d is not booked", seatNumber))
	}
	b.Seats[seatNumber-1].IsBooked = false
	b.BookedSeats--
	return nil
}

```
- The CancelSeat method is used to cancel a seat. It takes a seat number as an argument and returns an error if the seat number is invalid or the seat is not booked.
- The method uses a mutex to ensure that only one goroutine can cancel a seat at a time.

```go

func (b *Booking) ShowAvailableSeats() []int{
	b.mu.Lock()
	defer b.mu.Unlock()
	availableSeats := make([]int, 0)
	for i := 0; i < b.TotalSeats; i++{
		if !b.Seats[i].IsBooked{
			availableSeats = append(availableSeats, b.Seats[i].SeatNumber)
		}
	}
	return availableSeats
}
```
- The ShowAvailableSeats method is used to show the available seats. It returns a slice of seat numbers that are not booked.


### CLI
- The CLI is implemented using the flag package in Go.
- The main function parses the command-line arguments and calls the appropriate methods on the Booking struct.

```go
	bookFlag:= flag.String("book","","Book seats (e.g. -book 1,2,3)")
	cancelFlag:= flag.String("cancel","","Cancel seats (e.g. -cancel 1,2,3)")
	showFlag:= flag.Bool("show",false,"Show booked seats (e.g. -show)")
	flag.Parse()
```
- The flag package is used to define three flags, bookFlag, cancelFlag, and showFlag.

```go
	if *bookFlag != ""{
		seats, err := ParseSeatNumbers(*bookFlag)
		if err != nil{
			fmt.Println("Invalid input")
			ProperFormat()
			return
		}
		for _, seat := range seats{
			wg.Add(1)
			go func(seat int){
				defer wg.Done()
				if err:= b.BookSeat(seat); err != nil{
					fmt.Println(err)
				}
			}(seat)
		}
		wg.Wait()
	}
```
- If the bookFlag is set, the ParseSeatNumbers function is called to parse the seat numbers.
- goroutine is used for multiple booking requests at the same time.

```go

func ParseSeatNumbers(input string) ([]int, error){
	seatNumbers := strings.Split(input,",")
	var seats []int
	for _, seatNumber := range seatNumbers{
		seat, err := strconv.Atoi(seatNumber)
		if err != nil{
			return nil, err
		}
		seats = append(seats,seat)
	}
	return seats, nil
}
```
- The ParseSeatNumbers function is used to parse the seat numbers from the input string.
