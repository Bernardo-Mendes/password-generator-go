# password-generator-go

Simple password generator written in Go.

## Requirements

- Go installed on your machine. Official documentation and downloads:

	- Documentation: https://go.dev/doc/
	- Downloads: https://go.dev/dl/

	Use a recent stable Go release.

## Clone and run

Minimal steps to clone, build and run the project on Windows (PowerShell). Adjust the `git clone` URL if you use a fork.

1. Clone the repository

```powershell
git clone https://github.com/Bernardo-Mendes/password-generator-go.git
cd password-generator-go
```

2. Run without building (quick test)

```powershell
go run .
```

3. Build and run the binary (Windows / PowerShell)

```powershell
go build -o password-generator.exe .
.\password-generator.exe
```

Or on UNIX-like systems (Linux / macOS):

```bash
go build -o password-generator .
./password-generator
```

## What the program does

The program generates a random password containing uppercase letters, lowercase letters, numbers and special characters. It guarantees at least one character from each category and mixes the generated characters before printing the final password to stdout.

The implementation uses `crypto/rand` for cryptographically secure randomness.

## Notes

- Run the binary or `go run .` again to generate another password.
- If you want to use this logic in another program, you can extract the generation code into a package and call it from other Go code.

## License

See the `LICENSE` file in this repository.
