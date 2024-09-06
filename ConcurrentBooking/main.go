package main

import(
	"flag"
	"fmt"
	"strings"
	"sync"
	"strconv"
	booking "github.com/Dimpal-Kalita/RandomWorks/ConcurrentBooking/utils"
)


func main(){
	bookFlag:= flag.String("book","","Book seats (e.g. -book 1,2,3)")
	cancelFlag:= flag.String("cancel","","Cancel seats (e.g. -cancel 1,2,3)")
	showFlag:= flag.Bool("show",false,"Show booked seats (e.g. -show)")
	flag.Parse()
	b := booking.NewBooking(10)
	var wg sync.WaitGroup
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
	if *cancelFlag != ""{
		seats, err := ParseSeatNumbers(*cancelFlag)
		if err != nil{
			fmt.Println("Invalid input")
			ProperFormat()
			return
		}
		for _, seat := range seats{
			wg.Add(1)
			go func(seat int){
				defer wg.Done()
				if err:= b.CancelSeat(seat); err != nil{
					fmt.Println(err)
				}
			}(seat)
		}
		wg.Wait()
	}
	if *showFlag{
		AvailableSeats:= b.ShowAvailableSeats()
		for _, seat := range AvailableSeats{
			fmt.Println(seat)
		}
	}
}


func ProperFormat(){
	fmt.Println("Proper format for book: -book 1,2,3")
	fmt.Println("Proper format for cancel: -cancel 1,2,3")
	fmt.Println("Proper format for show: -show")
}

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