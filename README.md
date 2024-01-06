# Golang Library Command Line Project

## About the project

This application provides two featured library appliation from command line. This app take 6 commands such as **list** , **search**, **buy**, **delete**, **add** and **update**.

**List:** This command prints all books in the system.
```bash
foo@bar:~$ go run main.go list
1) Book1
2) Book2
foo@bar:~$ _
```
**Search:** This command is used with at least one another argument that is egual or longer tahn 3 char. It retuns books that contain string that given by you from console.
```bash
foo@bar:~$ go run main.go search Harry Potter
1) BookID: 7, Name: Harry Potter and the Philosophers Stone, Author: J. K. Rowling, Price: 28.59, ISBN: 8980017, Stock Code: 9458116, Stock Num: 2
2) BookID: 8, Name: Harry Potter and the Prisoner of Azkaban, Author: J. K. Rowling, Price: 12.55, ISBN: 3193181, Stock Code: 5292996, Stock Num: 11
foo@bar:~$ _
```

**Buy:** Takes two paramters.First one is book id of book that is wanted decrease the stock. Second paramter is how many books do you want.
```bash
foo@bar:~$ go run main.go buy 5 3
Harry Potter and the Prisoner of Azkaban 3 adet satin alindi.
foo@bar:~$ _
```
**Delete:** Takes one paramter. It is the ID of book that is wanted deleted.
```bash
foo@bar:~$ go run main.go delete 5
Little Prince Basariyla silindi
foo@bar:~$ _
```

**Add:** Takes 7 paramteres with comma between them. It adds database the book.
```bash
foo@bar:~$ go run main.go add Kitap Adi,Yazar Adi, ISBN, Stock Code, Stock Quantity, Page Number, Price
Kitap basariyla eklendi!!!
foo@bar:~$ _
```

**Update:** Takes 2 paramteres. First one is book ID , second one is increase amount of stock quantity. This command increase stock quantity of book.
```bash
foo@bar:~$ go run main.go update 3 5
Harry Potter 5 tane artırıldi.
foo@bar:~$ _
```

In the other situations, program will print error messages and terminated.


## Notes

*   You can put the searchelper directory in src directory that is in `GOPATH` . In this case you need to update import statment.
```go
package main


import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	data "workspace/data"
	"searchhelper"
	model "workspace/models"
)
```


* If there is any problem accsesing package or GOPATH you can set again GOPATH and you can use this command below and go visit this [link](https://stackoverflow.com/questions/68693154/package-is-not-in-goroot).
```bash
foo@bar:~$ go env -w GO111MODULE=off
```