## Using the project Makefile
The `dir` task creates a directory for the given day and populates it with a `main.go` file and a file for the day's input.
To run the `dir` task make sure you have an `.env` file in the project root directory with your Advent of Code website session cookie value.
```
// .env
session={SESSION_COOKIE_HERE}
```

The `run` task runs the solution for the given day against the input in the `input.txt` file. Make sure that file exists.

### Initializing a directory for a new day
```bash
make dir day=3
```

### Running the solution for a particular day
```bash
make run day=3
```
Make sure you have the appropriate input file in the solution directory.
See the `main.go` file of a particular day for more details.
