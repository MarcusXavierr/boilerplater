package boilerplater

import (
	"os"
	"reflect"
	"testing"
)

func TestParseString(t *testing.T) {
	check := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Error: got %q, want %q", got, want)
		}
	}
	t.Run("parse normal string", func(t *testing.T) {
		php := Php{Path: "tests/Feature/Withdraw/WithdrawRequestHistoryServiceTest.php"}
		namespace, className := php.ParsePath()

		wantedNamespace := "Tests\\Feature\\Withdraw"
		wantedClassName := "WithdrawRequestHistoryServiceTest"

		check(t, namespace, wantedNamespace)
		check(t, className, wantedClassName)
	})

	t.Run("parse fullpath string", func(t *testing.T) {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		php := Php{Path: cwd + "/tests/Feature/Withdraw/WithdrawRequestHistoryServiceTest.php"}
		namespace, className := php.ParsePath()

		wantedNamespace := "Tests\\Feature\\Withdraw"
		wantedClassName := "WithdrawRequestHistoryServiceTest"

		check(t, namespace, wantedNamespace)
		check(t, className, wantedClassName)
	})

	t.Run("remove string filename", func(t *testing.T) {
		str := "tests/Feature/Withdraw/WithdrawRequestHistoryServiceTest.php"
		path, filename := separateFileFromPath(str)

		wantedPath := []string{"tests", "Feature", "Withdraw"}
		wantedFilename := "WithdrawRequestHistoryServiceTest.php"

		check(t, filename, wantedFilename)
		if !reflect.DeepEqual(path, wantedPath) {
			t.Errorf("got %v wanted %v", path, wantedPath)
		}
	})

	t.Run("Create namespace", func(t *testing.T) {
		path := []string{"tests", "Feature", "Withdraw"}
		got := createNamespace(path)
		want := "Tests\\Feature\\Withdraw"

		check(t, got, want)
	})

	t.Run("Remove file extension", func(t *testing.T) {
		got := removeFileExtension("WithdrawRequestHistoryServiceTest.php")
		want := "WithdrawRequestHistoryServiceTest"
		check(t, got, want)
	})
}

func TestCreateBoilerplateString(t *testing.T) {
	check := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Error: got %q, want %q", got, want)
		}
	}
	t.Run("create string for laravel class", func(t *testing.T) {
		php := Php{Path: "tests/Feature/Withdraw/WithdrawRequestHistoryServiceTest.php"}
		got := php.CreateBoilerplateString()
		want := "<?php\nnamespace Tests\\Feature\\Withdraw;\n\nclass WithdrawRequestHistoryServiceTest\n{\n}"
		check(t, got, want)
	})
}
