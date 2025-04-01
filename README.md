# pathcodec

[![Go Version][GoVer-Image]][GoDoc-Url] [![License MIT][License-Image]][License-Url] [![GoDoc][GoDoc-Image]][GoDoc-Url] [![Go Report Card][ReportCard-Image]][ReportCard-Url]

[GoVer-Image]: https://img.shields.io/badge/Go-1.24%2B-blue
[GoDoc-Url]: https://pkg.go.dev/github.com/itpey/pathcodec
[GoDoc-Image]: https://pkg.go.dev/badge/github.com/itpey/pathcodec.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/itpey/pathcodec
[ReportCard-Image]: https://goreportcard.com/badge/github.com/itpey/pathcodec
[License-Url]: https://github.com/itpey/pathcodec/blob/main/LICENSE
[License-Image]: https://img.shields.io/github/license/itpey/pathcodec

A lightweight Go package for compressing and decompressing
[Telegram vector thumbnails](https://core.telegram.org/api/files#vector-thumbnails).

## Installation

```sh
go get -u github.com/itpey/pathcodec
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/itpey/pathcodec"
)

func main() {
	compressed, err := pathcodec.Compress("M257,455c-56,0-109-25-146-65-143-156,31-397,224-318,201,83,136,386-78,383z")
	if err != nil {
		fmt.Println("Compression error:", err)
		return
	}
	fmt.Println("Compressed:", compressed)

	decompressed, err := pathcodec.Decompress(compressed)
	if err != nil {
		fmt.Println("Decompression error:", err)
		return
	}
	fmt.Println("Decompressed:", decompressed)
}
```

## API

### `func Compress(path string) ([]byte, error)`

Compresses a path string into a compact byte format. The input must start with 'M' and end with 'z'.

### `func Decompress(encoded []byte) (string, error)`

Decompresses a byte slice back into a path string.

## Feedback and Contributions

If you encounter any issues or have suggestions for improvement, please [open an issue](https://github.com/itpey/pathcodec/issues) on GitHub.

We welcome contributions! Fork the repository, make your changes, and submit a pull request.

## License

pathcodec is open-source software released under the MIT License. You can find a copy of the license in the [LICENSE](https://github.com/itpey/pathcodec/blob/main/LICENSE) file.

## Author

pathcodec was created by [itpey](https://github.com/itpey)
