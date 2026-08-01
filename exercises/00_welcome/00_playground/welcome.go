// This is a Golang file. It is a plain text file with a `.go` extension.
// Every modern language supports comments. It is obvious that it would be no different with GO.
// Comments are ignored by the compiler; you can leverage them to annotate code with notes and
// explanations like this.
//
// Go provides C-style /* */ block comments and C++-style // line comments,
// the standard practice is to use line comments.
// block comments appear mostly as package comments,
// but are useful within an expression or to disable large swaths of code
//
// Exercises will include `TODO` or `__` markers to draw your attention to the lines
// where you need to write code.
// You'll need to replace these markers with your own code to complete the exercise.
// Sometimes it'll be enough to write a single line of code, other times you'll have to write
// longer sections.
//
// Each directory will contain the file you will edit and the test file.
//
// To check your solution, run the test file for this exercise:
//
//     go test ./exercises/01_intro/00_playground/
//
// Or, from inside this directory:
//
//     go test .
//
// Add -v to see each individual test case run:
//
//     go test -v .
//
// Once you've solved several exercises, you can check them all at once
// from the root of the repository:
//
//     go test ./...
//
// If you make a mistake, the output will look something like this:
//
// [bruno@arch-btw 00_playground]$ go test .
// --- FAIL: TestGreeting (0.00s)
//    welcome_test.go:10: got "I'm super excited to learn __!", want "I'm super excited to learn GO!"
// FAIL
// FAIL	github.com/xenodesire/go-hands-on/exercises/00_welcome/00_playground	0.002s
// FAIL
//
// But don't be alarmed—you'll learn by doing. For this first exercise, the expected output is:
//
// ok  	github.com/xenodesire/go-hands-on/exercises/00_welcome/00_playground	0.002s

package playground

func Greeting() string {

	// TODO: fix-me 👇
	return "I'm super excited to learn __!"
	// TIP: the two-letter language name
}
