
# gaf

Draw ASCII bar charts from a file


## Installation

### go get

```go
∃! go get github.com/Xx0w0wxX/gaf/cmd/gaf
```

## Usage

```go
∃! cat sample.csv
python,64
go,18
flutter,57
javascript,97
html,58
css,59
php,74

∃! gaf sample.csv
    python :  █████████████████████▎
        go :  ██████
   flutter :  ███████████████████
javascript :  ████████████████████████████████▎
      html :  ███████████████████▎
       css :  ███████████████████▋
       php :  ████████████████████████▋
```

## LICENSE

CC-NC




## Reference

- https://qiita.com/nayuneko/items/3c0b3c0de9e8b27c9548
- https://github.com/Xx0w0wxX/til/blob/master/go/read-csv.md
- https://stackoverflow.com/questions/33139020/can-golang-multiply-strings-like-python-can
- https://stackoverflow.com/questions/39245610/golang-converting-from-rune-to-string
- https://stackoverflow.com/questions/29914662/equivalent-of-pythons-ord-chr-in-go
- https://golang.org/doc/code.html