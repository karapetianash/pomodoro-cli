## Description
The Pomodoro is a time management approach that allows you to focus
on tasks by defining a short time interval to concentrate on it, called
Pomodoro, followed by short and long breaks to allow you to rest and
reprioritize tasks. In general, a Pomodoro interval lasts 25 minutes while
breaks are typically 5 and 15 minutes.

### Interactive Pomodoro Timer Screen:

![Pomodoro Screen](https://github.com/karapetianash/pomodoro-cli/blob/main/pomoFinalScreen.PNG "Pomodoro Screen")

### Usage:
`pomo [flags]`

### Flags:

        --config  string    config file (default is $HOME/.pomo.yaml)

    -d, --db      string    database file (default "pomo.db")
    -h, --help              help for pomo
    -l, --long    duration  long break duration (default 15m0s)
    -p, --pomo    duration  pomodoro duration (default 25m0s)
    -s, --short   duration  short break duration (default 5m0s)
