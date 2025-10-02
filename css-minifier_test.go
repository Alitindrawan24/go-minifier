package main

import "testing"

func TestCssMinifier_removeWhiteSpace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "basic rule",
			input: `
				body {
					color : red ;
				}
			`,
			expected: "body{color:red;}",
		},
		{
			name: "multiple selectors",
			input: `
				h1, h2, h3 {
					margin : 0 ;
					padding : 0 ;
				}
			`,
			expected: "h1,h2,h3{margin:0;padding:0;}",
		},
		{
			name: "nested braces",
			input: `
				@media screen and (max-width: 600px) {
					body {
						background-color : lightblue ;
					}
				}
			`,
			expected: "@media screen and (max-width:600px){body{background-color:lightblue;}}",
		},
		{
			name: "trailing semicolon before brace",
			input: `
				p {
					font-size: 14px ;
				}
			`,
			expected: "p{font-size:14px;}",
		},
		{
			name: "url with spaces inside quotes should stay untouched",
			input: `
				.bg {
					background: url("my image.png") no-repeat center center ;
				}
			`,
			expected: `.bg{background:url("my image.png") no-repeat center center;}`,
		},
		{
			name: "keep important",
			input: `
				.hidden {
					display : none !important ;
				}
			`,
			expected: ".hidden{display:none!important;}",
		},
		{
			name:     "multiple spaces and newlines",
			input:    "  div   {   color    :    blue   ;   }\n\n",
			expected: "div{color:blue;}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			css := CssMinifier{&Minifier{Content: tt.input}}
			css.removeWhiteSpace()
			if css.Content != tt.expected {
				t.Errorf("got: %q, want: %q", css.Content, tt.expected)
			}
		})
	}
}
