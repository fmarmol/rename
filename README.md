# Examples

```
touch foo.txt bar.txt baz.txt

rename '*.txt' '$basename.md'
```

or 

```
touch foo.txt bar.txt baz.txt

rename '*.txt' '$basename.$ext.backup'
```
