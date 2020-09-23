# PrimeOrNot

PrimeOrNot is a simple tool which tells you whether an integer number you give is prime or not. When the number is not prime, it tells you the divisors for the number as well.

## Usage


```bash
$ go run primeornot.go
```

Just execute the above command in the command line on Terminal or an equivalent terminal shell. Make sure the primeornot.go file is in the directory where you are executing the command.

When you do it, you have the below prompt message:

```bash
Enter a positive integer number > 
```

When the message appears, you can input a number and hit return key.
Then you'll get the message either

```bash
XXX is a prime number. 
```

or

```bash
XXX is not a prime number. 
Divisors for XXX: 1, A, B, C, ... , XXX
```

where XXX is the number you have input, and A, B, C, ... are the divisors for XXX.
