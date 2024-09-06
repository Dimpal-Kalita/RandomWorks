package booking

import(
	"sync"
	"testing"
)


func TestConcurrentBooking(t *testing.T){
	b := NewBooking(10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++{
		wg.Add(1)
		go func(seatNumber int){
			defer wg.Done()
			err := b.BookSeat(seatNumber)
			if err != nil{
				t.Error(err)
			}
		}(i+1)
	}
	wg.Wait()
	if b.BookedSeats != 10{
		t.Errorf("Expected 10 booked seats, got %d", b.BookedSeats)
	}
}

func TestConcurrentCancellation(t *testing.T){
	b:= NewBooking(10)
	var wg sync.WaitGroup
	for i:=0; i<10; i++{
		wg.Add(1)
		go func(seatNumber int){
			defer wg.Done()
			err := b.BookSeat(seatNumber)
			if err != nil{
				t.Error(err)
			}
		}(i+1)
	}
	wg.Wait()
	if b.BookedSeats != 10{
		t.Errorf("Expected 10 booked seats, got %d", b.BookedSeats)
	}
	wg.Add(1)
	go func(){
		defer wg.Done()
		err := b.CancelSeat(1)
		if err != nil{
			t.Error(err)
		}
	}()
	wg.Wait()
	if b.BookedSeats != 9{
		t.Errorf("Expected 9 booked seats, got %d", b.BookedSeats)
	}
}