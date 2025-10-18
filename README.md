# Go Minifier

A lightweight command-line tool designed to simplify and streamline the minification process for CSS and JavaScript files. Minification involves the removal of unnecessary characters and spaces from code, resulting in smaller file sizes that contribute to improved web page loading times.

## 🚀 Quick Start

Get up and running in under 2 minutes:

```bash
# Install the tool
go install github.com/Alitindrawan24/go-minifier

# Minify a CSS file
go-minifier -src styles.css

# Minify a JavaScript file  
go-minifier -src script.js -opt js
```

## 📋 Requirements

- Go 1.19 or later

## 🔧 Installation & Setup

### Option 1: Direct Installation
```bash
go install github.com/Alitindrawan24/go-minifier
```

### Option 2: Local Development

**Using HTTPS:**
```bash
git clone https://github.com/Alitindrawan24/go-minifier.git
cd go-minifier
```

**Using SSH:**
```bash
git clone git@github.com:Alitindrawan24/go-minifier.git
cd go-minifier
```

## 📖 Usage

### Minify CSS Files

**Basic usage** - Creates `filename.min.css`:
```bash
go run . -src styles.css
```

**Custom output filename:**
```bash
go run . -src styles.css -out production.min.css
```

**Working with multiple CSS files:**
```bash
# Minify multiple CSS files
go run . -src main.css
go run . -src components.css -out components.min.css
go run . -src responsive.css -out responsive.min.css
```

### Minify JavaScript Files

**Basic usage** - Creates `filename.min.js`:
```bash
go run . -src script.js -opt js
```

**Custom output filename:**
```bash
go run . -src app.js -out app.min.js -opt js
```

**Working with multiple JavaScript files:**
```bash
# Minify multiple JS files
go run . -src main.js -opt js
go run . -src utils.js -out utils.min.js -opt js
go run . -src vendor.js -out vendor.min.js -opt js
```

### Command Line Options

| Option | Description | Required | Default |
|--------|-------------|----------|---------|
| `-src` | Path to the source file | ✅ | - |
| `-out` | Path to the output file | ❌ | `filename.min.css/js` |
| `-opt` | Minifier type (`css` or `js`) | ❌ | `css` |


**Valid values for `-opt`:**
- `css` - For CSS files (default)
- `js` - For JavaScript files

## 📊 Example Output

After minification, you'll see file size statistics:

```
File styles.css original size: 15 KB
File styles.min.css output size: 8 KB (reduced by 46.67%)
```

## 🛠️ Development

### Running Tests
```bash
make test
```

### Code Formatting
```bash
make fmt
```

### Linting
```bash
make lint
```

## 📝 Todo

- [ ] Minify all CSS/JS files in a directory
- [ ] Batch processing capabilities
- [ ] Additional file format support