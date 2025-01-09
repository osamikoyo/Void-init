# Void-init

A command-line tool for initializing and managing Void Linux packages and templates efficiently.

## Features

- Quick package template initialization
- Package installation handling
- Modern project structure
- Follows Void Linux packaging guidelines

## Installation

```bash
# Using go install
go install github.com/osamikoyo/void-init@latest

# Or clone and build manually
git clone https://github.com/osamikoyo/void-init.git
cd void-init
go build
```

## Usage

```bash
# Initialize a new package template
void-initer init [package-name]

# Install a package
void-initer install [package-name]
```

### Basic Commands

```bash
# Initialize a new package
void-initer init mypackage

# Install a package
void-initer install mypackage

# Show help
void-initer --help
```

## Project Structure

When you initialize a new package, it will create the following structure:

```
mypackage/
├── cmd
│   └── [command files]
├── internal/
│   └── [package name]
├── pkg/
│   └── [package name]
└── README.md
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you have any questions or need help, please:
- Open an issue
- Check the Void Linux documentation
- Join the Void Linux community forums

## Acknowledgments

- Thanks to all contributors
- Inspired by Void Linux package management
- Built for the Void Linux community