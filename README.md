# Rationale and problem-solving

As this was meant to be a whiteboardish assessment, I would like to use this
section to walk you through the process that I used.

When I see 'Fibonacci', I grin with the same giddiness that I did the first
time I finally grokked the concept of recursion. Recursion makes a relatively
quick job out of the Fibonacci series. However, recursive calls greatly increase
overhead. I decided that I wanted to show off my recursion chops, and my optimization
chops, to boot.

This solution crawls a binary tree to calculate a number at Fibonacci[n]
by adding the sums of the tree nodes, working backwards from the greatest Fibonacci index
to the smallest. As nodes split out left and right, there will be repeated calculations.
To reduce this overhead, a cache is employed. Each node is checked against the cache. If
the number is already cached, no calculation is required - the cached value is returned.
In the event of a cache miss, the new Fibonacci number is calculated and cached.

Now, I had a pretty nice way to tell me what the /last/ number in the Fibonacci sequence was.
However, the requirements state that I must /expand/ the sequence up to n, not calculate
the final number. I employed a buffer in the form of a slice of uint64 to solve this. For every
cache miss, a new number is pushed to the buffer. At the end of recursion, a function is
returned that returns the buffer.

The last change that I made was to change all of the int datatypes to uint64 and catch negs.
Since the logic driving this calculator would have to be altered to return negative Fibonacci
sequences, this is a modicum of protection. The real solution would be to refactor the code
to gracefully handle negative Fibonacci sequences, but that was not a requirement.

Finally, to make sure everything was wired properly, I wrote tests.
