// Package new_year_chaos contains the solution for HackerRank challenge
// https://www.hackerrank.com/challenges/new-year-chaos/problem.
package new_year_chaos

import "errors"

// Run returns the total number of bribes for q or error if q is too chaotic.
func Run(q []int32) (int32, error) {
	// for each position from the end of the queue
	// 		call func(queue, position) {
	// 			if the their number is greater than the position then the person has bribed the person in front of them, so:
	//				- increment the total bribes counter
	//				- increment the person's bribes counter
	//					- if the person's bribes counter is greater than the maximum allowed bribes per person, then return error "Too chaotic"
	//				- move that person one position back in the queue to reset the queue back to the state before the bribe
	//				- func(queue, position + 1) // call the function for person that bribed
	//				- func(queue, position) // call the function for the person that was bribed
	//		}
	// 	return the total bribes counter

	const maxBribes = 2
	var totalBribes int32
	var bribes = map[int32]int{}
	var undoBribe func(q []int32, i int) error

	undoBribe = func(q []int32, i int) error {
		if i >= len(q) {
			// We can't move the last person any further back in the queue.
			return nil
		}

		n := q[i]
		if n <= int32(i+1) {
			// This person has been bribed and therefore has moved further back in the queue.
			return nil
		}

		if _, ok := bribes[n]; !ok {
			bribes[n] = 0
		}
		bribes[n]++
		totalBribes++
		if bribes[n] > maxBribes {
			return errors.New("Too chaotic")
		}

		// Undo the bribery for person with number n - move them one place further back in the queue.
		q[i], q[i+1] = q[i+1], q[i]

		// Callback to make sure we undo any other possible bribery by person with number n.
		if err := undoBribe(q, i+1); err != nil {
			return err
		}

		// Callback to make sure we check the place of the person that has been bribed.
		if err := undoBribe(q, i); err != nil {
			return err
		}

		return nil
	}

	for i := len(q) - 1; i >= 0; i-- {
		if err := undoBribe(q, i); err != nil {
			return 0, err
		}
	}

	return totalBribes, nil
}
