// For loop
for i := 0; i < 10; i++ {
		// Do a thing
	}

// loops
// Do
do {
    work();
} while (condition)

// infinite For
for ok := true; ok; ok = condition {
    work()
}

//OR 

for {
    work()
    if !condition {
        break
    }
}

// repeat
repeat
    work();
until condition;

// Switch statment
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
}