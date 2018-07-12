# 30 Days of Code
[Hacker Rank](https://www.hackerrank.com) Practice challengers

## How to run test codes:
Change GOPATH to 30-days-of-code/go directory:

```bash
export GOPATH=<your-project-folder>/coding/30-days-of-code/go
```

Enter on directory for each day and run `go test`

## Day 0: Hello, World
To complete this challenge, you must save a line of input from stdin to a variable, print Hello, World. on a single line, and finally print the value of your variable on a second line.

**Sample Input**
```
Welcome to 30 Days of Code!
```
**Sample Output**
```
Hello, World. 
Welcome to 30 Days of Code!
```

- [go](go/src/day0/main.go)

## Day 1: Data types
Complete the code in the editor below. The variables *d*, *i*, and *s* are already declared and initialized for you. You must:

1. Declare 3 variables: one of type int, one of type double, and one of type String.
2. Read 3 lines of input from stdin (according to the sequence given in the *Input Format* section below) and initialize your variables.
3. Use the *+* operator to perform the following operations:
    1. Print the sum of *i* plus your int variable on a new line.
    2. Print the sum of *d* plus your double variable to a scale of one decimal place on a new line.
    3. Concatenate *s* with the string you read as input and print the result on a new line.

**Sample Input**
```
12
4.0
is the best place to learn and practice coding!
```
**Sample Output**
```
16
8.0
HackerRank is the best place to learn and practice coding!
```

- [go](go/src/day1/main.go)

## Day 2: Operators
Given the meal price (base cost of a meal), tip percent (the percentage of the meal price being added as tip), and tax percent (the percentage of the meal price being added as tax) for a meal, find and print the meal's total cost.

**Explanation**

Given:

`mealCost = 12, tipPercent = 20, taxPercent = 8`

Calculations:

```
tip = 12 * (20 / 100) = 2.4
tax = 12 * (8 / 100) = 0.96
totalCost = mealCost + tip + tax = 12 + 2.4 + 0.96 = 15.36
round(totalCost) = 15
```

`We round *totalCost* to the nearest dollar (integer) and then print our result`

**Sample Input**
```
12.00
20
8
```
**Sample Output**
```
The total meal cost is 15 dollars.
```

- [go](go/src/day2/main.go)

## Day 3: Intro to Conditional Statements
Given an integer, _n_, perform the following conditional actions:

* If _n_ is odd, print Weird
* If _n_ is even and in the inclusive range of _2_ to _5_, print Not Weird
* If _n_ is even and in the inclusive range of _6_ to _20_, print Weird
* If _n_ is even and greater than _20_, print Not Weird

Complete the stub code provided in your editor to print whether or not _n_ is weird.

**Sample Input 1**
```
3
```
**Sample Output 1**
```
Weird
```

**Sample Input 2**
```
24
```
**Sample Output 2**
```
Not Weird
```

- [go](go/src/day3/)

## Day 4: Class vs. Instance
Write a Person class with an instance variable, _age_, and a constructor that takes an integer, _initialAge_, as a parameter.
The constructor must assign _initialAge_ to _age_ after confirming the argument passed as _initialAge_ is not negative; if a 
negative argument is passed as _initialAge_, the constructor should set _age_ to _0_ and print `Age is not valid, setting age to 0.`.
In addition, you must write the following instance methods:

1. yearPasses() should increase the _age_ instance variable by _1_.
2. amIOld() should perform the following conditional actions:
    * If _age < 13_ , print `You are young.`.
    * If _age >= 13_ and _age < 18_, print `You are a teenager.`.
    * Otherwise, print `You are old.`.

**Sample Input**
```
4
-1
10
16
18
```
**Sample Output**
```
Age is not valid, setting age to 0.
You are young.
You are young.

You are young.
You are a teenager.

You are a teenager.
You are old.

You are old.
You are old.
```

- [go](go/src/day4/)

## Day 5: Loops
Given an integer, _n_, print its first _10_ multiples.
Each multiple _n X i_ (where _1 <= i <= 10_) should be printed on a new line in the form: `n x i = result`.

**Constraints**

* _2 <= n <= 20_

**Sample Input**
```
2
```
**Sample Output**
```
2 x 1 = 2
2 x 2 = 4
2 x 3 = 6
2 x 4 = 8
2 x 5 = 10
2 x 6 = 12
2 x 7 = 14
2 x 8 = 16
2 x 9 = 18
2 x 10 = 20
```

- [go](go/src/day5/)

## Day 6: Let's Review
Given a string, _S_, of length _N_ that is indexed from _0_ to _N - 1_, print its even-indexed and odd-indexed characters as _2_ space-separated strings on a single line.

**Note:** _0_ is considered to be an even index

**Input format**

The first line contains an integer, _T_ (the number of test cases). 
Each line _i_ of the _T_ subsequent lines contain a String, _S_.

**Constraints**

* _1 <= T <= 10_
* _2 <= length of S <= 10000_

**Output format**

For each String _Sj_ (where _0 <= j <= T - 1_), print _Sj_'s even-indexed characters, followed by a space, followed by _Sj_'s odd-indexed characters.

**Sample Input**
```
2
Hacker
Rank
```
**Sample Output**
```
Hce akr
Rn ak
```

**Explanation**

Test Case 0: _S = "Hacker"_ 

_S[0] = "H"_

_S[1] = "a"_

_S[2] = "c"_

_S[3] = "k"_

_S[4] = "e"_

_S[5] = "r"_


The even indices are _0_, _2_, and _4_, and the odd indices are _1_, _3_, and _5_.
We then print a single line of _2_ space-separated strings; the first string contains the ordered characters from _S_'s even indices (**Hce**), and the second string contains the ordered characters from _S_'s odd indices (**akr**).

Test Case 1:  _S = "Rank"_

_S[0] = "R"_ 

_S[1] = "a"_ 

_S[2] = "n"_ 

_S[3] = "k"_ 

The even indices are _0_ and _2_, and the odd indices are _1_ and _3_.
We then print a single line of _2_ space-separated strings; the first string contains the ordered characters from _S_'s even indices (**Rn**), and the second string contains the ordered characters from _S_'s odd indices (**ak**).

- [go](go/src/day6/)

## Day 7: Arrays
Given an array, _A_, of _N_ integers, print _A_'s elements in reverse order as a single line of space-separated numbers.

**Input Format**

The first line contains an integer, _N_ (the size of our array).
The second line contains _N_ space-separated integers describing array _A_'s elements.

**Sample Input**
```
4
1 4 3 2
```
**Sample Output**
```
2 3 4 1
```

- [go](go/src/day7/)
