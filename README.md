# plugins-playground

This repo serves as a basic example for how to write standard library plugins with Golang.


## How to run this example
```bash
./build-plugins.sh && go run main.go
```

# Developer notes
The standard lib for `plugin` will only work on linux. See full restrictions on the documentation page.

The `build-plugins.sh` script will build every plugins inside the `plugins` folder. The output bin file will be stored inside `plugins/bin`.
