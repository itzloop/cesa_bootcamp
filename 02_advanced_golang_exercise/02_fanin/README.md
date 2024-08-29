# Sum of the N natrual number

You are asked to calculate the sum N natrual numbers using the [Fan-in](https://docs.google.com/presentation/d/19HieTZFLNatVq1FdT1V25tmZoL-m3qPpyfvFJiU7rNc/edit#slide=id.g2f2e8702a5d_0_367) pattern.
In the second session of Advanced go, we introduced a patternt called fan-in.
This pattern is used to aggregate the data from multiple channels into a single channel.
In the [code](https://github.com/itzloop/cesa_bootcamp/blob/main/02_advanced_golang/04_fan_in/fan_in.go#L35C1-L35C54), 
We have combined only 2 channels. Now you are asked to extend the functionality of `runFanIn` to accept `N` channels.


## Requirements

You need to complete the implementaion of `runFanIn(channels ...<-chan int)` and combine the result of
`N` channels. As you can see, `runFanIn` is receiving a variadic parameters meaning the number of channels
is not a constant and this function can be called with any number of channels.

### How to test

```golang
go test ./...
```

To avoid cached tests, run with `-count=1` flag

```golang
go test ./... -count=1
```

To see verbose logs, run with `-v` flag

```golang
go test ./... -v
```

## Notes

- If you are stuck or in need of any help, feel free to pop a question in the group chat or in private.
