# Earliest Start Timewarrior extension
This [Timewarrior extension](https://timewarrior.net/docs/extensions/) extracts the earliest start and the latest end from the list of intervals provided by Timewarrior.

## Definitions
Earliest start
: The interval with the `start` field having the smallest hour and minute of all the list.

Latest end
: The interval with the `end` field having the greatest hour and minute of all the list.

## Installation
1. Download the latest executable for your operating system from the [releases page](https://github.com/crossbone-magister/earliest-start-timew-extension/releases)
2. Add it to the Timewarrior extension folder as described in the [documentation](https://timewarrior.net/docs/api/).
3. Verify that the extension is active and installed by running `timew extensions`.

## Usage
In a terminal window, run `timew earliest-start`. An example output could be:
```bash
Earliest Start: 09:00 on 2023-01-23
Latest Stop: 18:37 on 2023-01-24
```
