# Parking Lot Design

## To run:
Create a folder MainDirectory. Create two folders "MainDirectory/bin" and "MainDirectory/src".
Clone this repository in src folder. Set your GOPATH to the current MainDirectory's path.

## Command :
1. Create a build.sh file inside bin folder:

                                             # Setting GOPATH
                                             export GOPATH=`pwd`
                                             echo $GOPATH
                                             PROJECT="parking_lot_design"
                                             echo "Building $PROJECT\n"
                                             go install $PROJECT

                                             # Running test cases
                                             ./bin/parking_slot.test

                                             # For checking coverage of test cases
                                             cd src/parking_lot_design/parking_lot
                                             go test -cover

                                             # For creating executable test file
                                             # go test -c
2. ./bin/build.sh : This will create parking_lot_design.exe executable file.
3. ./bin/parking_lot_design : This will run the main.go file inside parking-lot-design package.

## Input
You can give input in two ways:
1. Via input file: Place the .txt file inside the static folder. You can also replace the content
                   of already existing "commands.txt" file
2. Via interactive command line by entering commands line by line

## Output
Output will be printed on the console.
