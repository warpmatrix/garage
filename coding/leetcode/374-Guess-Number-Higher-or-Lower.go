/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    // ziNum: zero-indexed number
    return 1 + sort.Search(n, func(ziNum int) bool {
        return guess(ziNum + 1) <= 0
    })
}
