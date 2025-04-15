package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

// App struct
type App struct {
	ctx           context.Context
	ticker        *time.Ticker
	count         int
	humanReadeble string
	startingTime  time.Time
	expiredIn     time.Duration
	remainingTime *string
}

type Emiting struct {
	Count         int
	HumanReadable string
	RemainingTime *string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Print() string {
	return a.humanReadeble
}

func (a *App) GetCount() int {
	return a.count
}

func (a *App) StopTicker() {
	d := "00m00.00"
	a.remainingTime = &d
	if a.ticker != nil {
		a.ticker.Stop()
	}
	fmt.Println("Ticker stopped")
}

func (a *App) StartTickerLoop(interval string, expiration string) {
	a.remainingTime = def
	t, err := time.ParseDuration(interval)
	if err != nil {
		fmt.Println("Parse duration error")
		return
	}

	expiration_time, err := time.ParseDuration(expiration)
	if err != nil {
		fmt.Println("Parse duration error")
		return
	}

	expiredAt := time.Now().Add(expiration_time)
	a.ticker = time.NewTicker(t)
	a.count = 0
	go func() {
		for {
			select {
			case <-a.ticker.C:
				// Do something on each tick
				fmt.Println("going", a.count)
				a.count++
				r := time.Since(expiredAt).String()
				a.remainingTime = &r
				if time.Now().Add(1 * time.Second).After(expiredAt) {
					fmt.Println("stop")
					s := "00:00.00"
					a.remainingTime = &s
					a.ticker.Stop()
					return
				}
			}
		}
	}()

}

func (a *App) StartStopwatch(interval string) {
	t, err := time.ParseDuration(interval)
	if err != nil {
		fmt.Println("Parse duration error")
	}
	ticker := time.NewTicker(t)
	fmt.Println(interval)
	go func() {
		for {
			t := <-ticker.C
			a.humanReadeble = t.Local().Format("15:04:05")
			runtime.EventsEmit(a.ctx, "timer-update", Emiting{Count: a.count, HumanReadable: a.humanReadeble, RemainingTime: a.remainingTime})
		}
	}()
}
