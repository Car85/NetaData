# NetaData

![Go](https://img.shields.io/badge/Go-1.21-blue?style=flat-square&logo=go)
![Debian](https://img.shields.io/badge/Debian-Stable-red?style=flat-square&logo=debian)
![Linux](https://img.shields.io/badge/Linux-Universal-orange?style=flat-square&logo=linux)
![License](https://img.shields.io/badge/License-GPL--3.0-green?style=flat-square&logo=gnu)

## Overview

**NetaData** is a command-line application written in Go that processes YouTube-style transcription files, removing:
- Time markers.
- HTML-like tags (e.g., `<c>`).
- Repeated consecutive lines.
- Unnecessary line breaks, creating a single continuous text output.

The application is designed to simplify transcript files for easy readability and further analysis.

---

## Features

- **Time Marker Removal:** Eliminates timestamped text lines.
- **Tag Cleanup:** Strips out unwanted tags like `<c>`.
- **Duplicate Line Filtering:** Avoids redundant consecutive lines.
- **Text Flattening:** Produces a single continuous line of text without line breaks.

---

## Requirements

- **Go** 1.21 or later installed (for building manually).
  
- You can also install **NetaData** in any Debian based Linux distribution.
  
---

## Installation

### Option 1: Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/netadata.git
   cd netadata
   ```

2. Build the application:
   ```bash
   go build -o netadata main.go
   ```

3. Run the application:
   ```bash
   ./netadata <input_file>
   ```

   

### Option 2: Install via Debian Package

1. Build the `.deb` package:
   ```bash
   mkdir -p mypackage/usr/local/bin
   mkdir -p mypackage/DEBIAN
   cp netadata mypackage/usr/local/bin/
   echo "Package: netadata\nVersion: 1.0\nSection: utils\nPriority: optional\nArchitecture: amd64\nMaintainer: Your Name <your.email@example.com>\nDescription: A tool to clean transcription files." > mypackage/DEBIAN/control
   dpkg-deb --build mypackage
   ```

 You can also download directly the .deb file here:

```bash
wget https://github.com/Car85/NetaData/blob/main/netadata.deb

```

2. Install the package:
   ```bash
   sudo dpkg -i mypackage.deb
   ```


3. Run the application:
   ```bash
   netadata <input_file>
   ```

---

### Option 3 

## Usage

### Command Format

```bash
./netadata <input_file>
```

- `<input_file>`: Path to the transcription file (e.g., `transcript.txt`).
- The output will be a cleaned file named `<input_file>_cleaned.txt` in the same directory.

### Example

Input file (`transcript.txt`):
```text
00:00:01.210 --> 00:00:04.030 align:start position:0%
[Music]
hello<00:00:04.560><c> and</c><00:00:04.680><c> welcome</c>
hello and welcome
hello and welcome
```

Command:
```bash
./netadata transcript.txt
```

Output file (`transcript_cleaned.txt`):
```text
[Music] hello and welcome 
```

---


## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

---

## Contributing

Feel free to open issues or submit pull requests for enhancements and fixes. Contributions are welcome!


