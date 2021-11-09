# closest_in_tree
coding challenge problem

Find the nearest the number to the provided number in a binary tree. This problem came from a source I pay for so I won't say much more than that.

Though I am familiar with and have used similar node/graph based data structures in databases and in python, it has been a while (mid '90s) since I worked with the concept this granularly and my first time in Go.

Initially, I struggled to translate how I would have handled this in Python - with one or more classes - into a feasible Go implementation. Ultimately, I focused on lots of research to find how others had dealt with trees in the language. With this practice and numerous pseudocode revisions, I implemented this solution using Go structs.

Interestingly, the real trick was figuring out that one needs to determine whether to continue recursing using the absolute value of the next and current node. I won't spoil it in case you want to try solving the problem yourself.
