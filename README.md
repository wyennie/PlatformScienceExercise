## Assumptions
- the shipment addresses are in the same city, and therefore it is a list of 
  street addresses (no city, country, zip, etc.)
- the number of shipments is equal to the number of drivers
## File structure
ProjectScienceExercise
├── cmd
│   └── main.go
├── internal
|   ├── matrices/
|   │   ├── matrices.go
|   └── utils/
|       ├── calculateSS.go
└── README.md


## cmd/main.go
This is the entry point into the application. It takes in the txt file arguments
when the program is run, calls the appropriate functions, and outputs the
desired (hopefully!) results.
## internal
The rest of the code is stored here, as it's only intended to interact within
the project.
### matrices
Takes in the converted txt files and returns two matrices. Matrix A is filled
with Suitability Scores. Matrix B is a version of Matrix A that is altered so
it works in the hungarian algorithm.
## utils/calculateSS
The CalculateSS function takes two strings (dr)

## text
These are just two sample files we can use to test the app. 
Random addresses curtesy of https://www.randomlists.com/random-addresses
Random name curtesy of https://www.randomlists.com/random-names
1. get shipment destinations and drivers names in there own lists

2. create a matrix (let's call it matrixA) of Suitability Scores out of the lists from step 1
  we need to iterate through the shipment destinations list
    create a list
    inside current iteration, iterate through the drivers list
      push values to the list according to the top-secret algorithm
    push list to matrixA

This creates a matrix where the rows represent the shipment destinations and 
the columns represent the drivers

3. Make another matrix (let's call it matrixB) that uses the values in matrixA.
We neet to prep matrixB so that it's suitable to pass into the hungarian
algorithm. We multiply each score by 100, this is so we can convert them from
float64 to int. We then subtract the score from a very large number (99,999),
this is because the hungarian algorithm finds the minimum but we want to find
the maximum.

4. Pass matrixB into the Hungarian Algorithm. It will return a slice (let's call
it answerKey) where the index represents the row of matrixB and the value
represents the column of matrixB.

5. We can use the answerKey to add up the Suitability Score, by comparing it
with matrixA and at the same time Print/write out our pairings by comparing it
with the lists from step one

  The top-secret algorithm is:
  street names will be the characters from the begginning of a line to the first ",".
  ● If the length of the shipment's destination street name is even, the base suitability score (SS) is the number of vowels in the driver’s name multiplied by 1.5.
  ● If the length of the shipment's destination street name is odd, the base SS is the number of consonants in the driver’s name multiplied by 1.
  ● If the length of the shipment's destination street name shares any common factors (besides 1) with the length of the driver’s name, the SS is increased by 50% above the base SS.

Example of typical project structure
Project
  cmd
    main
      main.go
  pkg
    utils
    confing
    models
    routes
    controllers
  gomod
  README.md



Grok the answer

Shipments
  272 W. San Juan Dr.
  19 E. Lakeview Drive
  7847 Acacia Ave.
Drivers
  Ava Blevins     11  1, 11
  Jaxson Griffith 15  1, 3, 5, 15
  Marlon Navarro  14  1, 2, 7, 14


Answer Key
  [1 2 0]

19 11                     6 consonants
272 W. San Juan Dr.     Ava Blevins
baseSS   6
w/Factor 6

19 15                     10 consonants
272 W. San Juan Dr.     Jaxson Griffith
baseSS   10
w/Factor 10

19 14                     8 consonants
272 W. San Juan Dr.     Marlon Navarro
baseSS   8
w/Factor 8




20 11                     4 vowels
19 E. Lakeview Drive    Ava Blevins
baseSS   6
w/Factor 6

20 15                     4 vowels
19 E. Lakeview Drive    Jaxson Griffith
baseSS   6
w/Factor 9

20 14                     5 vowels
19 E. Lakeview Drive    Marlon Navarro
baseSS   7.5
w/Factor 11.25

[6, 10, 8]
[6, 9, 11.25]


16 11                   4 vowels
7847 Acacia Ave.        Ava Blevins
baseSS   6
w/Factor 6

16 15                   4 vowels
7847 Acacia Ave.        Jaxson Griffith
baseSS   6
w/Factor 6

16 14                   5 vowels
7847 Acacia Ave.        Marlon Navarro
baseSS   7.5
w/Factor 11.25

[6, 10, 8]
[6, 9, 11.25]
[6, 6, 11.25]