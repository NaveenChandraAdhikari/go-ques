package main


import (
"context"
"fmt"
"time"
)

//this func  simulates a task that takes time and can be canceled.....
func longRunningTask(ctx context.Context, taskID string) error {
	fmt.Printf("task %s started\n", taskID)
	for i := 1; i <= 5; i++ {
		select {

		//simulate the tasks with a with a 500ms delay per iteration
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("task %s: Step %d completed\n", taskID, i)
		case <-ctx.Done():

			//here context is canceled we return the cancellation err
			fmt.Printf("task %s canceled\n", taskID)
			return ctx.Err()
		}
	}
	fmt.Printf("task %s finished successfully\n", taskID)
	return nil
}

func main() {
	// create a context with a 1sec timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() //to release resources

	// run  the task in a goroutine
	done := make(chan error)
	go func() {
		done <- longRunningTask(ctx, "task1")
	}()

	// Wait for the task to complete or be canceled
	select {
	case err := <-done:
		if err != nil {
			fmt.Printf("main-- Task ended with error: %v\n", err)
		} else {
			fmt.Println("main--- Task completed successfully")
		}
	case <-ctx.Done():
		fmt.Println("main---- Context timeout triggered")
	}
}