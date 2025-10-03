package require

const (
	// strings for JSON tests.

	notJSONString            = "Not JSON"
	fooBarObject             = `{"foo": "bar"}`
	simpleJSONObject         = `{"hello": "world", "foo": "bar"}`
	simpleJSONObjectReversed = `{"foo": "bar", "hello": "world"}`
	simpleJSONNested         = `["foo", {"hello": "world", "nested": "hash"}]`
	simpleJSONNestedReversed = `["foo", {"nested": "hash", "hello": "world"}]`
	simpleJSONNestedNotEq    = `{"foo": "bar", {"nested": "hash", "hello": "world"}}`
	simpleJSONArray          = `["foo", {"hello": "world", "nested": "hash"}]`
	simpleJSONArrayReversed  = `[{ "hello": "world", "nested": "hash"}, "foo"]`

	nestedJSONObject = "{" +
		"\r\n\t\"numeric\": 1.5,\r\n\t\"array\": " +
		"[{\"foo\": \"bar\"}, 1, \"string\", " +
		"[\"nested\", \"array\", 5.5]],\r\n\t\"hash\": " +
		"{\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\"\r\n}"
	nestedJSONObjectShuffled = "{" +
		"\r\n\t\"numeric\": 1.5,\r\n\t\"hash\": " +
		"{\"nested\": \"hash\", \"nested_slice\": " +
		"[\"this\", \"is\", \"nested\"]},\r\n\t\"string\": " +
		"\"foo\",\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]]\r\n}"
)
