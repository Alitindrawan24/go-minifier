# Go Minifier

Go Minifier is a lightweight command-line tool designed to simplify and streamline the minification process for CSS and JavaScript files. Minification involves the removal of unnecessary characters and spaces from code, resulting in smaller file sizes that contribute to improved web page loading times.

### 1. Requirement
- Go

### 2. Installation & Setup

- Install Go Minifier directly from command line
```bash
go install github.com/Alitindrawan24/go-minifier
```

- For local development, clone project from the repository using http
```bash
git clone https://github.com/Alitindrawan24/go-minifier.git
```

- Clone project from the repository if using ssh

```bash
git clone git@github.com:Alitindrawan24/go-minifier.git
```

### 3. Usage

#### Minify CSS Files
- To minify a CSS file, run the following command. The minified file will be generated in the same directory with the same name but using `.min.css` extension

```bash
go run . -src filename.css
```

- To customize the output filename, run the following command

```bash
go run . -src filename.css -out custom.min.css
```

#### Minify JavaScript Files
- To minify a JavaScript file, run the following command. The minified file will be generated in the same directory with the same name but using `.min.js` extension

```bash
go run . -src filename.js -opt js
```

- To customize the output filename for JavaScript files, run the following command

```bash
go run . -src filename.js -out custom.min.js -opt js
```

#### Command Line Options
- `-src`: Path to the source file (required)
- `-out`: Path to the output file (optional, defaults to input filename with .min extension)
- `-opt`: Minifier option - 'css' for CSS files (default), 'js' for JavaScript files

### 4. Todo
- Minify all css/js in a directory