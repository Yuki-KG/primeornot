# PrimeOrNot

# This program makes a decision for a given positive integer number if it is a prime number or not.
# If the number isn't prime, this program returns all of its divisors.

class PrimeOrNot:
    def __init__(self):
        self.number = 0
        self.divisors = []
        self.divisorsChar = ''

    def getNumber(self):
        self.number = int(input('Enter an integer number > '))

    def getDivisors(self, number):
        self.divisors = [1]
        if number == 1:
            self.divisors.append(0)
        else:
            for i in range(2, number+1):
                if number % i == 0:
                    self.divisors.append(i)
            for j in self.divisors:
                self.divisorsChar += str(j) + '  '



# Press the green button in the gutter to run the script.
prm = PrimeOrNot()

while(prm.number < 1):
    prm.getNumber()
    if prm.number < 1:
        print('**Error: Number must be positive.')

prm.getDivisors(prm.number)

if prm.divisors[1] == prm.number:
    print('{0} is a prime number.'.format(str(prm.number)))
else:
    print('{0} is not a prime number.'.format(str(prm.number)))
    if prm.number != 1:
        print('Divisors for {0}: {1}'.format(str(prm.number), prm.divisorsChar))

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
