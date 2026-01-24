package playingwithjson

import (
	"encoding/json"
	"testing"
)

// TestMarshalPrivateFields tests that private fields are not marshaled
func TestMarshalPrivateFields(t *testing.T) {
	data := normalStruct{
		name:    "cat",
		email:   "cat@cat.cat",
		age:     5,
		address: "Catland",
	}

	response, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Private fields should result in empty JSON object
	expected := "{}"
	actual := string(response)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

// TestMarshalWithStructTags tests marshaling with struct tags but private fields
func TestMarshalWithStructTags(t *testing.T) {
	data := StructWithTick{
		name:  "aakku",
		email: "aakku106@gmail.com",
	}

	response, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Even with tags, private fields are not exported
	expected := "{}"
	actual := string(response)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

// TestMarshalPublicFields tests marshaling with public fields
func TestMarshalPublicFields(t *testing.T) {
	data := struct3{
		Name:    "Adarasha Gaihre",
		Address: "Nepal",
		age:     20, // Private field, won't be marshaled
	}

	response, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	actual := string(response)
	// Should only include public fields
	if actual != `{"Name":"Adarasha Gaihre","Address":"Nepal"}` {
		t.Errorf("Unexpected output: %s", actual)
	}

	// Verify private field 'age' is not included
	var result map[string]interface{}
	if err := json.Unmarshal(response, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	if _, exists := result["age"]; exists {
		t.Error("Private field 'age' should not be marshaled")
	}
}

// TestMarshalWithProperTags tests marshaling with proper struct tags
func TestMarshalWithProperTags(t *testing.T) {
	data := struct4{
		Name:    "Adarasha Gaihre",
		Address: "Nepal",
		Age:     20,
	}

	response, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	expected := `{"name":"Adarasha Gaihre","address":"Nepal","age":20}`
	actual := string(response)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

// TestMarshalIndent tests indented JSON output
func TestMarshalIndent(t *testing.T) {
	data := struct4{
		Name:    "Adarasha Gaihre",
		Address: "Nepal",
		Age:     20,
	}

	tests := []struct {
		name     string
		prefix   string
		indent   string
		expected string
	}{
		{
			name:     "No indent",
			prefix:   "",
			indent:   "",
			expected: "{\n\"name\": \"Adarasha Gaihre\",\n\"address\": \"Nepal\",\n\"age\": 20\n}",
		},
		{
			name:     "Two spaces",
			prefix:   "",
			indent:   "  ",
			expected: "{\n  \"name\": \"Adarasha Gaihre\",\n  \"address\": \"Nepal\",\n  \"age\": 20\n}",
		},
		{
			name:     "Custom indent",
			prefix:   "",
			indent:   "--->",
			expected: "{\n--->\"name\": \"Adarasha Gaihre\",\n--->\"address\": \"Nepal\",\n--->\"age\": 20\n}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := json.MarshalIndent(data, tt.prefix, tt.indent)
			if err != nil {
				t.Fatalf("MarshalIndent failed: %v", err)
			}

			actual := string(response)
			if actual != tt.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, actual)
			}
		})
	}
}

// TestUnmarshalUser tests unmarshaling JSON into User struct
func TestUnmarshalUser(t *testing.T) {
	jsonData := `{
		"id": 1,
		"name": "Leanne Graham",
		"age": 25,
		"address": {
			"email": "leanne@example.com",
			"location": {
				"country": "USA",
				"province": "California",
				"city": "San Francisco",
				"tole": "Market St"
			}
		}
	}`

	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if user.Id != 1 {
		t.Errorf("Expected Id to be 1, got %d", user.Id)
	}
	if user.Name != "Leanne Graham" {
		t.Errorf("Expected Name to be 'Leanne Graham', got %s", user.Name)
	}
	if user.Age != 25 {
		t.Errorf("Expected Age to be 25, got %d", user.Age)
	}
	if user.Address == nil {
		t.Fatal("Expected Address to be non-nil")
	}
	if user.Address.Email != "leanne@example.com" {
		t.Errorf("Expected Email to be 'leanne@example.com', got %s", user.Address.Email)
	}
	if user.Address.Location == nil {
		t.Fatal("Expected Location to be non-nil")
	}
	if user.Address.Location.Country != "USA" {
		t.Errorf("Expected Country to be 'USA', got %s", user.Address.Location.Country)
	}
}

// TestUnmarshalMultipleUsers tests unmarshaling array of users
func TestUnmarshalMultipleUsers(t *testing.T) {
	jsonData := `[
		{
			"id": 1,
			"name": "Alice",
			"age": 30
		},
		{
			"id": 2,
			"name": "Bob",
			"age": 25
		}
	]`

	var users []User
	err := json.Unmarshal([]byte(jsonData), &users)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("Expected 2 users, got %d", len(users))
	}

	if users[0].Id != 1 || users[0].Name != "Alice" {
		t.Errorf("First user data incorrect: %+v", users[0])
	}

	if users[1].Id != 2 || users[1].Name != "Bob" {
		t.Errorf("Second user data incorrect: %+v", users[1])
	}
}

// TestUnmarshalInvalidJSON tests error handling for invalid JSON
func TestUnmarshalInvalidJSON(t *testing.T) {
	invalidJSON := `{"id": 1, "name": "Invalid JSON"`

	var user User
	err := json.Unmarshal([]byte(invalidJSON), &user)
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

// TestMarshalUnmarshalRoundTrip tests that data survives marshal/unmarshal
func TestMarshalUnmarshalRoundTrip(t *testing.T) {
	original := struct4{
		Name:    "Test User",
		Address: "Test Address",
		Age:     99,
	}

	// Marshal
	jsonData, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result struct4
	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Compare
	if result != original {
		t.Errorf("Round trip failed. Expected %+v, got %+v", original, result)
	}
}
