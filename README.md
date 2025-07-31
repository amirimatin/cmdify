# cmdify

`cmdify` is a minimal, composable, and testable CLI utility toolkit for Go developers.

It simplifies common command-line interactions like:
- Executing system commands (with or without TTY)
- Asking interactive questions (input, password, select, yes/no)
- Building delightful command-line tools with minimal code

---

## ✨ Features

- ✅ Run shell commands with structured result (`stdout`, `exit code`, `error`)
- ✅ Full interactive TTY support (`htop`, `vim`, `docker exec -it`)
- ✅ User prompts: input, password, yes/no, select
- ✅ Default selection support for better UX
- ✅ Dependency-free runtime (except for `promptui` and `creack/pty`)

---
## ✨ DSL Design

- 🧠 Fluent DSL API: `cmdify.Run(...).WithTTY().Must()`
- 💬 Prompt input, password, confirmation, select, multi-select
- 🖥️ Run shell commands with or without TTY
- 🎯 Default option support
- ✅ Lightweight: no dependencies except `promptui` and `creack/pty`
---

## 📦 Installation

```bash
go get github.com/amirimatin/cmdify
```

## 🚀 Quick Examples

### 📌 Run a system command
✅ Ask for user input

```go
name := cmdify.Ask("What's your name?").Input()
pass := cmdify.Ask("Enter password").Password()
cmdify.Ask("Continue?").YesNo()
```
✅ Select single option

```go
choice := cmdify.Select("Choose environment").
  FromList([]cmdify.SelectOption{
    {"Dev", "Development"},
    {"Prod", "Production"},
  }).
  WithDefault(0).
  Run()

```

✅ Select multiple options

```go
tasks := cmdify.MultiSelect("Pick tasks").
  FromList([]cmdify.SelectOption{
    {"Build", "Compile code"},
    {"Test", "Run tests"},
    {"Deploy", "Ship to prod"},
  }).
  Run()

```

✅ Run system commands

```go
cmdify.Run("ls", "-la").Must()
cmdify.Run("htop").WithTTY().Must()

```

## 🔖 License
MIT © amirimatin

This project is licensed under the terms of the MIT license.
You are free to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, under the following conditions:

```text
MIT License

Copyright (c) 2025 amirimatin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights  
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell     
copies of the Software, and to permit persons to whom the Software is         
furnished to do so, subject to the following conditions:                      

The above copyright notice and this permission notice shall be included in    
all copies or substantial portions of the Software.                           

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR    
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,      
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE   
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER       
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN     
THE SOFTWARE.

---

## 🔒 Ethical and Legal Use Restriction

Use of this software for unlawful, unethical, or Islamically impermissible purposes — including but not limited to: surveillance, harm to innocent persons, support of oppressive regimes, or development of weapons — is strictly prohibited.

Any such use is considered a violation of this license, and the author bears no responsibility for resulting consequences.
```

```text
> ⚠️ Ethical Notice:  
> Use of this software for unlawful or Islamically impermissible purposes (such as surveillance, harming individuals, or enabling oppression) is strictly forbidden under the extended license terms.
```