package main

// Packages are, together with functions, about code reuse (DRY).

// Many of the libraries (packages) included with Go require specialized,
// domain specific knowledge (e.g. cryptography).

import ( "fmt" ; "strings" ; "os" ; "io/ioutil" ; "log" )

//
// Strings
//

// strings functions
func mystrings() {
    // func Contains(s, substr string) bool
    fmt.Println(strings.Contains("test", "es"))
    // => true

    // func Count(s, sep string) int
    fmt.Println(strings.Count("test", "t"))
    // => 2

    // func HasPrefix(s, prefix string) bool
    fmt.Println(strings.HasPrefix("test", "te"))
    // => true

    // func HasSuffix(s, suffix string) bool
    fmt.Println(strings.HasSuffix("test", "st"))
    // => true

    // func Index(s, substr string) int
    // -1 if not found
    fmt.Println(strings.Index("test", "e"))
    // => 1

    // func Join(a []string, sep string) string
    fmt.Println(strings.Join([]string{"a", "b"}, "-"))
    // => "a-b"

    // func Repeat(s string, count int) string
    fmt.Println(strings.Repeat("a", 5))
    // => "aaaaa"

    // func Replace(s, old, new string, n int) string
    fmt.Println(strings.Replace("aaaa", "a", "b", 2))
    // => "bbaa"

    // func Split(s, sep string) []string
    fmt.Println(strings.Split("a-b-c", "-"))
    // => []string{"a", "b", "c"}

    // func ToLower(s string) string
    fmt.Println(strings.ToLower("TEST"))
    // => "test"

    // func ToUpper(s string) string
    fmt.Println(strings.ToUpper("test"))
    // => "TEST"

    // Sometimes we need to work with strings as binary data. To convert a
    // string to a slice of bytes (and vice versa):
    arr := []byte("test")
    str := string([]byte{'T','E','S','T'})
    fmt.Println(arr)
    fmt.Println(str)
}

//
// Files and folders
//

// Open and print a file
func readFile (fileName string) {
    file, err := os.Open(fileName)
    if err != nil {
        // handle the error here
        return
    }
    defer file.Close()

    // get the file size
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // read the file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}

// Shorter way to open and print a file
func readFile2 (fileName string) {
    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("File contents: %s", content)
}

// Get the contents of a directory
func readDir (dirname string) {
    dir, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    defer dir.Close()

    fileInfos, err := dir.Readdir(-1) // -1 => all dir entries
    if err != nil {
        log.Fatal(err)
    }
    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}

func main() {
    mystrings()

    readFile("/etc/hosts")
    readFile2("/etc/hosts")

    readDir(".")
}
