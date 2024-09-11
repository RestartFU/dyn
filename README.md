# dyn
a linux package manager 

## installing from source (requires Go installed):
```sh
  git clone --depth=1 https://github.com/RestartFU/dyn
  cd dyn
  sudo make install
```

# installing a package:
```sh
dyn install <pkg>
```
## example:
```sh
dyn install discord
```

# removing a package:
```sh
dyn remove <pkg>
```
## example:
```sh
dyn remove discord
```

# updating a package:
```sh
dyn update <pkg>
```
## example:
```sh
dyn update discord
```

# fetching the dyn-pkg repository:
```sh
dyn fetch
```

## you may also fetch and install at the same time:
```sh
dyn fetch install discord
```
