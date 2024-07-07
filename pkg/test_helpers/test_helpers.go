package test_helpers

import "testing"

func HasErrorMessage(t *testing.T, errs []error, expectedError string) {
	t.Helper()

	if len(errs) == 0 {
		t.Errorf("Expected to have received errors")
	}

	result := false
	for _, e := range errs {
		if e.Error() == expectedError {
			result = true
		}
	}

	if !result {
		t.Errorf("Expected error message: %s\nErrors received: %v", expectedError, errs)
	}
}

func BlankError(t *testing.T, errs []error) {
	t.Helper()

	if len(errs) > 0 {
		t.Errorf("Expected to receive a blank errors slice. Received: %v", errs)
	}
}
