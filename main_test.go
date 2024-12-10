package main

// func TestMain(t *testing.T) {
// 	expected := "Hello, World!\n"
// 	out, err := exec.Command("go", "run", "main.go").Output()
// 	if err != nil {
// 		t.Errorf("Error running main.go: %v", err)
// 	}
// 	outstr := string(out[:])
// 	if outstr != expected {
// 		t.Errorf("Error running main.go: expected output to be %q, received %q instead", expected, outstr)
// 	}
// }

// func TestOneArg(t *testing.T) {
// 	expected := "starting crawl of: http://example.com"
// 	out, err := exec.Command("go", "run", "main.go", "http://example.com").Output()
// 	if err != nil {
// 		t.Errorf("Error running main.go: %v", err)
// 	}
// 	outstr := strings.TrimSpace(string(out[:]))
// 	if outstr != expected {
// 		t.Errorf("Error running main.go: expected output to be %q, received %q instead", expected, outstr)
// 	}
// }

// There's no good (to me) way to assert exist status 1
//
// func TestNoArgs(t *testing.T) {
// 	expected := "no website provided"
// 	out, err := exec.Command("go", "run", "main.go").Output()
// 	if err != nil {
// 		t.Errorf("Error running main.go: %v", err)
// 	}
// 	outstr := strings.TrimSpace(string(out[:]))
// 	if outstr != expected {
// 		t.Errorf("Error running main.go: expected output to be %q, received %q instead", expected, outstr)
// 	}
// }

// func TestTooManyArgs(t *testing.T) {
// 	expected := "too many arguments provided"
// 	out, err := exec.Command("go", "run", "main.go", "asdf", "qwer").Output()
// 	if err != nil {
// 		t.Errorf("Error running main.go: %v", err)
// 	}
// 	outstr := strings.TrimSpace(string(out[:]))
// 	if outstr != expected {
// 		t.Errorf("Error running main.go: expected output to be %q, received %q instead", expected, outstr)
// 	}
// }
