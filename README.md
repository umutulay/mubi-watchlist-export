# mubi-watchlist-export

This project is a Go application that interacts with the MUBI API to retrieve a user's watchlist and export it into an Excel file. The application fetches films from the user's MUBI watchlist and extracts details such as the film title and year, then saves this information in a well-structured Excel spreadsheet.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- `excelize` Go module for working with Excel files
- `mubi` Go package for interacting with MUBI API

### Installing

1. Clone the repository

`git clone https://github.com/umutulay/mubi_watchlist_export.git`

2. Navigate to the project directory

`cd mubi_watchlist_export`

3. Initialize the Go module
`go mod init mubi_export.go`

4. Download and install required Go modules
`go mod tidy`

### Usage

Change the userId
`go run mubi_export.go`

## Authors

* Umut Tulay - [umutulay](https://github.com/umutulay) 
