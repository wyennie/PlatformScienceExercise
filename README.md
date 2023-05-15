# Contents
```
Contents
Run the Program!
  use your own text files
Assumptions
Structure
  cmd/main.go
  calculate
    SuitabilityScore
    FinalResults
    MakeMatrices
  io
    ParseFiles
  matrices
    MakeMatrices
Hungarian Algorithm
```

# Run the Program!
1. Clone the GitHub reopsitory.
```
$ git clone git@github.com:wyennie/PlatformScienceExercise.git
```

2. Download necessary modules.
```
$ go mod download
```
4. Run the code!

**You can run it from the build file:**
```
$ ./main testdata/shipments.txt testdata/drivers.txt
```

**Otherwise**
```
$ go run cmd/main.go testdata/shipments.txt testdata/drivers.txt
```

This will run the program with the smaple txt files from `testdata`.

When the code is run it will create a file called `output.txt` if it doesn't
exist and overwrite that file if it does exist. You can see the output in that
file as well as in the console.

**Use your own txt files**
The code should work with other txt files as well, the two text files should
both have the same amount of lines (and preferrably no empty lines). 
Make sure to pass in the list of shipment addresses first and the list of
drivers names second.
```
$ ./main <shipments txt file> <drivers txt file>
```
```
$ go run cmd/main.go <shipments txt file> <drivers txt file>
```

# Tests
If you w
I have made some simple tests for the `calculate`, `io`, and `matrices` files.
Commands to run some tests from inside the project directory.

# Assumptions
- Shipment destination addresses are in the same city, and therefore it is a 
  list of street addresses (no city, country, zip, etc.)
- The number of shipments is equal to the number of available drivers.

# Structure
ProjectScienceExercise
```
├── cmd
│   └── main.go
├── internal
|   ├── calculate/
|   │   ├── calculate_test.go
|   │   └── calculate.go
|   ├── io/
|   │   ├── fileio_test.go
|   │   └── fileio.go
|   └── matrices/
|       ├── matrices_test.go
|       └── matrices.go
├── testdata
└── README.md
```
I created this simple structure becuase it helps me parse the different 
components of the application. If the application were to get significantly
larger, it may be benificial to move some of the packages to more traditionally
named ones (utils, vendor, etc.). I used *internal* because the packages are not
for use outside of the project.

## cmd/main.go
This is the entry point into the application. It takes in the txt file arguments
when the program is run via cli, it calls the appropriate functions, and writes
the results to a txt file (output.txt).

## calculate
Calculate has two main functions:

### SuitabilityScore

**params**: shipmentDestination string, driverName string

**return**: float64

Calculates the Suitabillity Score (using the top-secret algorithm) between
the shipment address and the driver's name that are passed in and returns it.

### FinalResults

**params**: shipmentLines []string, driverLines []string, matrixA [][]float64, answerKey []int

**return**: string

Calculates the Total Suitability Score via answerKey and matrixA. Uses the
answerKey and the driver/shipment Lines slices to put together the final output
string. It prints the final output string and returns it.

## io
We use io to parse txt files and create lists of strings from them

### ParseFiles

**params**: shipmentFile string, driverFile string

**return**: []string, []string


Takes in two txt files (specifically the ones passed to our program when it's
run). The files are read and each line is appended to a list made for that file.
The two lists are returned.

## matrices
Responsible for the creation of the various matrices we use in this project.

### MakeMatrices

**params**: shipments []string, drivers []string

**return**: [][]float64, [][]int

Creates and returns two matrices:
  1. **suitabilityScoreMatrix**: A matrix that contains all of the suitability 
    scores between the shipment addressess and driver names
  2. **hungarianMatrix**: A matrix that takes the *suitabilityScoreMatrix* and
    creates a modified version of it that is suited for the hungarian algorithm.

# Hungarian Algorithm
The hungarian algorithm finds the smallest cost, but we want the largest,
The implementation of the hungarian algorithm I'm using also accepts a
matrix of int not float64. We multiply the score by 100 so we can convert
between int and float64 to the hundreth place. Then subtract each score from
a very large number. This will make the biggest the smallest 