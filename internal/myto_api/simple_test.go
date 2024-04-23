package myto_api

import (
	"os"
	"testing"
)

func TestMytoExecJsonLike(t *testing.T) {
	// Set cwd to the root of the project
	os.Chdir("../../")
	// Restore the cwd after the test
	t.Cleanup(func() {
		// Reset cwd
		os.Chdir("./internal/myto_api")
	})

	// Run MytoExec
	res, err := MytoFileExec("./testdata/test_json_like.myto", "./testdata/.myto")
	if err != nil {
		t.Error(err)
	}

	// Check string length
	if len(res) == 0 {
		t.Fatal("Expected non-empty string, got empty string")
	}
	// Debug logging for result
	t.Log("\n" + res + "\n")
}

func TestMytoExecTxtLike(t *testing.T) {
	// Set cwd to the root of the project
	os.Chdir("../../")
	// Restore the cwd after the test
	t.Cleanup(func() {
		// Reset cwd
		os.Chdir("./internal/myto_api")
	})

	// Run MytoExec
	res, err := MytoFileExec("./testdata/test_txt_like.myto", "./testdata/.myto")
	if err != nil {
		t.Error(err)
	}
	// Check string length
	if len(res) == 0 {
		t.Fatal("Expected non-empty string, got empty string")
	}
	// Debug logging for result
	t.Log("\n" + res + "\n")
}
