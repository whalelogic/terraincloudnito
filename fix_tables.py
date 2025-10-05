import json

file_path = "/home/whaler/github/templ-go/allposts.json"

with open(file_path, 'r') as f:
    posts = json.load(f)

for post in posts:
    if post['slug'] == 'python-regex':
        content = post['content']
        
        content = content.replace('   Syntax       | Description                                      | Example                     |\n |--------------|--------------------------------------------------|-----------------------------|\n | `[a-z]`      | Matches any lowercase letter.                    | `a`, `b`, `z`               |\n | `[A-Z]`      | Matches any uppercase letter.                    | `A`, `B`, `Z`               |\n | `[0-9]`      | Matches any digit.                               | `0`, `1`, `9`               |\n | `[a-zA-Z0-9]`| Matches any alphanumeric character.              | `a`, `B`, `1`               |\n | `[^a-z]`     | Matches any character **except** lowercase letters. | `A`, `1`, `@`           |',
                              '| Syntax       | Description                                      | Example                     |
|--------------|--------------------------------------------------|-----------------------------|
| `[a-z]`      | Matches any lowercase letter.                    | `a`, `b`, `z`               |
| `[A-Z]`      | Matches any uppercase letter.                    | `A`, `B`, `Z`               |
| `[0-9]`      | Matches any digit.                               | `0`, `1`, `9`               |
| `[a-zA-Z0-9]`| Matches any alphanumeric character.              | `a`, `B`, `1`               |
| `[^a-z]`     | Matches any character **except** lowercase letters. | `A`, `1`, `@`           |')
        
        content = content.replace(' | Syntax | Description                                      | Example                     |
 |--------|--------------------------------------------------|-----------------------------|
 | `^`    | Matches the **beginning** of the string.         | `^hello` matches "hello world" |
 | `$`    | Matches the **end** of the string.               | `world$` matches "hello world" |', 
                              '| Syntax | Description                                      | Example                     |
|--------|--------------------------------------------------|-----------------------------|
| `^`    | Matches the **beginning** of the string.         | `^hello` matches "hello world" |
| `$`    | Matches the **end** of the string.               | `world$` matches "hello world" | 
')

        content = content.replace(' | Syntax  | Description                                      | Example                     |
 |---------|--------------------------------------------------|-----------------------------|
 | `*`     | Matches **zero or more** occurrences.             | `a*` matches "", "a", "aa" |
 | `+`     | Matches **one or more** occurrences.              | `a+` matches "a", "aa"  |
 | `?`     | Matches **zero or one** occurrence.               | `a?` matches "", "a"    |
 | `{n}`   | Matches **exactly n** occurrences.                | `a{3}` matches "aaa"      |
 | `{n,}`  | Matches **n or more** occurrences.                | `a{2,}` matches "aa", "aaa" |
 | `{n,m}` | Matches **between n and m** occurrences.          | `a{2,4}` matches "aa", "aaa", "aaaa" |', 
                              '| Syntax  | Description                                      | Example                     |
|---------|--------------------------------------------------|-----------------------------|
| `*`     | Matches **zero or more** occurrences.             | `a*` matches "", "a", "aa" |
| `+`     | Matches **one or more** occurrences.              | `a+` matches "a", "aa"  |
| `?`     | Matches **zero or one** occurrence.               | `a?` matches "", "a"    |
| `{n}`   | Matches **exactly n** occurrences.                | `a{3}` matches "aaa"      |
| `{n,}`  | Matches **n or more** occurrences.                | `a{2,}` matches "aa", "aaa" |
| `{n,m}` | Matches **between n and m** occurrences.          | `a{2,4} ` matches "aa", "aaa", "aaaa" |')

        content = content.replace(' | Syntax | Description                                      | Example                     |
 |--------|--------------------------------------------------|-----------------------------|
 | `.`    | Matches **any character** (except newline).      | `a.c` matches "abc", "a1c" |
 | `\d`   | Matches a **digit**.                             | `\d` matches `1`, `2`       |
 | `\w`   | Matches a **word character**.                    | `\w` matches `a`, `_`, `1`  |
 | `\s`   | Matches **whitespace**.                          | `\s` matches " ", `\t`    |
 | `\D`   | Matches a **non-digit** character.               | `\D` matches `a`, `@`       |
 | `\W`   | Matches a **non-word** character.                | `\W` matches `@`, `#`       |
 | `\S`   | Matches a **non-whitespace** character.          | `\S` matches `a`, `1`       |', 
                              '| Syntax | Description                                      | Example                     |
|--------|--------------------------------------------------|-----------------------------|
| `.`    | Matches **any character** (except newline).      | `a.c` matches "abc", "a1c" |
| `\d`   | Matches a **digit**.                             | `\d` matches `1`, `2`       |
| `\w`   | Matches a **word character**.                    | `\w` matches `a`, `_`, `1`  |
| `\s`   | Matches **whitespace**.                          | `\s` matches " ", `\t`    |
| `\D`   | Matches a **non-digit** character.               | `\D` matches `a`, `@`       |
| `\W`   | Matches a **non-word** character.                | `\W` matches `@`, `#`       |
| `\S`   | Matches a **non-whitespace** character.          | `\S` matches `a`, `1`       |')

        content = content.replace(' | Syntax       | Description                                      | Example                     |
 |--------------|--------------------------------------------------|-----------------------------|
 | `(pattern)`  | Groups the pattern.                              | `(abc)`                     |
 | `\1`, `\2`   | Refer to captured groups (**backreferences**).   | `(a).\1` matches "aba"    |', 
                              '| Syntax       | Description                                      | Example                     |
|--------------|--------------------------------------------------|-----------------------------|
| `(pattern)`  | Groups the pattern.                              | `(abc)`                     |
| `\1`, `\2`   | Refer to captured groups (**backreferences**).   | `(a).\1` matches "aba"    |')
        
        content = content.replace(' | Type | Syntax |
|------|--------|
| Email Validation | `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\$` |
| Extracting URLs | `https?://[^\s]+` |
| Finding Dates | `\d{2}-\d{2}-\d{4}` |
| Password Strength Check | `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@\$!%*?&])[A-Za-z\d@$!%*?&]{8,}$` |', 
                              '| Type | Syntax |
|------|--------|
| Email Validation | `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\$` |
| Extracting URLs | `https?://[^\s]+` |
| Finding Dates | `\d{2}-\d{2}-\d{4}` |
| Password Strength Check | `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@\$!%*?&])[A-Za-z\d@$!%*?&]{8,}$` |')

        post['content'] = content

with open(file_path, 'w') as f:
    json.dump(posts, f, indent=2)