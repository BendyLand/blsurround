# Bland Surround Tool

A simple text surround tool written in Go!

## Usage 

This tool reads the system clipboard to a string and applies the first provided command-line argument to the beginning and end of that string. Then it saves the new string back to the system clipboard. Certain tokens, like '(', '{', and the other opening/closing pairs are automatically matched. 

**Note:** Most non-alphanumeric arguments will requite a backslash:
```bash
blsurround \"
```

## Future Plans

 - Add flag to ignore opening/closing pairs and surround with the same token.
     - e.g. `echo "test" | pbcopy | "blsurround \{` -> {test}
