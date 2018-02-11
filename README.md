# meerkat-cli

[![Build Status](https://travis-ci.org/mattevans/meerkat-cli.svg?branch=master)](https://travis-ci.org/mattevans/meerkat-cli)

CLI based Toy Robot Simulator written in Go.

> The application is a simulation of a toy robot moving on a square tabletop, of dimensions 5 units x 5 units.

### Install & Run

1) Install any [dependencies](#dependencies)
```
go get -u github.com/mattevans/meerkat-cli
```

2) Run it!
```
meerkat-cli
```

> Tested on [Go versions 1.5.x - 1.9.x](https://travis-ci.org/mattevans/meerkat-cli)

### Tests

```
go test -v ./...
```

or, to test and see coverage, run...

```
./coverage.sh --html
```

### Valid Commands

| Command        | Arguments            | Description                            |
| ------------- |:--------------------- | -------------------------------------- |
| `PLACE`         | `x`, `y`, `facing`    | PLACE will put the toy robot on the table in position `X`, `Y` and `facing` |
| `MOVE`          | -                     | MOVE will move the toy robot one unit forward in the direction it is currently facing |
| `RIGHT`         | -                     | RIGHT will rotate the robot 90 degrees right without changing the position of the robot  |
| `LEFT`          | -                     | LEFT will rotate the robot 90 degrees left without changing the position of the robot |
| `REPORT`        | -                     | REPORT will announce the X,Y and orientation of the robot |
| `EXIT`          | -                     | Will close the process |

### Example

![Example meerkat-cli](/_docs/example.png?raw=true "Example CLI App in use")

### Dependencies

Packages used:

- [Table Writer](https://github.com/olekukonko/tablewriter) - Used to output a visual representation of your robots position when executing `REPORT` command
- [CLI](https://github.com/urfave/cli) - A fast and simple package leveraged to bootstrap cli components of `meerkat-cli`.