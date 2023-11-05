Usage: filescanner [flags] [directory]

Description:
filescanner is a Go command-line application that performs two tasks, which can be controlled using the "--task" flag.

Flags:
  -h, --help          Show this help message.
  -d, --directory     The directory to scan for text files (default is the current directory).
  --task              Toggle between tasks: Choose task 'line' or 'word'. ( default is line)
                        Task 1 : Choosing 'line' will read files and number of lines in each file.
                        Task 2 : Choosing 'word' will read files and print top 10 frequent word in files count.
  --ext"            Provide the comma seprated file extension to read data ( default is .txt)

Example Usage:
  - To scan the current directory for text files and display the top 10 most frequent words:
    $ go run main.go --task word


  - To display the numbers of lines in each txt file in the current directory:
    $  run main.go --task line
    Or
    $  run main.go 

  - To change the directory for files and display the top 10 most frequent words:
    $ go run main.go --task word -d ./path/to/desired/dir

  - To change the directory for files and display the numbers of lines in each txt file in the current directory:
    $ go run main.go --task line -d ./path/to/desired/dir
    Or
    $ go run main.go  -d ./path/to/desired/dir

  - To provide the different file extension for txt files.:
    $ go run main.go --ext .txt,.docs,.docx


