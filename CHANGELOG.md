# Changelog

## [v0.1.0] - 2025-07-31

### Added
- DSL-style API for building interactive CLI tools
- `cmdify.Run(...).WithTTY().Must()` for shell commands
- `cmdify.Ask(...).Input()`, `.Password()`, `.YesNo()` for prompts
- `cmdify.Select(...).FromList(...).WithDefault(...).Run()`
- `cmdify.MultiSelect(...).FromList(...).Run()`
- Full examples in `examples/` folder
- Ethical license restriction notice