# Go Minifier

Go Minifier is is a lightweight command-line tool designed to simplify and streamline the minification process for CSS and JS (upcoming) files,. Minification involves the removal of unnecessary characters and spaces from code, resulting in smaller file sizes that contribute to improved web page loading times.

### 1. Requirement
- Go

### 2. Installation & Setup
- Clone project from the repository using http
```bash
git clone https://github.com/Alitindrawan24/go-minifier.git
```

- Clone project from the repository if using ssh

```bash
git clone git@github.com:Alitindrawan24/go-minifier.git
```

### 3. Usage
- To minify a css file, run the following command and the minify file will be generated on the same directory with same name as the css file but using min.css extension

```bash
go run . -src=filename.css
```

- To customize the output filename and extension, run the following command

```bash
go run . -src=filename.css -out=out.min.css
```

### 4. Todo
- Minify js file
- Minify all css/js in a directory