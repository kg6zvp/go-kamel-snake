package kamel_snake

import(
	"testing"
)

func AssertEqual(t *testing.T, e string, a string) {
	if e != a {
		t.Errorf("Expected: %s, Actual: %s", e, a);
	}
}

func TestCamelToSnake(t *testing.T) {
	AssertEqual(t, "test_string", CamelToSnake("TestString"));
	AssertEqual(t, "other_test", CamelToSnake("otherTest"));
	AssertEqual(t, "third_test_string", CamelToSnake("ThirdTestString"));
	AssertEqual(t, "none", CamelToSnake("None"));
	AssertEqual(t, "none", CamelToSnake("none"));
}

func TestSnakeToCamel(t *testing.T) {
	AssertEqual(t, "TestString", SnakeToCamel("test_string"));
	AssertEqual(t, "OtherTest", SnakeToCamel("other_test"));
	AssertEqual(t, "ThirdTestString", SnakeToCamel("third_test_string"));
	AssertEqual(t, "None", SnakeToCamel("none"));
}
