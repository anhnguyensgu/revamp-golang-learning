(Optional) O(1) extra space trick

You can do a one-pass mountain counting (track lengths of up-slope and down-slope, add triangular numbers, and fix the peak). It stays O(n) time and O(1) space but is easy to bug; the two-sweep method above is the interview staple.