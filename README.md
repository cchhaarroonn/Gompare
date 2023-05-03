# Gompare

Gompare is a command-line tool written in Go that allows you to compare the contents of two directories and check if they have the same files. It does this by calculating the MD5 hash of the filenames in each directory and comparing them. If there are any files missing or if the hashes don't match, Gompare will print an error message.

## Installation

To use Gompare, you need to have Go installed on your system. Once Go is installed, you can download and install Gompare using the following command:

go get https://github.com/cchhaarroonn/Gompare.git


## Configuration

To run this program just configure the paths to directories you want to compare basically just paste paths in variables sourceDir and targetDir which are on lines 17 and 18 in main.go

## Usage

1. Download the gompare.go file.
2. Open your terminal and navigate to the directory where you saved the gompare.go file.
3. Run the following command: go run gompare.go

## Contributing

If you find any bugs or issues with Gompare, please feel free to open an issue on the GitHub repository. Pull requests are also welcome!

## Note

This project is rewrite of @narumii File comparator that he made in java so i rewrote it in Go just for fun and for my practice of learning Go!
