package main

import "testing"

func TestJsMinifier_removeComments(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "single line comments",
			input: `
// This is a comment
var x = 5;
console.log(x); // inline comment
`,
			expected: "\n\nvar x = 5;\nconsole.log(x); \n",
		},
		{
			name: "multi line comments",
			input: `
/* This is a multi-line
   comment that spans
   multiple lines */
function test() {
	return true;
}
/* Another comment */
`,
			expected: "\n\nfunction test() {\n\treturn true;\n}\n\n",
		},
		{
			name: "comments with URLs in strings",
			input: `
const config = {
	apiUrl: "https://api.example.com/users", // API endpoint
	timeout: 5000, // timeout in ms
};
console.log("Visit https://example.com for more info"); // URL in string
`,
			expected: "\nconst config = {\n\tapiUrl: \"https://api.example.com/users\", \n\ttimeout: 5000, \n};\nconsole.log(\"Visit https://example.com for more info\"); \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsMinifier{&Minifier{Content: tt.input}}
			js.removeComments()
			if js.Content != tt.expected {
				t.Errorf("got: %q, want: %q", js.Content, tt.expected)
			}
		})
	}
}

func TestJsMinifier_removeWhiteSpace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "basic function",
			input: `
				function hello ( ) {
					console.log("Hello World");
				}
			`,
			expected: "function hello(){console.log(\"Hello World\");}",
		},
		{
			name: "variable declarations",
			input: `
				var x = 10 ;
				let y = 20 ;
				const z = 30 ;
			`,
			expected: "var x=10;let y=20;const z=30;",
		},
		{
			name: "if statement with operators",
			input: `
				if ( x > 5 && y < 10 ) {
					return x + y;
				}
			`,
			expected: "if(x>5&&y<10){return x+y;}",
		},
		{
			name: "object literal",
			input: `
				const obj = {
					name : "John" ,
					age : 25 ,
					city : "New York"
				} ;
			`,
			expected: "const obj={name:\"John\",age:25,city:\"New York\"};",
		},
		{
			name: "array with spaces",
			input: `
				const arr = [ 1 , 2 , 3 , 4 , 5 ] ;
			`,
			expected: "const arr=[1,2,3,4,5];",
		},
		{
			name: "function with parameters",
			input: `
				function calculate ( a , b ) {
					return a * b + 10 ;
				}
			`,
			expected: "function calculate(a,b){return a*b+10;}",
		},
		{
			name: "switch statement",
			input: `
				switch ( day ) {
					case 1 :
						console.log ( "Monday" ) ;
						break ;
					case 2 :
						console.log ( "Tuesday" ) ;
						break ;
					default :
						console.log ( "Other day" ) ;
				}
			`,
			expected: "switch(day){case 1:console.log(\"Monday\");break;case 2:console.log(\"Tuesday\");break;default:console.log(\"Other day\");}",
		},
		{
			name:     "multiple spaces and newlines",
			input:    "  var   x    =    5   ;  \n\n  console.log(x);  ",
			expected: "var x=5;console.log(x);",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsMinifier{&Minifier{Content: tt.input}}
			js.removeWhiteSpace()
			if js.Content != tt.expected {
				t.Errorf("got: %q, want: %q", js.Content, tt.expected)
			}
		})
	}
}

func TestJsMinifier_Minify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "javascript with comments",
			input: `
// This is a single line comment
function greet ( name ) {
	/* This is a multi-line
	   comment that spans
	   multiple lines */
	console.log ( "Hello " + name ) ; // inline comment
	return true ; // another comment
}
/* Final comment */
`,
			expected: "function greet(name){console.log(\"Hello \"+name);return true;}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := JsMinifier{&Minifier{Content: tt.input}}
			js.Minify()
			if js.Content != tt.expected {
				t.Errorf("got: %q, want: %q", js.Content, tt.expected)
			}
		})
	}
}
