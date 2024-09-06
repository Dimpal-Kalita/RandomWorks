package booking

import(
	"errors"
	"fmt"
	"sync"
)

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

func NewBooking(totalSeats int) *Booking{
	seats := make([]Seat, totalSeats)
	for i := 0; i < totalSeats; i++{
		seats[i] = Seat{SeatNumber: i+1, IsBooked: false}
	}
	return &Booking{Seats: seats, TotalSeats: totalSeats, BookedSeats: 0}
}

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

